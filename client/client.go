package main

import (
	"context"
	"fmt"
	"io"

	"github.com/emrahsariboz/grpc-weather/api"
	"google.golang.org/grpc"
)

const (
	addr = "localhost:8080"
)

func main() {

	conn, err := grpc.Dial(addr, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	client := api.NewWeatherServiceClient(conn)

	ListCities(client, &api.ListCitiesRequest{})

	QueryWeather(client, &api.WeatherRequest{CityCode: "tr_ank"})

	defer conn.Close()

}

func ListCities(l api.WeatherServiceClient, req *api.ListCitiesRequest) {

	fmt.Println("Listing Cities...")

	resp, err := l.ListCities(context.Background(), req)

	if err != nil {
		panic(err)
	}

	for _, v := range resp.Items {
		fmt.Printf("City code is %v and the name %v\n", v.GetCityCode(), v.GetCityName())
	}
}

func QueryWeather(l api.WeatherServiceClient, req *api.WeatherRequest) {
	fmt.Println("Printing the temperatures...")

	stream, err := l.QueryWeather(context.Background(), req)

	if err != nil {
		panic(err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Printf("\t temperature: %.2f\n", msg.GetTemperature())
	}

	fmt.Println("The server stopped sending...")
}
