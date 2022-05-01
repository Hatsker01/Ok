package main

import "fmt"


func main()  {
	var a int
	fmt.Scan(&a)
	array := make([]int,a,a)
	for i:=0;i<a;i++{
		fmt.Scan(&array[i])
	}
	for i:=0;i<a;i+=2{
		fmt.Print(array[i]," ")

	}
	
}


