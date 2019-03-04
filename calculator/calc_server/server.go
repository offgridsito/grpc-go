package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go/calculator/calcpb"
	"io"
	"log"
	"net"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calcpb.SumRequest) (*calcpb.SumResponse, error){
	fmt.Printf("Sum func was invoked with %v\n", req)

	firstAddend := req.FirstAddend
	secondAddend := req.SecondAddend

	sum := firstAddend + secondAddend

	res := &calcpb.SumResponse{
		SumResult: sum,
	}

	return res, nil
}

func(*server) PrimeNumberDecomp(req *calcpb.PrimeNumberDecompRequest, stream calcpb.CalculatorService_PrimeNumberDecompServer) error {
	fmt.Printf("PrimeNumberDecomp func was invoked with: %v\n", req)

	number := req.GetFirstNum()

	divisor := int64(2)

	for number > 1 {
		if number % divisor == 0 {
			stream.Send(&calcpb.PrimeNumberDecompResponse{
				PrimeFactor: divisor,
			})

			number = number / divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased to %v\n", divisor)
		}

	}

	return nil
}

func(*server) ComputeAverage(stream calcpb.CalculatorService_ComputeAverageServer) error {
	fmt.Printf("Received ComputeAverage RPC\n")

	sum := int32(0)
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			average := float64(sum) / float64(count)
			return stream.SendAndClose(&calcpb.ComputeAverageResponse{
				Average: average,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		sum += req.GetNumber()
		count++
	}
}

func main(){
	fmt.Println("Calculator Server...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}

	//create server
	s := grpc.NewServer()
	calcpb.RegisterCalculatorServiceServer(s, &server{})

	//bind port to server
	if err:= s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
}
