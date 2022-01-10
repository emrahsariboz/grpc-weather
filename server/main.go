package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/emrahsariboz/grpc-weather/api"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Something is not right", err)
	}

	srv := grpc.NewServer()

	//TODO: register service

	api.RegisterWeatherServiceServer(srv, &WeatherService{})

	fmt.Println("Starting the server!")

	srv.Serve(lis)

}

type WeatherService struct {
	api.UnimplementedWeatherServiceServer
}

func (w WeatherService) QueryWeather(req *api.WeatherRequest, resp api.WeatherService_QueryWeatherServer) error {

	for {
		err := resp.Send(&api.WeatherResponse{Temperature: rand.Float32()*10 + 10})
		if err != nil {
			break
		}
		time.Sleep(time.Second)
	}

	return nil
}

func (w WeatherService) ListCities(ctx context.Context, req *api.ListCitiesRequest) (*api.ListCitiesResponse, error) {
	return &api.ListCitiesResponse{
		Items: []*api.CityEntry{
			&api.CityEntry{
				CityCode: "tr_ank",
				CityName: "Ankara",
			},
			&api.CityEntry{
				CityCode: "tr_ist",
				CityName: "Istanbul",
			},
		},
	}, nil

}
