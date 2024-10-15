# Simple_Go_gRPC
What I learned: 
+ The fundamentals of gRPC and its architecture 
+ How to define service interfaces using Protocol Buffers (protobuf) 
+ Implementing unary and server-streaming RPCs 
+ Understanding the basics of gRPC clients and servers

```bash
      $ go install google.golang.org/grpc@latest
      $ go install google.golang.org/protobuf@latest
```
```

      $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
      $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
# Regenerate gRPC code
    $ protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative    proto/*.proto
