package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
	"gomall/apps/seckill/cmd/mq/internal/config"
	"gomall/apps/seckill/cmd/mq/internal/service"
)

var configFile = flag.String("f", "etc/seckill.yaml", "the etc file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	srv := service.NewService(c)
	queue := kq.MustNewQueue(c.Kafka, kq.WithHandle(srv.Consume))
	defer queue.Stop()

	fmt.Printf("seckill start")
	queue.Start()
}
