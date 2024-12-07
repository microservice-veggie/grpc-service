package routes

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/microservice-veggie/grpc-service/common/utils"
	gwauth "github.com/microservice-veggie/grpc-service/protogen/golang/auth" // Update
	"google.golang.org/grpc"
)

func AuthRoutes(ctx context.Context, mux *runtime.ServeMux, opts []grpc.DialOption) {
	domain := utils.GoDoEnvVariable("DOMAIN")
	portAuth := fmt.Sprintf("%s:8081", domain)
	var (
		// gRPC server endpoint
		grpcServerEndpoint = flag.String("grpc-server-endpoint-auth", portAuth, "gRPC server endpoint auth service")
	)
	//Login
	err := gwauth.RegisterLoginAPIHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
