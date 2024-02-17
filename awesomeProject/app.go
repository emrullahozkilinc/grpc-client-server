package main

import (
	"awesomeProject/protobuff"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	protobuff.RegisterTimeServiceServer(server, &srv{})
	log.Printf("Starting grpc server")
	err = server.Serve(lis)
	log.Fatal(err)
}

type srv struct{}

func (s *srv) Now(ctx context.Context, req *protobuff.NowRequest) (*protobuff.TimeUpdate, error) {
	return &protobuff.TimeUpdate{Time: &protobuff.Time{
		Value: time.Now().String(),
	}}, nil
}

func (s *srv) Stream(request *protobuff.TimeStreamRequest, server protobuff.TimeService_StreamServer) error {
	deadline := time.Now().Add(time.Duration(request.Length) * time.Second)
	for !time.Now().After(deadline) {
		time.Sleep(time.Millisecond * 300)

		err := server.Send(&protobuff.TimeUpdate{Time: &protobuff.Time{
			Value: time.Now().String(),
		}})
		if err != nil {
			return err
		}
	}
	return nil
}
