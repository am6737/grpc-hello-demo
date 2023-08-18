package main

import (
	"context"
	"flag"
	"fmt"
	helloworld "github.com/am6737/grpc-hello-demo/pb"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	helloworld.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	message := fmt.Sprintf("Hello, %s!", req.Name)
	return &helloworld.HelloResponse{Message: message}, nil
}

var (
	Endpoint string
)

func init() {
	flag.StringVar(&Endpoint, "endpoint", "/tmp/hello.sock", "socket addr")
	flag.Parse()
}

func main() {

	listener, err := net.Listen("unix", Endpoint)
	if err != nil {
		fmt.Printf("Failed to listen: %v\n", err)
		return
	}

	srv := grpc.NewServer()
	helloworld.RegisterHelloServiceServer(srv, &server{})

	fmt.Println("Server is listening on UNIX socket:", Endpoint)
	err = srv.Serve(listener)
	if err != nil {
		fmt.Printf("Failed to serve: %v\n", err)
		return
	}
}
