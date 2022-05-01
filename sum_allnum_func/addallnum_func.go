package main

import "fmt"


func main()  {
	var a int
	fmt.Scan(&a)
	fmt.Print(a/100+(a%100)/10+a%10)
	
}

