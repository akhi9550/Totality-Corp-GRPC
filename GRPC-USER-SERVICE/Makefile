run:
	go run cmd/main.go

proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/user.proto

mock:
	mockgen -source pkg\repository\interface\user.go -destination pkg\repository\mock\user_mock.go -package mock
	mockgen -source pkg\usecase\interface\user.go -destination pkg\usecase\mock\user_mock.go -package mock