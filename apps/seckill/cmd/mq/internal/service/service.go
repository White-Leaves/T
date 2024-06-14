package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"gomall/apps/order/cmd/rpc/order"
	"gomall/apps/product/cmd/rpc/product"
	"gomall/apps/seckill/cmd/mq/internal/config"
	"log"
	"sync"
)

const (
	chanCount   = 10
	BufferCount = 1024
)

type KafkaData struct {
	Uid int64 `json:"uid"`
	Pid int64 `json:"pid"`
}

type Service struct {
	c          config.Config
	ProductRpc product.Product
	OrderRpc   order.Order

	waiter      sync.WaitGroup
	massageChan []chan *KafkaData
}

func (s *Service) consume(ch chan *KafkaData) {
	defer s.waiter.Done()

	for {
		m, ok := <-ch
		if !ok {
			log.Fatal("seckill rmq exit")
		}
		fmt.Printf("consume msg:%+v\n", m)
		_, err := s.ProductRpc.CheckAndUpdateStock(context.Background(), &product.CheckAndUpdateStockReq{ProductId: m.Pid})
		if err != nil {
			logx.Errorf("CheckAndUpdataStock pid:%d error:%v", m.Pid, err)
			return
		}
		_, err = s.OrderRpc.CreateOrder(context.Background(), &order.CreateOrderReq{UserId: m.Uid, ProductId: m.Pid})
		if err != nil {
			logx.Errorf("Create order uid:%d pid:%d error:%v", m.Uid, m.Pid, err)
		}

		_, err = s.ProductRpc.UpdateProductStock(context.Background(), &product.UpdateProductStockReq{ProductId: m.Pid, Num: 1})
		if err != nil {
			logx.Errorf("Update product stock uid:%v pid:%v error:%v", m.Uid, m.Pid, err)
		}

	}
}

func (s *Service) Consume(_ string, value string) error {
	logx.Infof("Consume value: %s\n", value)
	var data []*KafkaData
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return err
	}
	for _, d := range data {
		s.massageChan[d.Pid%chanCount] <- d
	}
	return nil
}

func NewService(c config.Config) *Service {
	s := &Service{
		c:           c,
		ProductRpc:  product.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
		OrderRpc:    order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		massageChan: make([]chan *KafkaData, chanCount),
	}

	for i := 1; i < chanCount; i++ {
		ch := make(chan *KafkaData, BufferCount)
		s.massageChan[i] = ch
		s.waiter.Add(1)
		go s.consume(ch)
	}

	return s
}

//var dtmServer = "etcd://localhost:2379/dtmservice"
//
//func (s *Service) consumeDTM(ch chan *KafkaData) {
//	defer s.waiter.Done()
//
//	productServer, err := s.c.ProductRpc.BuildTarget()
//	if err != nil {
//		log.Fatalf("s.c.ProductRpc error:%v", err)
//	}
//	orderServer, err := s.c.OrderRpc.BuildTarget()
//	if err != nil {
//		log.Fatalf("s.c.OrderRpc error:%v", err)
//	}
//
//	for {
//		m, ok := <-ch
//		if !ok {
//			log.Fatalf("seckill rmq exit")
//		}
//		fmt.Printf("consume msg: %+v\n", m)
//
//		gid := dtmgrpc.MustGenGid(dtmServer)
//		err := dtmgrpc.TccGlobalTransaction(dtmServer, gid, func(tcc *dtmgrpc.TccGrpc) error {
//			if e := tcc.CallBranch(
//				&product.UpdateProductStockReq{ProductId: m.Pid, Num: 1},
//				productServer+"/product.Product/CheckProductStock",
//				productServer+"/product.Product/UpdateProductStock",
//				productServer+"/product.Product/DecrStockCancel",
//				&product.UpdateProductStockResp{}); err != nil {
//				logx.Errorf("tcc.CallBarch server:%s error: %v", productServer, err)
//				return e
//			}
//			if e := tcc.CallBranch(
//
//				&order.CreateOrderReq{UserId: m.Uid, ProductId: m.Pid},
//				orderServer+"/order.Order/CreateOrderCheck",
//				orderServer+"/order.Order/CreateOrder",
//				orderServer+"/order.Order/CreateOrderDTMCancel",
//				&order.CreateOrderResp{},
//			); err != nil {
//				logx.Errorf("tcc.CallBranch server: %s error: %v", orderServer, err)
//				return e
//			}
//			return nil
//		})
//		logger.FatalIfError(err)
//	}
//}
