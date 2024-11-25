package main

import "fmt"

func IsOdd(n int) bool {
	return n%2 != 0
}

func main() {
	fmt.Println(IsOdd(14))
	fmt.Println(IsOdd(7))
}
