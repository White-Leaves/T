package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/limit"
	"gomall/apps/product/cmd/rpc/product"
	"gomall/pkg/batcher"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"time"

	"gomall/apps/seckill/cmd/rpc/internal/svc"
	"gomall/apps/seckill/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	limitPerid        = 10
	limitQuota        = 10
	seckillUserPrefix = "seckill#u#"
	localCacheExpire  = time.Second * 60

	batcherSize     = 100
	batcherBuffer   = 100
	batcherWorker   = 10
	batcherInterval = time.Second
)

type KafkaData struct {
	Uid int64 `json:"uid"`
	Pid int64 `json:"pid"`
}

type SeckillOrderLogic struct {
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	limiter    *limit.PeriodLimit
	localCache *collection.Cache
	batcher    *batcher.Batcher
	logx.Logger
}

func NewSeckillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillOrderLogic {
	localCache, err := collection.NewCache(localCacheExpire)
	if err != nil {
		panic(err)
	}

	s := &SeckillOrderLogic{
		ctx:        ctx,
		svcCtx:     svcCtx,
		Logger:     logx.WithContext(ctx),
		localCache: localCache,
		limiter:    limit.NewPeriodLimit(limitPerid, limitQuota, svcCtx.BizRedis, seckillUserPrefix),
	}

	b := batcher.New(
		batcher.WithSize(batcherSize),
		batcher.WithBuffer(batcherBuffer),
		batcher.WithWorker(batcherWorker),
		batcher.WithInterval(batcherInterval),
	)

	b.Sharding = func(key string) int {
		pid, _ := strconv.ParseInt(key, 10, 64)
		return int(pid) % batcherWorker
	}

	b.Do = func(ctx context.Context, val map[string][]interface{}) {
		var msgs []*KafkaData
		for _, vs := range val {
			for _, v := range vs {
				msgs = append(msgs, v.(*KafkaData))
			}
		}
		kd, err := json.Marshal(msgs)
		if err != nil {
			logx.Errorf("Batcher.Do json.Marshal msgs:%v error:%v", msgs, err)
		}
		if err = s.svcCtx.KafkaPusher.Push(string(kd)); err != nil {
			logx.Errorf("KafkaPusher.Push kd:%v error: %v", string(kd), err)
		}
	}
	s.batcher = b
	s.batcher.Start()
	return s
}

func (l *SeckillOrderLogic) SeckillOrder(in *pb.SeckillOrderReq) (*pb.SeckillOrderResp, error) {
	code, _ := l.limiter.Take(strconv.FormatInt(in.Uid, 10))
	if code == limit.OverQuota {
		return nil, status.Errorf(codes.OutOfRange, "Number of requests exceeded the limit")
	}
	p, err := l.svcCtx.ProductRpc.Product(l.ctx, &product.ProductItemRequest{Id: in.Pid})
	if err != nil {
		return nil, err
	}
	if p.Stock <= 0 {
		return nil, status.Errorf(codes.OutOfRange, "Insufficient stock")
	}
	if err = l.batcher.Add(strconv.FormatInt(in.Pid, 10), &KafkaData{Uid: in.Uid, Pid: in.Pid}); err != nil {
		logx.Errorf("l.batcher,Add uid: %d pid: %d error: %v", in.Uid, in.Pid, err)
	}

	return &pb.SeckillOrderResp{}, nil
}
