# 为什么要用Protobuf？

​	之前在Netty的学习中已经学习过Protobuf的使用，他的优点总的来说就是数据体积压缩率高，最终会被序列化为二进制数据。同一份数据转成JSON的大小是转成Protobuf的5倍，因此在微服务数据序列化中使用Protobuf'能减少网络传输量。

​	回忆一下Protobuf三要素：protobuf文件本身（需要用protobuf语言定义）、protobuf编译器（将protobuf文件编译成对应语言的源码文件，**这个文件也称为存根Stub**）、protobuf-API（对编译后的结构体进行序列化、反序列化等操作）。之前已熟悉过Java那一套东西，接下来需要学Go的。

​	上面说到Protobuf的传输体积小，其实还有一个优点：同一份Protobuf文件能编译多种语言的存根（只要是Google提供支持的语言），这样保证了复用性。



# Protobuf的使用

