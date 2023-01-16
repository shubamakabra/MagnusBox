.PHONY: proto
.SILENT: proto

proto:
	protoc \
  --go_out=. --go_opt=module=github.com/shubamakabra/MagnusBox \
  --go-grpc_out=. --go-grpc_opt=module=github.com/shubamakabra/MagnusBox \
  ./proto/**/*.proto

	go mod tidy