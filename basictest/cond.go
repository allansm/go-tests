package main

import "fmt"

func main(){
	var n int
	fmt.Print("n:")
	fmt.Scanln(&n)

	if n > 10{
		fmt.Println("n > 10")
	}else if n > 5{
		fmt.Println("n > 5")
	}else{
		fmt.Print("n < 6")
	}
}
