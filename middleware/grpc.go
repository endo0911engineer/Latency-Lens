package middleware

import (
	"context"
	"log-analyzer/collector"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

// UnaryServerInterceptor returns a new unary server interceptor that records latency.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		start := time.Now()
		resp, err = handler(ctx, req)
		elapsed := time.Since(start)

		label := info.FullMethod
		if p, ok := peer.FromContext(ctx); ok {
			label = p.Addr.String() + " " + info.FullMethod
		}

		collector.Record(label, elapsed)
		return resp, err
	}
}
