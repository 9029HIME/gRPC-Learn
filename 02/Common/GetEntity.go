package Common

import (
	"gRPC-Learn/02/ProtoEntity"
	"google.golang.org/protobuf/types/known/anypb"
)

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
