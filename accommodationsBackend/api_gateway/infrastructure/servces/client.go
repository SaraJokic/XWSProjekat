package servces

import (
	"accommodationsBackend/common/proto/accommodation_service"
	auth_service "accommodationsBackend/common/proto/auth-service"
	"accommodationsBackend/common/proto/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewUserClient(address string) user_service.UserServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to User service: %v", err)
	}
	return user_service.NewUserServiceClient(conn)
}

func NewAccommodationClient(address string) accommodation_service.AccommodationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Accommodation service: %v", err)
	}
	return accommodation_service.NewAccommodationServiceClient(conn)
}

func NewAuthClient(address string) auth_service.AuthServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Accommodation service: %v", err)
	}
	return auth_service.NewAuthServiceClient(conn)
}
func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
