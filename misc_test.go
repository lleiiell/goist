package goist

import (
	"encoding/json"
	"fmt"
)

func ExampleJsonPreUnmarshal() {
	type People struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	js := `{
		"name":"link"
	}`
	var p People
	p.Age = 22
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p)

	// Output: people:  {link 22}
}
