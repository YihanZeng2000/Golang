package main

import (
	"add_server/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	proto.UnimplementedCalcServiceServer
}

func (s server) Add(ctx context.Context, in *proto.AddRequest) (*proto.AddResponse, error) {
	sum := int64(in.GetX()) + int64(in.GetY())
	return &proto.AddResponse{Result: sum}, nil
}

func main() {
	// add rpc server
	l, err := net.Listen("tcp", ":8973")
	if err != nil {
		log.Fatalf("net.Listen failed, error:%v\n", err)
		return
	}
	s := grpc.NewServer()
	// 注册
	proto.RegisterCalcServiceServer(s, &server{})
	// 启动该服务
	err = s.Serve(l)
	if err != nil {
		log.Fatalf("s.Serve failed, error:%v\n", err)
		return
	}
}
