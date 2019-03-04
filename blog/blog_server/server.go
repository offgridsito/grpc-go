package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"grpc-go/blog/blogpb"
	"log"
	"net"
	"os"
	"os/signal"
)

var collection *mongo.Collection

type server struct{

}

type blogItem struct {
	ID 			objectid.ObjectID		`bson:"_id,omitempty"`
	AuthorID 	string		`bson:"author_id"`
	Content		string		`bson:"content"`
	Title 		string		`bson:"title"`
}

func main(){
	// if we crash go code, we get file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)


	fmt.Println("Blog Server Started")

	// connect to mongo db
	client, err := NewClient(options.Client().ApplyURI("mongodb://foo:bar@localhost:27017"))
	if err != nil { return err }
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil { return err }

	//from client connect to db and use blog "table"
	collection = client.Database("mydb").Collection("blog")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
	log.Fatal("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}

	//create server
	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, &server{})

	go func(){
		fmt.Println("Starting blog server...")

		if err:= s.Serve(lis); err != nil {
			log.Fatal("Failed to serve: %v", err)
		}
	}()

	//Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until signal is received
	<-ch
	fmt.Println("\nStopping the server")
	s.Stop()
	fmt.Println("Closing the listener...")
	lis.Close()
	fmt.Println("closing mongodb conn")
	client.Disconnect(ctx)
	fmt.Println("Closing program...")

}
