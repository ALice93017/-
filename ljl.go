package main

import (
	"fmt"
	"math"
)

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Print("请输入一个整数：")
	fmt.Scan(&n)

	if IsPrime(n) {
		fmt.Printf("%d是素数\n", n)
	} else {
		fmt.Printf("%d不是素数\n", n)
	}
}
