package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func RedisBasic() {
	// 创建Redis客户端
	fmt.Println("## 创建Redis客户端，连接到本地Redis服务器")
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址
		Password: "",               // 密码
		DB:       0,                // 使用默认DB
	})

	ctx := context.Background()

	// 设置键值对
	fmt.Println("设置键值对")
	err := rdb.Set(ctx, "name", "Alice", 10*time.Minute).Err()
	if err != nil {
		panic(err)
	}

	// 获取值
	fmt.Println("获取值")
	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name:", val)

	// 设置哈希表
	fmt.Println("设置哈希表")
	err = rdb.HSet(ctx, "user:1000", map[string]interface{}{
		"name":  "Bob",
		"email": "bob@example.com",
	}).Err()
	if err != nil {
		panic(err)
	}

	// 获取哈希表字段
	fmt.Println("获取哈希表字段")
	email, err := rdb.HGet(ctx, "user:1000", "email").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("email:", email)

	// 删除键
	fmt.Println("删除键")
	rdb.Del(ctx, "name")
	fmt.Println("Redis基础操作完成")
}
