package configs

import (
	"context"
	"github.com/beego/beego/v2/core/logs"
	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
	"log"
)

var Rdb *redis.Client
var Ctx context.Context

// NewRedis 初始化Redis数据库
func NewRedis() {
	cfg := Config().Redis
	Rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	ping := Rdb.Ping(context.Background())
	if ping.Err() != nil {
		log.Fatalf("redis 启动失败: %v", ping.Err())
	}
	_ = redislock.New(Rdb)
	Ctx = context.Background()
	logs.Info("Redis数据库初始化连接成功")
}
