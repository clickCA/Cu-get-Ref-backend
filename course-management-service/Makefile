generate_grpc_code:
	protoc --go_out=coursemanagement --go_opt=paths=source_relative --go-grpc_out=coursemanagement --go-grpc_opt=paths=source_relative .\coursemanagement protorpc_opt=paths=source_relative .\coursemanagement.proto

dev:
	go run main.go