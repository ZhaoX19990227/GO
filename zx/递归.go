package main

func main() {

	print(getSum(5))
}

func getSum(n int) int {
	if n == 1 {
		return 1
	}
	return getSum(n-1) + n
}
