package main

import "redis/internal"

func main() {
	internal.RedisBasic()
	internal.RedisPool()
	internal.RedisTx()
	internal.RedisPubsub()
}
