package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/pkg/pubsub"
	"google.golang.org/grpc"
	"net"
	"strings"
	"time"
)

type PubsubService struct {
	pub *pubsub.Publisher
}

func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (p *PubsubService) Publish(ctx context.Context, arg *String) (*String, error) {
	p.pub.Publish(arg.GetValue())
	return &String{}, nil
}

func (p *PubsubService) Subscribe(arg *String, stream PubsubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&String{Value: v.(string)}); err != nil {
			return err
		}
	}
	return nil
}

func (p *PubsubService) mustEmbedUnimplementedPubsubServiceServer() {}

func main() {
	grpcServer := grpc.NewServer()
	RegisterPubsubServiceServer(grpcServer, NewPubsubService())
	lis, err := net.Listen("tcp", ":9091")
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	grpcServer.Serve(lis)
}
