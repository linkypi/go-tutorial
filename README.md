
``` shell
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    /path/to/xx.proto
    
    
protoc --go_out=. --go-grpc_out=. hello.proto
```