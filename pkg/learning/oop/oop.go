package oop

import "fmt"

type Attr struct {
	name   string
	age    int
	mobile string
}

type person interface {
	Name() string
	age() int
	Mobile() string
}

type Employee struct {
	Attr
}

func New() Employee {
	return Employee{
		Attr{
			name:   "name",
			age:    123,
			mobile: "123",
		},
	}
}

func (e Employee) Name() string {
	return e.name
}

func (e Employee) age() int {
	return e.Attr.age
}

func (e Employee) Mobile() string {
	return e.mobile
}

type Student struct{}

func (Student) Name() string {
	//TODO implement me
	panic("implement me")
}

func (Student) age() int {
	//TODO implement me
	panic("implement me")
}

func (Student) Mobile() string {
	//TODO implement me
	panic("implement me")
}

type LearnOop struct{}

func (LearnOop) Learn() {
	employee := New()
	fmt.Println("name => ", employee.Name())
	fmt.Println("age => ", employee.name)
}

func (LearnOop) Kind() string {
	return "OOP"
}
