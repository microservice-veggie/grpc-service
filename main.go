package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"                       // Update
	gwauth "github.com/microservice-veggie/grpc-service/protogen/golang/auth" // Update
	"google.golang.org/grpc"
)

func main() {
	// creating mux for gRPC gateway. This will multiplex or route request different gRPC service
	mux := runtime.NewServeMux()

	// Creating a normal HTTP server
	server := http.Server{
		Handler: mux,
	}
	// setting up a dail up for gRPC service by specifying endpoint/target url
	err := gwauth.RegisterLoginAPIHandlerFromEndpoint(context.Background(), mux, "auth-microservice-veggie-auth-service-1-1:8082", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal(err)
	}
	// creating a listener for server
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server started at :8081")
	// start server
	err = server.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
