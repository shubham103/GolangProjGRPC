package main

import (
        "context"
        "log"
 
        "google.golang.org/grpc"

        pb "grpc-hello-world/helloworld"
)

const (
        address  = "localhost:50051"
)

func main(){
        conn, err := grpc.Dial(address, grpc.WithInsecure())
        if err != nil {
                log.Fatalf("Failed to connect: %v", err)
        }
        defer conn.Close()

        c:= pb.NewGreeterClient(conn)
        name := "John"

        r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
        if err != nil {
                log.Fatalf("Error calling SayHello: %v", err)
        }
        log.Printf("Response from server %s", r.Message)
}
