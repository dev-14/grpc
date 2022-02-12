package chat

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) sayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("received message body from client: %s", message.Body)
	return &Message{Body: "Hello from the server"}, nil
}
