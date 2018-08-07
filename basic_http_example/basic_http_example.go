package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	port := 8888
	http.HandleFunc("/helloworld", helloworldHandler)
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloworldHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World! %s", time.Now())
}
