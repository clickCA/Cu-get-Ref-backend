PRO_PATH := .

PROTO_PATH := $(PRO_PATH)\coursemanagement.proto
OUTPUT_FOLDER := .\coursemanagement

generate_grpc_code:
	protoc --go_out=$(OUTPUT_FOLDER) --go_opt=paths=source_relative --go-grpc_out=$(OUTPUT_FOLDER) --go-grpc_opt=paths=source_relative $(PROTO_PATH)

dev:
	go run main.go

build:
	go build -o course-management-service.exe main.go