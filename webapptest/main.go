package main

import (
    "fmt"
    "log"
    "net/http"
)

func form(w http.ResponseWriter, r *http.Request){
	html := "<html><form method='POST' action='/test'><input type='text' name='msg'><input type='submit' value='test'></form></html>"
	fmt.Fprintf(w, "%s", html)
}

func action(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		msg := r.FormValue("msg")
		fmt.Fprintf(w, "%s", msg)
	}
}

func main() {
	fmt.Println("running on :8080")
	
	http.HandleFunc("/", form)
	http.HandleFunc("/test", action)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
