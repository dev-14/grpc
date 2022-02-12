package main

import (
	"log"

	"github.com/tutorialedge/go-grpc-tutorial/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from client side",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("error when responding: %s", err)

	}

	log.Printf("response from server: %s", response.Body)
}
