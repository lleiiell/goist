package goist

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"testing"
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

/*
=== RUN   TestPipeCommands
Wed Nov 13 16:11:04 CST 2024
Wed Nov 13 16:11:09 CST 2024

--- PASS: TestPipeCommands (5.01s)
PASS
*/
func TestPipeCommands(t *testing.T) {
	rt, err := exec.Command("bash", "-c", "date && sleep 5 && date").Output()
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%s\n", rt)
}
