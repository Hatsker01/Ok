package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scan(&i)
	if i%10 != i/100 && i/100 != (i%100)/10 && i%10 != (i%100)/10 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
