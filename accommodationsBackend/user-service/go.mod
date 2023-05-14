module accommodationsBackend/user-service

go 1.20

replace accommodationsBackend/common => ../common

replace accommodationsBackend/api_gateway => ../api_gateway

require (
	accommodationsBackend/api_gateway v0.0.0-00010101000000-000000000000
	accommodationsBackend/common v0.0.0-00010101000000-000000000000
	go.mongodb.org/mongo-driver v1.11.6
	golang.org/x/crypto v0.8.0
	google.golang.org/grpc v1.54.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2 // indirect
	github.com/klauspost/compress v1.16.5 // indirect
	github.com/montanaflynn/stats v0.7.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
