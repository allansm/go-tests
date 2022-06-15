package main

import(
	"encoding/json"
	"fmt"
)

type Test struct{
	Id  int `json:"id"`
	Title string `json:"title"`
}

func main(){
	xyz := map[string]int{"x":2,"y":4,"z":5}
	tmp, _ := json.Marshal(xyz)
	fmt.Println(string(tmp))

	test := &Test{
		Id:1,
		Title:"i am a test"}

	tmp, _ = json.Marshal(test)
	fmt.Println(string(tmp))

	var msg map[string]interface{}

	json.Unmarshal([]byte(`{"msg1":"hello", "msg2":"world"}`), &msg)

	fmt.Print(msg["msg1"])
	fmt.Print(msg["msg2"])
}
