package main

import (
	"context"
	"fmt"
	pb "go_grpc/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to the gRPC server
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:8080", opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPersonServiceClient(conn)

	// Timeout for context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// Example 1: Create a new person
	fmt.Println("Creating a new person...")
	createReq := &pb.CreatePersonRequest{
		Name:        "Swapnil Kamble",
		Email:       "swapnil.kamble2701@outlook.com",
		PhoneNumber: "+91 9028864181",
	}
	createRes, err := client.Create(ctx, createReq)
	if err != nil {
		log.Fatalf("Error during Create: %v", err)
	}
	fmt.Printf("Person created: %+v\n", createRes)

	// Example 2: Read the created person by ID
	fmt.Println("Reading the person by ID...")
	readReq := &pb.SinglePersonRequest{
		Id: createRes.GetId(),
	}
	readRes, err := client.Read(ctx, readReq)
	if err != nil {
		log.Fatalf("Error during Read: %v", err)
	}
	fmt.Printf("Person details: %+v\n", readRes)

	// Example 3: Update the person's details
	fmt.Println("Updating the person's details...")
	updateReq := &pb.UpdatePersonRequest{
		Id:          createRes.GetId(),
		Name:        "Luke Skywalker",
		Email:       "luke.skywalker@codeheim.io",
		PhoneNumber: "987-654-3210",
	}
	updateRes, err := client.Update(ctx, updateReq)
	if err != nil {
		log.Fatalf("Error during Update: %v", err)
	}
	fmt.Printf("Update response: %s\n", updateRes.GetResponse())

	// Example 4: Delete the person by ID
	fmt.Println("Deleting the person by ID...")
	deleteReq := &pb.SinglePersonRequest{
		Id: createRes.GetId(),
	}
	deleteRes, err := client.Delete(ctx, deleteReq)
	if err != nil {
		log.Fatalf("Error during Delete: %v", err)
	}
	fmt.Printf("Delete response: %s\n", deleteRes.GetResponse())
}