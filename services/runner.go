package services

import (
	"context"
	"log"
	"net"
	"net/http"
	"users/services/pb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Params struct {
	GrpcPort int
	RestPort int
	Postgres IP
}

func runGrpc() {
	lis, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func runRest() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterUsersHandlerFromEndpoint(ctx, mux, "localhost:9080", opts)
	if err != nil {
		panic(err)
	}

	log.Printf("http server listening at 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
