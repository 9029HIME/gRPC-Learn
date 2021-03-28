package Server

import (
	"context"
	"gRPC-Learn/03/ProtoEntity"
)

type ServerImpl struct {
}

func (server *ServerImpl) GetUserInfo(context context.Context, request *ProtoEntity.Request) (*ProtoEntity.Response, error) {
	return &ProtoEntity.Response{}, nil
}
func (server *ServerImpl) GetUserInfoStream(streamServer ProtoEntity.UserService_GetUserInfoStreamServer) error {
	return nil
}
