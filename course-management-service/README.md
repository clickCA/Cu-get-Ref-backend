## Get Started

Generate the proto files

```pwsh
protoc --go_out=coursemanagement --go_opt=paths=source_relative --go-grpc_out=coursemanagement --go-grpc_opt=paths=source_relative .\coursemanagement.proto
```
