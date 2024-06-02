package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	user "shopito/pkg/protobuf/users"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(":3002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := user.NewUserServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.CreateUser(ctx, &user.CreateUserRequest{
		Name:     "Danil",
		Email:    "danil.li24@mail.ru",
		Password: "12345",
	})
	if err != nil {
		log.Fatalf("could not send request: %v", err)
	}
	fmt.Println(r.Id, r.Success)

	// r, err := c.GetUsers(ctx, &user.GetUsersRequest{})
	// if err != nil {
	// 	log.Fatalf("could not send request: %v", err)
	// }
	// users := r.GetUsers()
	// for _, user := range users {
	// 	fmt.Println(user.Id, user.Email, user.Name)
	// }
	// log.Printf("Get Users: %s, %v, %v, %v", r.GetEmail(), r.GetName(), r.GetPassword(), r.GetIsAdmin())
}
