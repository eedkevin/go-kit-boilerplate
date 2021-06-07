package grpc

import (
	"context"
	pb "github.com/eedkevin/go-kit-boilerplate/echo"
	"log"
	"testing"

	"google.golang.org/grpc"
)

func TestClient(t *testing.T) {
	conn, err := grpc.Dial("localhost:5040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	c, err := New(conn)
	if err != nil {
		panic(err)
	}

	resp, err := c.Echo(context.TODO(), &pb.EchoRequest{In: "hi hi hi"})
	if err != nil {
		panic(err)
	}
	log.Printf("%+v", resp)
}
