package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"

	subPb "my/grpc/test/proto"
)

type Coffee struct {
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

type subServer struct {
	subPb.SubServer
}

func (s *subServer) Func(ctx context.Context, input *subPb.Input) (*subPb.Coffee, error) {
	fmt.Println("RemoteFunc is called, index: ", input.Index)
	data, err := os.Open("../data.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := io.ReadAll(data)
	var coffees []Coffee
	json.Unmarshal(byteValue, &coffees)
	var coffee Coffee
	if int(input.Index) >= 0 && int(input.Index) < len(coffees) {
		coffee = coffees[input.Index]
	} else {
		coffee = Coffee{}
	}
	return &subPb.Coffee{
		Name:  coffee.Name,
		Price: coffee.Price,
	}, err
}

func main() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	subPb.RegisterSubServer(server, &subServer{})
	fmt.Println("Sub server is running")
	if err := server.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %s", err)
	}
}
