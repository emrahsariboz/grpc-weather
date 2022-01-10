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
		fmt.Println("Something is wrong!")
	}

	client := api.NewWeatherServiceClient(conn)

	ctx := context.Background()

	resp, err := client.ListCities(ctx, &api.ListCitiesRequest{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Cities: ")

	for _, city := range resp.Items {
		fmt.Printf("\t%s: %s \n", city.GetCityCode(), city.GetCityName())
	}

	stream, err := client.QueryWeather(ctx, &api.WeatherRequest{CityCode: "tr_ank"})

	if err != nil {
		fmt.Println("Something is wrong!!!")
	}

	fmt.Println("Weather in Ankara")

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Printf("\t temperature: %.2f\n", msg.GetTemperature())
	}

	fmt.Println("Server stopped sending")

}
