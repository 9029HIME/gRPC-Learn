// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: school.proto

package ProtoEntity

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//定义一个枚举来代表老师类型，注意值是顺序
type TeacherType int32

const (
	Teacher_MATH    TeacherType = 0
	Teacher_CHINESE TeacherType = 1
	Teacher_ENGLISH TeacherType = 2
)

// Enum value maps for TeacherType.
var (
	TeacherType_name = map[int32]string{
		0: "MATH",
		1: "CHINESE",
		2: "ENGLISH",
	}
	TeacherType_value = map[string]int32{
		"MATH":    0,
		"CHINESE": 1,
		"ENGLISH": 2,
	}
)

func (x TeacherType) Enum() *TeacherType {
	p := new(TeacherType)
	*p = x
	return p
}

func (x TeacherType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TeacherType) Descriptor() protoreflect.EnumDescriptor {
	return file_school_proto_enumTypes[0].Descriptor()
}

func (TeacherType) Type() protoreflect.EnumType {
	return &file_school_proto_enumTypes[0]
}

func (x TeacherType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TeacherType.Descriptor instead.
func (TeacherType) EnumDescriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{0, 0}
}

type Teacher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Teacher) Reset() {
	*x = Teacher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_school_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Teacher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Teacher) ProtoMessage() {}

func (x *Teacher) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Teacher.ProtoReflect.Descriptor instead.
func (*Teacher) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{0}
}

func (x *Teacher) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Student struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Teacher *Teacher `protobuf:"bytes,1,opt,name=teacher,proto3" json:"teacher,omitempty"`
	Name    string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Student) Reset() {
	*x = Student{}
	if protoimpl.UnsafeEnabled {
		mi := &file_school_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student.ProtoReflect.Descriptor instead.
func (*Student) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{1}
}

func (x *Student) GetTeacher() *Teacher {
	if x != nil {
		return x.Teacher
	}
	return nil
}

func (x *Student) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type People struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*People_Teacher
	//	*People_Student
	Type isPeople_Type `protobuf_oneof:"type"`
	Desc string        `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *People) Reset() {
	*x = People{}
	if protoimpl.UnsafeEnabled {
		mi := &file_school_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *People) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*People) ProtoMessage() {}

func (x *People) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use People.ProtoReflect.Descriptor instead.
func (*People) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{2}
}

func (m *People) GetType() isPeople_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *People) GetTeacher() *Teacher {
	if x, ok := x.GetType().(*People_Teacher); ok {
		return x.Teacher
	}
	return nil
}

func (x *People) GetStudent() *Student {
	if x, ok := x.GetType().(*People_Student); ok {
		return x.Student
	}
	return nil
}

func (x *People) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type isPeople_Type interface {
	isPeople_Type()
}

type People_Teacher struct {
	Teacher *Teacher `protobuf:"bytes,1,opt,name=teacher,proto3,oneof"`
}

type People_Student struct {
	Student *Student `protobuf:"bytes,2,opt,name=student,proto3,oneof"`
}

func (*People_Teacher) isPeople_Type() {}

func (*People_Student) isPeople_Type() {}

var File_school_proto protoreflect.FileDescriptor

var file_school_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49,
	0x0a, 0x07, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x54, 0x48, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x53, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x45, 0x4e, 0x47, 0x4c, 0x49, 0x53, 0x48, 0x10, 0x02, 0x22, 0x41, 0x0a, 0x07, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52,
	0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x70, 0x0a, 0x06,
	0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x48, 0x00, 0x52, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x07,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0e,
	0x5a, 0x0c, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_school_proto_rawDescOnce sync.Once
	file_school_proto_rawDescData = file_school_proto_rawDesc
)

func file_school_proto_rawDescGZIP() []byte {
	file_school_proto_rawDescOnce.Do(func() {
		file_school_proto_rawDescData = protoimpl.X.CompressGZIP(file_school_proto_rawDescData)
	})
	return file_school_proto_rawDescData
}

var file_school_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_school_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_school_proto_goTypes = []interface{}{
	(TeacherType)(0), // 0: Teacher.type
	(*Teacher)(nil),  // 1: Teacher
	(*Student)(nil),  // 2: Student
	(*People)(nil),   // 3: People
}
var file_school_proto_depIdxs = []int32{
	1, // 0: Student.teacher:type_name -> Teacher
	1, // 1: People.teacher:type_name -> Teacher
	2, // 2: People.student:type_name -> Student
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_school_proto_init() }
func file_school_proto_init() {
	if File_school_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_school_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Teacher); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_school_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Student); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_school_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*People); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_school_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*People_Teacher)(nil),
		(*People_Student)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_school_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_school_proto_goTypes,
		DependencyIndexes: file_school_proto_depIdxs,
		EnumInfos:         file_school_proto_enumTypes,
		MessageInfos:      file_school_proto_msgTypes,
	}.Build()
	File_school_proto = out.File
	file_school_proto_rawDesc = nil
	file_school_proto_goTypes = nil
	file_school_proto_depIdxs = nil
}
