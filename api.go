package logger

import (
	"log/slog"
)

type SlogHandler interface {
	slog.Handler
	SetLevel(level Level)
	SetAddSource(b bool)
}

// NOTE: 返回值中的 SlogHandler 可选用于进一步配置 Handler,从后台动态修改调试日志级别
func New(opts ...*Option) (*slog.Logger, SlogHandler) {
	h := NewFastTextHandler(opts...)
	return slog.New(h), h
}

func SetDefault(opts ...*Option) SlogHandler {
	sl, handler := New(opts...)
	slog.SetDefault(sl)
	return handler
}
