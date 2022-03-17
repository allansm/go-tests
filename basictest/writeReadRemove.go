package main

import (
	"os"
	"io/ioutil"
	"fmt"
)

func main() {
	f, _ := os.Create("hello.txt")
	_, _ = f.WriteString("helloworld\n")
	f.Close()

	text, _ := ioutil.ReadFile("hello.txt")
	fmt.Println(string(text))

	_ = os.Remove("hello.txt")
}
