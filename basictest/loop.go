package main

import "fmt"

func main(){
	i := 0

	for i = 0; i < 100; i++ {
		fmt.Println(i)
	}

	i = 100

	for i > 0{
		i -= 10
		fmt.Println(i)
	}

	i = 0

	for{
		fmt.Println(i)

		if i < 10{
			i+=1
		}else{
			break;
		}
	}
}
