package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func RedisPool() {
	fmt.Println("## 创建Redis连接池，连接到本地Redis服务器")
	// 创建Redis连接池
	// 连接池配置可以根据实际需求调整
	rdb := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		PoolSize:     50,               // 连接池大小
		MinIdleConns: 10,               // 最小空闲连接数
		MaxConnAge:   30 * time.Minute, // 连接最大存活时间
		PoolTimeout:  30 * time.Second, // 获取连接超时时间
	})
	ctx := context.Background()
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong) // 输出PONG表示连接成功
}
