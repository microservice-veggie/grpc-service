# Makefile

gen-proto:
	cd proto && protoc --go_out=../protogen/golang --go_opt=paths=source_relative \
	--go-grpc_out=../protogen/golang --go-grpc_opt=paths=source_relative \
	./**/*.proto
gen-proto-gateway:
	cd proto && protoc --go_out=../protogen/golang --go_opt=paths=source_relative \
	--go-grpc_out=../protogen/golang --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out ../protogen/golang --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
	./**/*.proto
start:
	go run main.go
run-proxy:
	go run proxy/proxy.go