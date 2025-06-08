package main

import (
	"context"
	"fmt"
	"log"
	"log-analyzer/middleware"
	"net"
	"time"

	"google.golang.org/grpc"
)

type HelloRequest struct {
	Name string
}

type HelloReply struct {
	Message string
}

type greeterServer struct{}

func (s *greeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	time.Sleep(150 * time.Millisecond) // 擬似的な処理
	return &HelloReply{Message: "Hello, " + req.Name}, nil
}

func registerGreeterServer(s *grpc.Server, srv GreeterServer) {
	type wrapper struct {
		GreeterServer
	}

	grpc.RegisterService(&grpc.ServiceDesc{
		ServiceName: "example.Greeter",
		HandlerType: (*GreeterServer)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "SayHello",
				Handler: func(srv interface{}, ctx context.Context, dec func(interface{}) error, _ grpc.UnaryServerInterceptor) (interface{}, error) {
					in := new(HelloRequest)
					if err := dec(in); err != nil {
						return nil, err
					}
					return srv.(GreeterServer).SayHello(ctx, in)
				},
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "",
	}, wrapper{srv})
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.UnaryServerInterceptor()),
	)
	registerGreeterServer(s, &greeterServer{})
	fmt.Println("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
