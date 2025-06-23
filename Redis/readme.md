# 使用go语言进行redis操作

1. 安装Redis服务器
    - 使用 redis-server.exe redis.windows.conf 命令来启动redis服务
    - 使用select命令来切换库，总共有16个库，默认使用0号库
    - 使用set 和get来设置键值和取出
    - 使用shutdown来关闭redis服务
2. 安装go的Redis客户端库:   go get github.com/go-redis/redis/v8

3. 基础操作示例  redis_basic.go  包括连接设置、键值操作和哈希表操作。
4. 连接池配置    redis_pool.go  这段代码展示了如何配置Redis连接池，优化高并发场景下的性能。
5. 事务操作      redis_tx.go 这段代码展示了如何使用Redis事务，确保多个操作的原子性。
6. 发布订阅模式   redis_pubsub.go 这段代码展示了Redis的发布订阅功能，实现简单的消息队列。

## 操作流程总结
1. 安装Redis服务器并确保其运行
2. 在Go项目中引入Redis客户端库
3. 创建Redis客户端实例并配置连接参数
4. 使用context.Context管理请求生命周期
5. 执行各种Redis操作（SET/GET/HASH等）
6. 处理返回结果和错误
7. 合理使用连接池和事务
8. 根据需求选择合适的数据结构和操作