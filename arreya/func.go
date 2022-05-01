package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	b := make([]int, a, a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b[i])
	}
	fmt.Println(b[3])
}
