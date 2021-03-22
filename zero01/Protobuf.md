# 为什么要用Protobuf？

​	之前在Netty的学习中已经学习过Protobuf的使用，他的优点总的来说就是数据体积压缩率高，最终会被序列化为二进制数据。同一份数据转成JSON的大小是转成Protobuf的5倍，因此在微服务数据序列化中使用Protobuf'能减少网络传输量。

​	回忆一下Protobuf三要素：protobuf文件本身（需要用protobuf语言定义）、protobuf编译器（将protobuf文件编译成对应语言的源码文件，**这个文件也称为存根Stub**）、protobuf-API（对编译后的结构体进行序列化、反序列化等操作）。之前已熟悉过Java那一套东西，接下来需要学Go的。

​	上面说到Protobuf的传输体积小，其实还有一个优点：同一份Protobuf文件能编译多种语言的存根（只要是Google提供支持的语言），这样保证了复用性。



# Protobuf的使用

## 事前准备

​	添加依赖

​	go get google.golang.org/grpc

​	go get github.com/golang/protobuf/protoc-gen-go

​	注意！！！此时$GOPATH/bin会出现protoc-gen-go.exe，还要再去[https://github.com/protocolbuffers/protobuf/releases](https://links.jianshu.com/go?to=https%3A%2F%2Fgithub.com%2Fprotocolbuffers%2Fprotobuf%2Freleases)里下一个win64的压缩包，将protoc.exe解压到$GOPATH/bin内（此时要保证$GOPATH/bin被添加到Path变量里）。

## Protobuf编写

```protobuf
syntax = "proto3";
//注意这个坑！！！！一定要指定生成的go文件的包名，且前面要带/
option go_package = "/pendingTest";

message CPU{
//  记得！这个1代表属性的序号，代表值
  string brand = 1;
  string name = 2;
//  属性的命名用下划线
  uint32 number_cores = 3;
}
```

## 文件的生成

​	在proto文件所在的文件夹输入 protoc --go_out=. *.proto，即可根据具体包名生成对应的go源文件

![生成的go文件](https://user-images.githubusercontent.com/48977889/111899110-65d94580-8a65-11eb-8bd5-3bb4445a0eed.jpg)



# Protobuf再一次详细介绍

## 嵌套使用

​	定义经典的学生-老师实体

```protobuf
syntax = "proto3";
option go_package = "/ProtoEntity";


message Teacher{
  //定义一个枚举来代表老师类型，注意值是顺序
  enum type{
    MATH = 0;
    CHINESE = 1;
    ENGLISH = 2;
  }
  string name = 1;
}

message Student{
  Teacher teacher = 1;
  string name = 2;
}

message People{
  //oneof即type的类型只能是Teacher或Student，type在Go源码中以接口形式存在
  oneof type{
    Teacher teacher = 1;
    Student student = 2;
  }
  string desc = 3;
}
```

​	