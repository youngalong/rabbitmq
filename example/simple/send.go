// Package main @Author: youngalone [2023/8/15]
package main

import (
	"context"
	"github.com/gookit/slog"
	amqp "github.com/rabbitmq/amqp091-go"
	"rabbitmq/config"
	"rabbitmq/util"
	"time"
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

	// 声明context 防止超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	msg := "这是一条数据"
	err = ch.PublishWithContext(
		ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
	util.FailOnError(err, "发布消息失败")
	slog.Info("发送完毕")
}
