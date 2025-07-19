package middleware

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type ctxKey string

const TraceIDKey ctxKey = "trace_id"

func TraceIDFromContext(ctx context.Context) string {
	if v := ctx.Value(TraceIDKey); v != nil {
		if id, ok := v.(string); ok {
			return id
		}
	}
	return "unknown"
}

func SimpleTraceIDInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		traceID := uuid.New().String()
		ctx = context.WithValue(ctx, TraceIDKey, traceID)
		return handler(ctx, req)

	}
}
