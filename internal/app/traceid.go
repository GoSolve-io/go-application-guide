package app

import (
	"context"
)

type ctxTraceIDKeyType uint32

const (
	ctxTraceIDKeyKey ctxTraceIDKeyType = iota
)

// TraceIDFromCtx returns trace id from context.
// If id is not present in context, returns empty string.
func TraceIDFromCtx(ctx context.Context) string {
	if id, ok := ctx.Value(ctxTraceIDKeyKey).(string); ok {
		return id
	}
	return ""
}

// CtxWithTraceID returns new context with trace id.
func CtxWithTraceID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, ctxTraceIDKeyKey, id)
}
