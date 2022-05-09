package main

import (
	"fmt"
	"strconv"
)
	




func main() {
	var a int
	var b int 
	var c string
	fmt.Scan(&a)
	for a!=0{
		b=a%10

		b=b*b
		a/=10
		c=strconv.Itoa(b)+c
	}
	fmt.Println(c)
	
	


}
