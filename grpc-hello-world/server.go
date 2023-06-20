package main

import (
        "log"
        "net"

        "google.golang.org/grpc"

        pb "grpc-hello-world/helloworld"
)

const (
        port = ":50051"
)

func main() {
        log.Println("gRPC server tutorial in Go")

        lis, err := net.Listen("tcp", port)
        if err != nil {
                panic(err)
        }


        s := grpc.NewServer()
        pb.RegisterGreeterServer(s, &pb.Server{})
        if err := s.Serve(lis); err != nil {
                log.Fatalf("failed to serve: %v", err)
        }
}
