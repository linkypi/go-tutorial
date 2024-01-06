package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "hiraeth.com/tutorial/grpc/proto"
	"io"
	"net"
	"strconv"
	"testing"
)

type MyServer struct {
	pb.GreeterServer
}

func (s *MyServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("get request: ", req.Name)
	return &pb.HelloReply{
		Message: "Hello",
		Code:    200,
	}, nil
}

func (c *MyServer) ClientStreamPing(stream pb.Greeter_ClientStreamPingServer) error {
	count := 0
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Println("client stream closed")
				return nil
			}
			return err
		}

		fmt.Println("recv: ", req.GetId())
		count++
		if count > 10 {
			err := stream.SendAndClose(&pb.PingReply{Message: "client all msg has been received."})
			if err != nil {
				return err
			}
		}

	}
}

func (*MyServer) ServerStreamPing(req *pb.PingRequest, stream pb.Greeter_ServerStreamPingServer) error {
	count := 0
	for {
		count++
		err := stream.Send(&pb.PingReply{Message: "server msg" + strconv.Itoa(count)})
		if err != nil {
			return err
		}
		if count > 10 {
			return nil
		}
	}
}

func TestServer1(t *testing.T) {
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &MyServer{})
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
	server.GracefulStop()
}
