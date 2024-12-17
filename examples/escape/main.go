package main

import "math/rand"

/*
go run -gcflags '-m -l' main.go
*/
func main() {

	printlnPerson()

}

func printlnPerson() {
	person := pointEscapePerson("Alice", 10, "Ali")
	println(person.name, person.age)

	interfaceEscape(person.name)

	stackEscapeGenerate8191()

	stackEscapeGenerate8192()

	stackEscapeGenerate(1)

	incr := closureEscape()
	incr()
}

type Person struct {
	name string
	nick string
	age  int
}

/*
leaking param: name
leaking param: nick
&Person{...} escapes to heap
*/
func pointEscapePerson(name string, age int, nick string) *Person {
	return &Person{name: name, age: age, nick: nick}
}

/*
leaking param content: name
new(Person) does not escape
*/
func interfaceEscape(name interface{}) {
	var userInfo = new(Person)
	userInfo.name = name.(string)
	return
}

// make([]int, 8192) does not escape
func stackEscapeGenerate8191() {
	nums := make([]int, 8192) // = 64KB
	for i := 0; i < 8191; i++ {
		nums[i] = rand.Int()
	}
}

// make([]int, 8193) escapes to heap
func stackEscapeGenerate8192() {
	nums := make([]int, 8193) // > 64KB
	for i := 0; i < 8192; i++ {
		nums[i] = rand.Int()
	}
}

// make([]int, n) escapes to heap
func stackEscapeGenerate(n int) {
	nums := make([]int, n) // 不确定大小
	for i := 0; i < n; i++ {
		nums[i] = rand.Int()
	}
}

/*
moved to heap: n
func literal escapes to heap
*/
func closureEscape() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}
