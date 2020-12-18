package grpcServer

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"

    pb "github.com/EDDYCJY/go-grpc-example/proto"
)

type EchoServer struct{}

func (server *EchoServer) Echo(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
    return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9001"

// func main() {
//     server := grpc.NewServer()
//     pb.RegisterSearchServiceServer(server, &SearchService{})

//     lis, err := net.Listen("tcp", ":"+PORT)
//     if err != nil {
//         log.Fatalf("net.Listen err: %v", err)
//     }

//     server.Serve(lis)
// }