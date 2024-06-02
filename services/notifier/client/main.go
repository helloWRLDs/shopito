package main

import (
	"context"
	"fmt"
	"log"
	"shopito/services/notifier/protobuf"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:3003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := protobuf.NewNotifierServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := c.SendEmail(ctx, &protobuf.SendEmailRequest{
		To:      "danil.li24x@gmail.com",
		Subject: "Some subject",
		Body:    "some body wants to know",
	})
	if err != nil {
		log.Fatalf("could not send request: %v", err)
	}
	fmt.Println(r.Success)
}
