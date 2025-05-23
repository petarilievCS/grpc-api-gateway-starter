PROTO_DIR=proto
OUT_DIR=gen

generate:
	protoc \
  --proto_path=proto \
  --go_out=. \
  --go-grpc_out=. \
  proto/user.proto proto/order.proto proto/gateway.proto