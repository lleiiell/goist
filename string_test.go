package goist

import (
	"fmt"
	"testing"
)

func TestUuidSimple(t *testing.T) {
	uid := UuidSimple()
	if len(uid) < 22 {
		t.Error(len(uid), uid)
		return
	}
	fmt.Println(uid)
}

func TestRandomStr(t *testing.T) {
	str := RandStr(5)
	if len(str) != 5 {
		t.Error(str)
		return
	}
	str1 := RandStr(5)
	if len(str1) != 5 {
		t.Error(str1)
		return
	}

	fmt.Println(str)
	fmt.Println(str1)
}

func ExampleStrIsEmpty() {
	s1 := ""
	s2 := "a bc"
	s3 := "  "
	s4 := "\n \t"
	fmt.Println(IsStrEmpty(s1), IsStrEmpty(s2), IsStrEmpty(s3), IsStrEmpty(s4))
	// Output: true false true true
}

func ExampleMD5() {
	fmt.Println(MD5("hello"))
	// Output: 5d41402abc4b2a76b9719d911017c592
}

func TestToken(t *testing.T) {
	fmt.Println(Token())
	fmt.Println(Token())
	fmt.Println(Token())
	fmt.Println(Token())
	fmt.Println(Token())
}

func TestIsAscii(t *testing.T) {
	fmt.Println(IsAscii("你好"))
	fmt.Println(IsAscii("hello"))

	// Output: true
	// false

}

func TestStringDiff(t *testing.T) {
	a := "颊琰👁霓"
	b := "颊琰👁霓斡"

	same, diffA, diffB := StringDiff(a, b)
	if same != "颊琰👁霓" || diffA != "" || diffB != "斡" {
		t.Error(same, diffA, diffB)
		return
	}
	fmt.Println("same", same, "diffA", diffA, "diffB", diffB)
}

func TestWordFrequency(t *testing.T) {
	a := "test word frequency, world  word"
	b := "hello world word"

	reta := wordFrequency(a)
	retb := wordFrequency(b)
	fmt.Println("reta", reta, "retb", retb)
}
