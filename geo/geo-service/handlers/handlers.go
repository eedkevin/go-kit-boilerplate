package handlers

import (
	"context"

	pb "geo"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.GeoServer {
	return geoService{}
}

type geoService struct{}

func (s geoService) Status(ctx context.Context, in *pb.StatusRequest) (*pb.StatusResponse, error) {
	var resp pb.StatusResponse
	return &resp, nil
}

func (s geoService) Distance(ctx context.Context, in *pb.DistanceRequest) (*pb.DistanceResponse, error) {
	var resp pb.DistanceResponse
	resp.Distance = 100
	return &resp, nil
}
