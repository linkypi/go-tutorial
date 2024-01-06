package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"strconv"
	"testing"

	pb "hiraeth.com/tutorial/grpc/proto"
)

func TestClient1(t *testing.T) {
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	serverClient := pb.NewGreeterClient(conn)
	response, err := serverClient.SayHello(context.Background(), &pb.HelloRequest{Name: "Hello"})
	fmt.Println("say hello response", response)

	// client stream
	quit := make(chan struct{})
	streamPing, err := serverClient.ClientStreamPing(context.Background())
	go func() {
		count := 0
		// 往流中发 10 个请求
		for {
			count++
			err2 := streamPing.Send(&pb.PingRequest{Id: "client msg" + strconv.Itoa(count)})
			if err2 != nil {
				panic(err2)
			}

			if count > 10 {
				quit <- struct{}{}
				break
			}
		}
	}()

	<-quit

	reply, err := streamPing.CloseAndRecv()
	fmt.Println("recv: ", reply.GetMessage())

	// server stream
	serverStream, err := serverClient.ServerStreamPing(context.Background(), &pb.PingRequest{Id: "server msg1"})
	if err != nil {
		panic(err)
	}
	for {
		req, err := serverStream.Recv()
		if err != nil {
			if err == io.EOF {
				// 客户端已接收完成
				fmt.Println("client has received all msg")
				err := serverStream.CloseSend()
				if err != nil {
					panic(err)
				}
				break
			}
			panic(err)
		}

		fmt.Println("recv: ", req.GetMessage())
		if req.GetMessage() == "server msg10" {
			break
		}
	}
}
