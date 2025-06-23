package internal

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func RedisTx() {
	fmt.Println("## Redis事务操作示例")
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()
	// 开启事务
	tx := rdb.TxPipeline()
	// 事务操作
	tx.Set(ctx, "key1", "value1", 0)
	tx.Set(ctx, "key2", "value2", 0)
	tx.Incr(ctx, "counter")
	// 执行事务
	_, err := tx.Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// 验证结果
	val1, _ := rdb.Get(ctx, "key1").Result()
	val2, _ := rdb.Get(ctx, "key2").Result()
	counter, _ := rdb.Get(ctx, "counter").Result()
	fmt.Printf("key1: %s, key2: %s, counter: %s\n", val1, val2, counter)
}
