run:
	go run cmd/main.go

proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/user.proto

mock:
	mockgen -source=pkg/pb/user.pb.go -destination=pkg/client/mocks/mock_userclient.go -package=mocks UserServiceClient
