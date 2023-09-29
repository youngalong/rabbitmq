// Package main @Author: youngalone [2023/8/15]
package main

import (
	"github.com/gookit/slog"
	amqp "github.com/rabbitmq/amqp091-go"
	"rabbitmq/config"
	"rabbitmq/util"
)

func main() {
	// 连接rabbitmq
	conn, err := amqp.Dial(config.URL)
	util.FailOnError(err, "rabbitmq连接失败")
	defer conn.Close()

	// 创建通道
	ch, err := conn.Channel()
	util.FailOnError(err, "创建通道失败")
	defer ch.Close()

	// 声明队列
	q, err := ch.QueueDeclare(
		"hello",
		false, // 队列是否持久化 持久化队列在服务器重启之后可以恢复
		false, // 队列无消费者时是否自动删除
		false, // 是否专属与声明队列的连接
		false, // 是否等待服务器的响应
		nil,   // 其他参数
	)
	util.FailOnError(err, "声明队列失败")

	msgs, err := ch.Consume(
		q.Name, // queue
		"消费者",
		true, // 自动应答
		false,
		false, // RabbitMQ不支持这一项配置
		false,
		nil,
	)
	util.FailOnError(err, "注册消费者失败")

	var forever chan struct{}
	go func() {
		for msg := range msgs {
			slog.Infof("接收到消息: %v", string(msg.Body))
		}
	}()
	<-forever
}
