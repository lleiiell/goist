package goist

import "fmt"

func ExampleStruct2Map() {

    st := struct {
        Name string `json:"name,omitempty"`
        True bool   `json:"true"`
    }{
        "till",
        true,
    }

    // m, err := Struct2Map(st)
    // fmt.Println(m["name"], m["true"], err)
    // Output: till true <nil>

    fmt.Println(Struct2Map(st))
    // Output: map[name:till true:true] <nil>
}
