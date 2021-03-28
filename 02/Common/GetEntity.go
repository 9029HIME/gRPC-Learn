package Common

import (
	"fmt"
	"gRPC-Learn/02/ProtoEntity"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/anypb"
)

type ServerEntity struct {
}

//实现在protobuf里定义好的TestRequest接口
func (server *ServerEntity) TestRequest(context context.Context, request *ProtoEntity.Request) (*ProtoEntity.Response, error) {
	payload := request.GetPayload()
	requestPayload := new(ProtoEntity.RequestPayload)
	payload.UnmarshalTo(requestPayload)
	jsonContent := requestPayload.JsonContent
	fmt.Println("服务器收到的请求数据是:", jsonContent)

	responsePayload := &ProtoEntity.ResponsePayload{
		JsonContent: fmt.Sprintf("我是服务端，我已收到你的消息:%s", jsonContent),
	}
	any, err := anypb.New(responsePayload)
	response := &ProtoEntity.Response{
		Payload: any,
	}

	return response, err
}

func GetServer() *ServerEntity {
	return new(ServerEntity)
}

func GetRequest(pageNo int64, pageSize int64, payload *ProtoEntity.ResponsePayload) *ProtoEntity.Request {
	any, err := anypb.New(payload)
	if err != nil {
		panic(err)
	}
	return &ProtoEntity.Request{
		PageNo:   pageNo,
		PageSize: pageNo,
		Payload:  any,
	}
}
