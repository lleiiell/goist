package goist_test

import (
	"fmt"
	"github.com/lleiiell/goist"
	"testing"
)

func ExampleSum() {
	fmt.Println(goist.Sum([]int{1, 2, 3, 4, 5}))
	fmt.Println(goist.Sum([]int64{1, 2, 3, 4, 5}))
	fmt.Println(goist.Sum([]float64{1.1, 2.2, 3.3, 4.4, 5.5}))

	// Output: 15
	// 15
	// 16.5
}

func TestIdMask(t *testing.T) {
	fmt.Println(goist.IdMask(1))
	fmt.Println(goist.IdMask(999))
	fmt.Println(goist.IdMask(goist.Rand().Int63()))
}

func ExampleIdUnmask() {
	fmt.Println(goist.IdUnmask("1clb"))
	fmt.Println(goist.IdUnmask("rr5fk"))
	fmt.Println(goist.IdUnmask("1djlwzxe5tj1a4qw"))
	// Output: 1
	// 999
	// 6521156348675224222
}

func ExampleFloatDivide() {
	a := uint(10)
	b := uint(3)

	fmt.Println(float64(a / b))
	fmt.Println(float32(a) / float32(b))
	fmt.Println(float64(a) / float64(b))
	// Output: 3
	// 3.3333333
	// 3.3333333333333335
}

func ExampleXor() {
	var a, b uint64
	a = 0xfeb6372a8750eb1d
	b = 0xfeb6372a8750eb1e
	fmt.Println(fmt.Sprintf("%b, %b", a, b))
	fmt.Println(a ^ b)

	fmt.Println(8 ^ 7)
	// Output: 1111111010110110001101110010101010000111010100001110101100011101, 1111111010110110001101110010101010000111010100001110101100011110
	// 3
	// 15
}

func ExampleRound2NearestHalf() {

	fs := []float64{0.1, 0.9, 1.1, 1.5, 1.9, 2.1, 2.5, 2.9, 3.1, 3.5, 3.9, 4.1, 4.5, 4.9, 5.1, 5}

	for _, v := range fs {
		fmt.Println(goist.Round2NearestHalf(v))
	}

	// Output: 0
	// 1
	// 1
	// 1.5
	// 2
	// 2
	// 2.5
	// 3
	// 3
	// 3.5
	// 4
	// 4
	// 4.5
	// 5
	// 5
	// 5
}
