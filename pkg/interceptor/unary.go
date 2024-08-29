package interceptor

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func UnaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	slog.Debug("request", "method", info.FullMethod)

	resp, err = handler(ctx, req)
	if err != nil {
		st := status.Convert(err)
		slog.With("error", st.Message()).Error("")
	} else {
		slog.Debug("success")
	}

	return resp, err
}
