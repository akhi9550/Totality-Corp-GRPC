run:
	go run cmd/main.go

proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/user.proto

mock:
	mockgen -source pkg\client\interface\user.go -destination pkg\client\mock\handler_mock.go -package mock
