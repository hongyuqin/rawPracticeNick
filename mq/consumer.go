package main

import (
	"context"
	"klook.libs/errors"
	"klook.libs/logger"
	"klook.libs/mq/kmq"
	"time"
)

func testConsume() error {
	// 消息处理的handler
	handleFunc := func(ctx context.Context, body []byte) error {
		logger.Infof("msg body:%s", body)
		time.Sleep(time.Second)
		// 返回错误时消息会重新入队
		return nil
	}
	// 链式调用启动consumer
	// 设置topic channel
	// 设置并发处理数
	// 设置消息处理的handler
	// Startup 是非阻塞的
	c, err := kmq.NewConsumer().SetTopicChannel("kmq-test", "xxx").
		SetConcurrent(5).
		SetHandleFunc(handleFunc).
		Startup()
	if err != nil {
		return errors.Errorf(err, "failed to startup")
	}
	// 关闭consumer
	// kmq.ShutdownAllConsumers()
	// 放在 main.go 的 o.BeforeGracefulShutdown 中, 服务关闭时shutdown所有consumer
	// Shutdown 是阻塞的, 会等待正在处理的消息处理完毕再关闭

	// 如果在单元测试中, 可以单独关闭consumer
	err = c.Shutdown()
	if err != nil {
		logger.Errorf("failed to shutdown", err)
	}
	return nil
}

func main() {
	testConsume()
}
