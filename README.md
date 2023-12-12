# Protobuf Locations

```shell
protoc --proto_path=. \
  --go_out=./locations --go_opt=paths=source_relative \
  --go-grpc_out=./locations --go-grpc_opt=paths=source_relative \
  locations.proto

```