package main

//go:generate go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
//go:generate go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
//go:generate protoc --go_out=./ --go-grpc_out=./ --proto_path=../proto ../proto/raft.proto
import (
	"flag"
	"fmt"
	"net/http"

	raft "github.com/FrancescoLuzzi/raft/raft"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpc_port = flag.Int("grpc_port", 50051, "The grpc server port")
	http_port = flag.Int("http_port", 80, "The http server port")
	grpc_dst  = flag.String("grpc_dst", "http://localhost:50051", "Destination grpc server to call")
)

// server is used to implement helloworld.GreeterServer.

func rootHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.NewClient(*grpc_dst, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		http.Error(w, fmt.Sprintf("Something went wrong while connecting: %v", err), 500)
		return
	}
	client := raft.NewRaftModuleServiceClient(conn)
	// TODO:
	// - create server to handle other servers Requests
	// - create client to handle requesting to other servers
	// - create interface that implements the server interface (called by the server) and uses a client
	// - (TO UNDERSTAND) abstract the call to clients using ids:
	//	 the server impl hold the connections and ip informations of other nodes, it gives the ability to read all ids registered
	//   the raft modules interact with this client by using a clientId instead of connections and ips
}
