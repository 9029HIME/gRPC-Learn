# 服务端与客户端的代码

​	首先来看一下请求方法的定义

```protobuf
service CommonService{
  rpc testRequest(Request) returns (Response){};
}
```

​	定义了一个CommonService接口，里面一个方法testRequest，传参的Request，响应是Response。这个protobuf经过protoc的编译后会生成两个接口与两套实现：客户端接口&客户端实现、服务端接口&服务端实现

## 客户端代码

​	客户端接口：

```go
type CommonServiceClient interface {
   TestRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}
```

​	客户端实现：

```go
func (c *commonServiceClient) TestRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ProtoEntity.CommonService/testRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
```

​	**一般情况下客户端远程调用，只需要调用默认实现的方法即可**

## 服务端代码

​	服务端接口：

```go
type CommonServiceServer interface {
   TestRequest(context.Context, *Request) (*Response, error)
}
```

​	服务端实现：

```go
type UnimplementedCommonServiceServer struct {
}

func (*UnimplementedCommonServiceServer) TestRequest(context.Context, *Request) (*Response, error) {
   return nil, status.Errorf(codes.Unimplemented, "method TestRequest not implemented")
}
```

​	然而服务端默认实现是不合法的，这一部分就需要我们手动去实现，这样当客户端调用TestRequest时，本质上是调用了我们实现的服务端方法。

​	手动实现：

```go
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
```

​	本次demo的手动实现只是简单的拿到请求后打印，并告知客户端服务端已成功接收信息（值得注意的是Payload其实可以直接改成字符串json，而不用再包装一个requestPayload&responsePayload）。



# 服务端与客户端的配置

## 服务端配置

```go
func main() {
   // 先启动tcp监听
   listen, err := net.Listen("tcp", ":8003")
   if err != nil{
      panic(err)
   }
   // 新建一个gRPC服务器
   grpcServer := grpc.NewServer()
   
   // 将上面自定义的服务端注册到gRPC服务器内 
   ProtoEntity.RegisterCommonServiceServer(grpcServer,&Common.ServerEntity{})
   // TODO 如果有多个服务端（多个接口），还是可以用同样的方式将自定义服务端注册到gRPC服务器内
   
   // 注册gRPC服务器
   reflection.Register(grpcServer)
   // 将启动的tcp监听注册到gRPC服务器内，通过gRPC来代替监听
   err = grpcServer.Serve(listen)
   if err != nil{
      panic(err)
   }
}
```

​	总的来说是：

​	1.一个gRPC服务器对应多个自定义服务端（即多个protobuf接口实现）

​	2.一个自定义服务端对应多个接口方法（即一个protobuf接口实现对应多个rpc接口方法实现）

​	3.配置好gRPC服务器后，注册gRPC服务器

​	4.用配置好的gRPC服务器代理已开启好的tcp监听，等待请求的到来

## 客户端配置

```go
func main() {
    // 启动一个gRPC连接
   conn, err := grpc.Dial("localhost:8003", grpc.WithInsecure())
   defer conn.Close()
   if err != nil{
      panic(err)
   }
    // 通过连接 + protobuf包来生成一个客户端
   client := ProtoEntity.NewCommonServiceClient(conn)
    // TODO 如果有多个Protobuf包（不同包下有不同的接口），可以用同样的方式生成多个客户端
   	
    // 组装请求
   jsonContent := "Hello,I am gRPC client"
   requestPayload := &ProtoEntity.RequestPayload{
      JsonContent: jsonContent,
   }
   any,_ := anypb.New(requestPayload)
   request := &ProtoEntity.Request{
      Payload:any,
   }
   
    // 通过特定客户端来调用特定方法，获取响应
   response, err := client.TestRequest(context.Background(), request)
   if err != nil{
      panic(err)
   }
   responsePayload := &ProtoEntity.ResponsePayload{}
   response.GetPayload().UnmarshalTo(responsePayload)
   fmt.Println("来自服务端的响应是:",responsePayload.JsonContent)
}
```

​	总的来说是：

​	1.根据不同的protobuf包和不同连接，生成对应的客户端

​	2.直接通过客户端调对应的rpc接口方法实现，实现远程调用

​	3.一般来说客户端用默认的方法即可，当然也可以手动重写