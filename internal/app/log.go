package app

import (
	"context"

	"github.com/sirupsen/logrus"
)

type ctxLoggerKeyType uint32

const (
	ctxLoggerKey ctxLoggerKeyType = iota
)

// AugmentLogFromCtx augments logger with data from context.
func AugmentLogFromCtx(ctx context.Context, l logrus.FieldLogger) logrus.FieldLogger {
	data := logDataFromCtx(ctx)
	for k, v := range data {
		l = l.WithField(k, v)
	}
	return l
}

// CtxWithLogField adds extra log information to context.
func CtxWithLogField(ctx context.Context, key string, value string) context.Context {
	return ctxWithLogData(ctx, map[string]string{key: value})
}

func logDataFromCtx(ctx context.Context) map[string]string {
	if data, ok := ctx.Value(ctxLoggerKey).(map[string]string); ok {
		return data
	}
	return make(map[string]string)
}

func ctxWithLogData(ctx context.Context, data map[string]string) context.Context {
	if data == nil {
		return ctx
	}
	ctxData := logDataFromCtx(ctx)
	for k, v := range data {
		ctxData[k] = v
	}

	return context.WithValue(ctx, ctxLoggerKey, ctxData)
}
