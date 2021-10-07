package main

import (
	"context"
	"log"
	"net"

	"github.com/service_content/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {

	listen, err := net.Listen("tcp", ":4040")

	if err != nil {
		panic(err)
	}

	// ctx := context.Background()
	srv := grpc.NewServer()
	model.RegisterEventStoreServer(srv, &server{})
	reflection.Register(srv)

	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)

	// go func() {
	// 	for range c {
	// 		log.Println("shutting down server gRPC")

	// 		srv.GracefulStop()

	// 		<-ctx.Done()
	// 	}
	// }()

	if err := srv.Serve(listen); err != nil {
		panic(err)
	}
}

func (s *server) GetEvents(ctx context.Context, eventData *model.EventFilter) (*model.EventResponse, error) {
	return &model.EventResponse{}, nil
}

func (*server) CreateEvent(ctx context.Context, evParams *model.EventParam) (*model.ResponseParam, error) {
	// conn := database.NewConnection()
	// repo := database.NewEventConnection(conn)
	// use := repository.NewEventEntity(repo)
	// if err := use.CreateEvent(ctx, evParams); err != nil {
	// 	return nil, err
	// }
	log.Print(evParams)

	//and publish to nats
	return &model.ResponseParam{}, nil
}
