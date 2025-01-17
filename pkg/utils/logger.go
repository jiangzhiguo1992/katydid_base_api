package utils

import (
	"context"
	"log/slog"
)

func Init() {
	// TODO:GG 初始化配置
}

func Debug(msg string, args ...any) {
	slog.Debug(msg, args...)
}

func DebugContext(ctx context.Context, msg string, args ...any) {
	slog.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	slog.Info(msg, args...)
}

func InfoContext(ctx context.Context, msg string, args ...any) {
	slog.InfoContext(ctx, msg, args...)
}

func Warn(msg string, args ...any) {
	slog.Warn(msg, args...)
}

func WarnContext(ctx context.Context, msg string, args ...any) {
	slog.WarnContext(ctx, msg, args...)
}

func Error(msg string, args ...any) {
	slog.Error(msg, args...)
}

func ErrorContext(ctx context.Context, msg string, args ...any) {
	slog.ErrorContext(ctx, msg, args...)
}
