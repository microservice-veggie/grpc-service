package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/microservice-veggie/grpc-service/common/utils"
	api "github.com/microservice-veggie/grpc-service/routes" // Update
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

func run() error {
	domain := utils.GoDoEnvVariable("DOMAIN")
	endPoint := fmt.Sprintf("%s:8080", domain)
	//init context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	//manage routes
	api.AuthRoutes(ctx, mux, opts)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Printf("server is running:%s", endPoint)
	return http.ListenAndServe(endPoint, mux)
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}
