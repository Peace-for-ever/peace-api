package main

import (
	"context"
	"log"
	"time"

	"github.com/Peace-for-ever/peace-api/eventpb"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("peaceland.thomaslacaze.fr:8083", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := eventpb.NewEventServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.EventService(ctx, &eventpb.EventRequest{Amount: 10})
	if err != nil {
		log.Fatalf("could not use EventService: %v", err)
	}
	log.Printf("Events: %s", r.GetEvents()[0])
}
