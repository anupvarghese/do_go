package main

import "net/http"

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8008", nil)
}

func helloWorld(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello Wold"))
}
