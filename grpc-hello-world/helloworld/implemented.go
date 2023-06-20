package helloworld

import context "context" 

type Server struct{
        UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error){
        return &HelloResponse{Message: "Hello, "+ in.GetName()},nil
}
