package rpc_5

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"time"
)

type HelloServiceImpl struct {
}

func (s *HelloServiceImpl) Channel(server HelloService_ChannelServer) error {
	for {
		req, err := server.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		reply := &String{
			Value: "Hello:" + req.GetValue(),
		}
		log.Println(reply)
		//err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func (s *HelloServiceImpl) Hello(ctx context.Context, args *String) (*String, error) {
	reply := &String{
		Value: "Hello:" + args.GetValue(),
	}
	return reply, nil
}

func GrpcServer() {
	grpcServer := grpc.NewServer()
	RegisterHelloServiceServer(grpcServer, &HelloServiceImpl{})
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer.Serve(lis)
}

func GrpcClient() {
	conn, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewHelloServiceClient(conn)
	reply, err := c.Hello(context.Background(), &String{Value: "World"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(reply)
}

func GrpcDemoMain() {
	go GrpcServer()
	time.Sleep(5 * time.Second)
	GrpcClient()
}
