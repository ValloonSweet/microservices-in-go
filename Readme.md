# Installing protoc 
```
go install google.golang.org/protobuf/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

# Build protobuf
Move to logs
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative logs.proto
```