package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go/calculator/calcpb"
	"io"
	"log"
)

func main(){
	fmt.Println("Hello I am calculator client.")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect: %v", err)
	}

	defer cc.Close()

	// create client

	c := calcpb.NewCalculatorServiceClient(cc)

	//doUnary(c)
	//doServerStreaming(c)

	doClientStreaming(c)

}

func doUnary(c calcpb.CalculatorServiceClient){
	fmt.Println("Starting Sum Unary RPC...")

	req := &calcpb.SumRequest{
		FirstAddend: 3,
		SecondAddend: 10,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}

	log.Println("Response from Sum", res.SumResult)
}

func doServerStreaming(c calcpb.CalculatorServiceClient){
	fmt.Println("Starting to do a PrimeDecomp Server Streaming RPC...")

	req := &calcpb.PrimeNumberDecompRequest{
		FirstNum: 958640834830,
	}

	stream, err := c.PrimeNumberDecomp(context.Background(), req)
	if err != nil {
		log.Fatalf("error while caling PrimeDecomp RPC: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}
}

func doClientStreaming(c calcpb.CalculatorServiceClient){
	fmt.Println("Starting to do a ComputeAvg client streaming RPC\n")

	stream, err := c.ComputeAverage(context.Background())

	if err != nil {
		log.Fatalf("Error while opening stream: %v", err)
	}

	numbers := []int32{3,5,7,6,11,23}

	for _, number := range numbers {
		fmt.Printf("Sending number: %v\n", number)
		stream.Send(&calcpb.ComputeAverageRequest{
			Number: number,
		})
	}


	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}

	fmt.Printf("The Average is: %v\n", res.GetAverage())

}