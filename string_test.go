package goist

import (
    "fmt"
    "golang.org/x/example/stringutil"
)

func ExampleMD5() {
    fmt.Println(MD5("hello"))
    // Output: 5d41402abc4b2a76b9719d911017c592
}

func ExampleReverse() {
    fmt.Println(stringutil.Reverse("hello"))
    // Output: olleh
}
