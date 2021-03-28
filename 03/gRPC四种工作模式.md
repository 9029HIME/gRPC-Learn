# gRPC四种工作模式

## 简单RPC模式

​	客户端传入一个请求对象，服务端响应一个响应对象

​	定义方法：

```protobuf
service CommonService{
	// 和之前的gRPC简单demo一样
  rpc testRequest(Request) returns (Response){};
}
```

## 服务端流式 RPC

​	客户端传入一个请求对象，服务端响应多个响应对象

```protobuf
service CommonService{
	// 和之前不同的是，结果类型前面多了个stream
  rpc ServerStream(Request) returns (stream Response){};
}
```

## 客户端流式 RPC

​	客户端传入多个请求对象，服务端响应一个响应对象

```protobuf
service CommonService{
	// 这次是请求类型多了个stream
  rpc ClientStream(stream Request) returns (Response){};
}
```

## 双向流式 RPC

​	结合客户端流式RPC和服务端流式RPC，可以传入多个请求对象，返回多个结果对象

```protobuf
service CommonService{
  rpc testRequest(Request) returns (Response){};
}
```



# 为什么会有四种工作模式？他们的应用场景是？

​	

​	