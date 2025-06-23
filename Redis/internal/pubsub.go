package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func RedisPubsub() {
	fmt.Println("## Redis发布/订阅示例，连接到本地Redis服务器")
	// 创建Redis客户端
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()
	// 订阅频道
	pubsub := rdb.Subscribe(ctx, "mychannel")
	defer pubsub.Close()
	// 启动goroutine接收消息
	go func() {
		for {
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				panic(err)
			}
			fmt.Printf("收到消息: %s\n", msg.Payload)
		}
	}()
	// 发布消息
	for i := 0; i < 5; i++ {
		err := rdb.Publish(ctx, "mychannel", fmt.Sprintf("消息%d", i)).Err()
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
	}
}
