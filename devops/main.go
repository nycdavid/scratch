package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(r http.ResponseWriter, req *http.Request) {
		fmt.Println(fmt.Sprintf("%s %s", req.Method, req.URL.Path))
		r.Write([]byte("Hello world!"))
	})

	fmt.Println("Listening on 0.0.0.0:3000...")
	http.ListenAndServe(":3000", nil)
}
