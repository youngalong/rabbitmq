// Package util @Author: youngalone [2023/8/15]
package util

import (
	"github.com/gookit/slog"
)

func FailOnError(err error, msg string) {
	if err != nil {
		slog.Panicf("%s: %s", msg, err)
	}
}
