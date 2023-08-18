package main

import (
	"context"
	"flag"
	"fmt"
	helloworld "github.com/am6737/grpc-hello-demo/pb"
	"google.golang.org/grpc"
)

var (
	Endpoint string
)

func init() {
	flag.StringVar(&Endpoint, "endpoint", "/tmp/hello.sock", "socket addr")
	flag.Parse()
}

func main() {
	conn, err := grpc.Dial("unix://"+Endpoint, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Failed to connect to socket: %v\n", err)
		return
	}
	defer conn.Close()

	client := helloworld.NewHelloServiceClient(conn)
	req := &helloworld.HelloRequest{Name: "John"}
	resp, err := client.SayHello(context.Background(), req)
	if err != nil {
		fmt.Printf("SayHello RPC failed: %v\n", err)
		return
	}

	fmt.Println("Response:", resp.Message)
}
