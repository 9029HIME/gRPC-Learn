# 方法的定义与泛型

​	书接上一回protobuf01.md，这次要通过protobuf来定义方法。protobuf的方法是基于service的，一个service能有多个方法。最终落实到Go源码里是interface，这个interface名为XXXClient（XXX为protobuf里service的名字），最终在代码上要通过实现接口的形式来定义具体的业务方法（这里和SpringCloud、Dubbo有相似之处）。接下来是一个protobuf方法的定义示例。

```protobuf
syntax = "proto3";
option go_package = "/ProtoEntity";
import "google/protobuf/any.proto";

// 定义一个通用的请求类
message Request{
  int64 page_no = 1;
  int64 page_size = 2;
  // 用google.protobuf.Any来代替泛型
  google.protobuf.Any payload = 3;
}
// 定义一个通用的响应类
message Response{
  string status = 1;
  int32 code = 2;
  // 用google.protobuf.Any来代替泛型
  google.protobuf.Any payload = 3;
}

// 测试用Service
service CommonService{
	// rpc为关键字，代表这是一个protobuf方法
  rpc testRequest(Request) returns (Response){};
}
```

​	然而不清楚为什么，我直接编译这个protobuf会给我抛google/protobuf/any.proto不存在的异常，具体解决方法是：将google/protobuf这个包扔到${GOPATH}/src里，然后通过protoc --go_out=plugins=grpc:. -I=${GOPATH}/src -I=. *.proto来编译。

​	这个ANY代表的是任意类型，可以**暂时**认为是泛型，但和泛型不同的是需要调用anypb.New()返回一个any对象来赋值。如果你希望any是A结构体对象，就要anypb.New(A)，如果希望any是B结构体的对象，就要anypb.New(B)。**但是！！！A和B必须是proto.Message的派生，既需要被protobuf编译形成的结构体，因此兼容性有点低**

​	每一个Protobuf编译形成的结构体都会实现function ProtoReflect() protoreflect.Message方法，这是proto.Message的标志。**那么如果我有一个不是protobuf编译形成的任意结构体实现了proto.Message，将他转成Any可行吗？这个有待考究。**

```go
 foo := &pb.Foo{...}
 any, err := anypb.New(foo)
 if err != nil {
   ...
 }
 ...
 foo := &pb.Foo{}
 if err := any.UnmarshalTo(foo); err != nil {
   ...
 }
```
