// Package config @Author: youngalone [2023/8/15]
package config

import (
	"fmt"
	"github.com/gookit/slog"
	"github.com/spf13/viper"
)

var URL string

func init() {
	viper.SetConfigFile("../../config/settings.yml")
	if err := viper.ReadInConfig(); err != nil {
		slog.Errorf("加载配置信息失败 %v", err)
	} else {
		slog.Debug("加载配置信息成功")
	}
	URL = fmt.Sprintf("amqp://%s:%s@%s:%s",
		viper.GetString("settings.rabbitMQ.username"),
		viper.GetString("settings.rabbitMQ.password"),
		viper.GetString("settings.rabbitMQ.host"),
		viper.GetString("settings.rabbitMQ.port"),
	)
}
