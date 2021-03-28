package main

import (
	"context"
	"fmt"
	"gRPC-Learn/02/ProtoEntity"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
)

func main() {
	conn, err := grpc.Dial("localhost:8003", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	client := ProtoEntity.NewCommonServiceClient(conn)
	jsonContent := "Hello,I am gRPC client"
	requestPayload := &ProtoEntity.RequestPayload{
		JsonContent: jsonContent,
	}
	any, _ := anypb.New(requestPayload)
	request := &ProtoEntity.Request{
		Payload: any,
	}
	response, err := client.TestRequest(context.Background(), request)
	if err != nil {
		panic(err)
	}
	responsePayload := &ProtoEntity.ResponsePayload{}
	response.GetPayload().UnmarshalTo(responsePayload)
	fmt.Println("来自服务端的响应是:", responsePayload.JsonContent)
}
