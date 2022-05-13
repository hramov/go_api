package api

import (
	"api/src/core/logger"
	pb "api/src/proto"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

var conn net.Listener
var server *grpc.Server

func initGrpcApi() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("GRPC_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(grpcServer, &Server{})
	conn = lis
	server = grpcServer
}

func createGRPCServer() {
	if server == nil || conn == nil {
		logger.Error("GRPC not initialized")
	}
	if err := server.Serve(conn); err != nil {
		logger.Error("Cannot start GRPC server: " + err.Error())
	}
}

func (s *Server) SayHello(ctx context.Context, helloRequest *pb.HelloRequest) (*pb.HelloReply, error) {
	greeting := fmt.Sprintf("Hello, %s", helloRequest.Name)
	return &pb.HelloReply{Message: greeting}, nil
}

func (s *Server) SayMultipleHello(req *pb.HelloRequest, stream pb.Greeter_SayMultipleHelloServer) error {

	for i := 0; i < 10; i++ {
		message := &pb.HelloReply{
			Message: fmt.Sprintf("Hello, %s | %d\n", req.Name, i),
		}
		if err := stream.Send(message); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) SayHelloToMultiplePeople(stream pb.Greeter_SayHelloToMultiplePeopleServer) error {
	name := ""
	for {
		request, err := stream.Recv()
		if request != nil {
			name = request.Name
		}
		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloReply{
				Message: fmt.Sprintf("That's all, folks. Last name was %s\n", name),
			})
		}
		if err != nil {
			return err
		}
		log.Println(request.Name)
	}
}

func (s *Server) SayHelloToEachPerson(stream pb.Greeter_SayHelloToEachPersonServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		name := in.Name

		log.Println(name)

		if err := stream.Send(&pb.HelloReply{
			Message: fmt.Sprintf("Hello, %s\n", name),
		}); err != nil {
			return err
		}
	}
}
