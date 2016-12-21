// +build server

package main

import (
	"log"
	"net"

	pb "github.com/arith-abe/grpcexample/grpcexample"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type handler struct{}

func (h *handler) GetPerson(ctx context.Context, req *pb.Request) (*pb.Person, error) {
	return &pb.Person{Name: "Hi!"}, nil
}

func (h *handler) ListPeople(req *pb.Request, stream pb.GRPCExample_ListPeopleServer) error {
	for i := 0; i < 5; i++ {
		if err := stream.Send(&pb.Person{Id: int32(i)}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	h := &handler{}
	pb.RegisterGRPCExampleServer(s, h)
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
