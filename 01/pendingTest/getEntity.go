package pendingTest

import "gRPC-Learn/01/ProtoEntity"

func main() {

}

func NewTeacher(name string) *ProtoEntity.Teacher {
	return &ProtoEntity.Teacher{
		Name:     name,
		TypeName: ProtoEntity.Teacher_MATH,
	}
}

func NewStudent(teacher *ProtoEntity.Teacher, name string) *ProtoEntity.Student {
	return &ProtoEntity.Student{
		Teacher: teacher,
		Name:    name,
	}
}

func NewPeople(name string) *ProtoEntity.People {
	return &ProtoEntity.People{
		Type: &ProtoEntity.People_Teacher{
			NewTeacher(name),
		},
		Desc: name,
	}
}
