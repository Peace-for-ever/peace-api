package main

import (
	"context"
	"log"
	"net"

	"github.com/Peace-for-ever/peace-api/eventpb"
	"github.com/Peace-for-ever/peace-api/models"
	"google.golang.org/grpc"
)

type server struct {
	eventpb.UnimplementedEventServiceServer
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:8083")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	eventpb.RegisterEventServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
}

func (*server) EventService(ctx context.Context, in *eventpb.EventRequest) (*eventpb.EventResponse, error) {

	var events []*eventpb.Event

	for i := 0; i < int(in.GetAmount()); i++ {
		events = append(events, models.GenerateEvent())
	}

	res := eventpb.EventResponse{
		Events: events,
	}

	return &res, nil
}
