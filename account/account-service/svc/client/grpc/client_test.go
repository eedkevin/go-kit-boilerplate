package grpc

import (
	"context"
	"log"
	"testing"

	"google.golang.org/grpc"

	pb "account"
	client "account/account-service/svc/client/grpc"
)

func TestClient(t *testing.T) {
	conn, err := grpc.Dial("localhost:5040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	c, err := client.New(conn)
	if err != nil {
		panic(err)
	}

	resp, err := c.Status(context.TODO(), &pb.StatusRequest{Full: true})
	if err != nil {
		panic(err)
	}
	log.Printf("%+v", resp)

	respAuthToken, err := c.AuthToken(context.TODO(), &pb.AuthTokenRequest{Username: "fhidahf", Password: "492338902"})
	if err != nil {
		panic(err)
	}
	log.Printf("%+v", respAuthToken)

}
