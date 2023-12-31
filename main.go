// Package main @Author: youngalone [2023/8/15]
package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"rabbitmq/config"
)

func main() {
	connection, err := amqp.Dial(config.URL)
	if err != nil {
		log.Fatalf("连接失败 %v", err)
	}
	fmt.Println(connection)
	n := 0
	fmt.Scanf("%d", &n)
	fmt.Println(n)
}
