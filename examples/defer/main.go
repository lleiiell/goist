package main

/*
deferAdd a 1 b 2 rt 3
deferAdd a 0 b 2 rt 2
4
3
2
1
0
deferAdd a 0 b 2 rt 2
deferAdd a 1 b 3 rt 4
*/
func main() {
	a := 1
	b := 2
	defer deferAdd(a, deferAdd(a, b))

	a = 0
	defer deferAdd(a, deferAdd(a, b))

	b = 1

	deferFor()
}

func deferAdd(a, b int) int {
	rt := a + b
	println("deferAdd", "a", a, "b", b, "rt", rt)
	return rt
}

func deferFor() {
	for i := 0; i < 5; i++ {
		defer println(i)
	}
}
