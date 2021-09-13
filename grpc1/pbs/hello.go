package pbs

import "golang.org/x/net/context"

type Server struct{}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	return &Message{Body: "Hello from go micro server!"}, nil
}
