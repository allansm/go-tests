package main

import(
	"fmt"
	"strings"
)

func main(){
	if strings.Contains("helloworld", "hello"){
		fmt.Println(":D")
	}else{
		fmt.Println(":(")
	}
}

