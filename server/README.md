1. Install Protobuf Compiler (protoc):

If you haven't already, download and install the Protocol Buffers compiler (protoc) for your platform from the official GitHub repository: https://github.com/protocolbuffers/protobuf/releases

2. Install the Go Protocol Buffers Plugin (protoc-gen-go):

You can install the Go Protocol Buffers plugin using the go get command:

````
go install google.golang.org/protobuf/cmd/protoc-gen-go

protoc --go_out=. --go-grpc_out=. proto/task.proto
````

This command generates Go code for both Protocol Buffers messages and gRPC service definitions. The --go_out flag generates code for Protocol Buffers, and the --go-grpc_out flag generates code for gRPC.

Make sure to navigate to the directory containing your .proto file before running the protoc command.

Generated Code:

After running the protoc command, you will find the generated Go files in the current directory. The generated files will include:

_your.pb.go: Contains the Protocol Buffers message definitions.
_your_grpc.pb.go: Contains the gRPC service and client definitions.
You can import and use these generated files in your Go code to work with the gRPC service and Protocol Buffers messages defined in your .proto file.
