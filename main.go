package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 2; i <= n; i++ {
		fmt.Println(i * i)
	}
}
