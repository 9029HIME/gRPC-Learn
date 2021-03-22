package UnitTest

import (
	"fmt"
	"gRPC-Learn/01/ProtoEntity"
	"gRPC-Learn/01/pendingTest"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"testing"
)

/**
将一个protobuf实体转为二进制与转为json文件比较
本次测试结果：
protobuf：24KB
json：51KB
*/
func TestMarshal(t *testing.T) {
	file := "C:\\WorkPlace\\Golang\\FreePlace\\gRPC-Learn\\01\\tempFile\\marshal.bin"
	people := pendingTest.NewPeople("黄老师")
	data, err := proto.Marshal(people)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(file, data, 0644)
	if err != nil {
		panic(err)
	}
}

func TestUnMarshal(t *testing.T) {
	file := "C:\\WorkPlace\\Golang\\FreePlace\\gRPC-Learn\\01\\tempFile\\marshal.bin"
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	people := new(ProtoEntity.People)
	err = proto.Unmarshal(data, people)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Printf("反序列化后的peopl.desce是%s，type是%s", people.Desc, people.Type))
}

func TestJsonMarshal(t *testing.T) {
	file := "C:\\WorkPlace\\Golang\\FreePlace\\gRPC-Learn\\01\\tempFile\\people.json"
	people := pendingTest.NewPeople("黄老师")
	marshaler := jsonpb.Marshaler{}
	json, err := marshaler.MarshalToString(people)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(file, []byte(json), 0644)
	if err != nil {
		panic(err)
	}
}
