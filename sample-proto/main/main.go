package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	subPb "my/grpc/test/proto"
)

type Coffee struct {
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

func RemoteFunc(index int64) *subPb.Coffee {
	conn, _ := grpc.Dial(":9001",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	subClient := subPb.NewSubClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := subClient.Func(ctx, &subPb.Input{Index: index})
	if err != nil {
		panic(err)
	}
	return res
}

func main() {
	fmt.Println("Main server is running")
	var index string
	for {
		fmt.Print("Input index: ")
		fmt.Scanln(&index)
		i, _ := strconv.ParseInt(index, 10, 64)
		fmt.Println(RemoteFunc(i))
	}
}
