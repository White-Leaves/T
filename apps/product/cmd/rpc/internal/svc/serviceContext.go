package svc

import (
	"github.com/jinzhu/gorm"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/sync/singleflight"

	"gomall/apps/product/cmd/rpc/internal/config"
	"gomall/apps/product/model"
	"gomall/pkg/orm"

	"time"
)

const localCacheExpire = time.Duration(time.Second * 60)

type ServiceContext struct {
	Config config.Config

	ProductModel model.ProductModel
	BizRedis     *redis.Redis
	LocalCache   *collection.Cache
	orm          *gorm.DB
	SingleGroup  singleflight.Group
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	localCache, err := collection.NewCache(localCacheExpire)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,

		ProductModel: model.NewProductModel(sqlConn, c.Cache),
		BizRedis:     redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
		LocalCache:   localCache,
		orm: orm.NewMysql(&orm.Config{
			DSN:         c.DB.DataSource,
			Active:      20,
			Idle:        10,
			IdleTimeout: time.Hour * 24,
		}),
	}
}
