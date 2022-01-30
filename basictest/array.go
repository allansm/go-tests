package main

import "fmt"

func main(){
	x := [5] int {1,2,3,4,5}

	fmt.Println(x[0]+x[2])
	fmt.Println(x[3]-x[1])
	fmt.Println("")

	for _,n := range x{
		fmt.Println(n)
	}
}
