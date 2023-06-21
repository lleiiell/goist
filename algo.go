package goist

/*
ClimbStairs3
Function: ClimbStairs3
Language: Go
Description: Given a number of stairs, the task is to find the total number of ways to climb the stairs where at each step, the climber can either climb 1, 2, or 3 steps.
Parameters:
x: Total number of stairs
Output: Number of ways to climb the stairs
*/
func ClimbStairs3(x int) int {
	if x <= 2 {
		return x
	} else if x == 3 {
		return 4
	}
	f1, f2, f3, f4 := 1, 2, 4, 0
	for i := 4; i <= x; i++ {
		f4 = f1 + f2 + f3
		f1 = f2
		f2 = f3
		f3 = f4
	}
	return f4
}
