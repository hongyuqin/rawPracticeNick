package main

import (
	"context"
	"fmt"
	"klook.libs/errors"
	"klook.libs/logger"
	"klook.libs/mq/kmq"
)

func testProduce() error {
	// Startup 是非阻塞的
	fmt.Println("begin testProduce")
	logger.Info("begin testProduce")
	p, err := kmq.NewProducer().SetTopic("kmq-test").Startup()
	if err != nil {
		fmt.Println("failed to startup")
		return errors.Errorf(err, "failed to startup")
	}
	ctx := context.Background()
	msg := &struct {
		Seq int
	}{
		Seq: 10,
	}
	err = p.Produce(ctx, msg)
	if err != nil {
		// ...
	}
	// 关闭producer
	// kmq.ShutdownAllProducers()
	// 放在 main.go 的 o.DeferFunc 中, 服务关闭时shutdown所有producer
	// Shutdown 是阻塞的, 会等待至消息处理完毕

	// 如果在单元测试中, 可以单独关闭 producer
	err = p.Shutdown()
	if err != nil {
		logger.Errorf("failed to shutdown", err)
	}
	return nil
}
func main() {
	testProduce()
}
