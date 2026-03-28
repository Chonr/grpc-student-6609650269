package main

import (
	"context"
	"log"
	"net"

	pb "grpc-student-6609650269/studentpb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStudentServiceServer
}

func (s *server) GetStudent(ctx context.Context, req *pb.StudentRequest) (*pb.StudentResponse, error) {

	log.Printf("Received request for student ID: %d", req.Id)

	// Mock data
	return &pb.StudentResponse{
		Id:    req.Id,
		Name:  "Alice Johnson",
		Major: "Computer Science",
		Email: "alice@university.com",
		Phone: "066-096-0269",
	}, nil
}

func (s *server) ListStudents(ctx context.Context, req *pb.Empty) (*pb.StudentListResponse, error) {
	log.Println("Received request for student list")

	students := []*pb.StudentResponse{
		{Id: 660, Name: "Chon N", Major: "Computer Science", Email: "chon.s@mail.com", Phone: "012-345-6789"},
		{Id: 269, Name: "Sin M", Major: "Computer Science", Email: "sin.d@mail.com", Phone: "098-765-4321"},
	}

	return &pb.StudentListResponse{Student: students}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterStudentServiceServer(grpcServer, &server{})

	log.Println("gRPC Server running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
