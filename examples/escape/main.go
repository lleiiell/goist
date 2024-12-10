package main

/*
go run -gcflags '-m -l' main.go
# command-line-arguments
./main.go:29:16: leaking param: name
./main.go:29:38: leaking param: nick
./main.go:32:9: &Person{...} escapes to heap
Alice 10

*/
func main() {

	printlnPerson()

}

func printlnPerson() {
	person := NewPerson("Alice", 10, "Ali")
	println(person.name, person.age)
}

type Person struct {
	name string
	nick string
	age  int
}

func NewPerson(name string, age int, nick string) *Person {
	// p := Person{name: name} // 局部变量 p 在函数返回后逃逸
	// return &p               // 返回指向局部变量的指针
	return &Person{name: name, age: age, nick: nick}
}
