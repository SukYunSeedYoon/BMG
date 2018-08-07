package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type helloworldResponse struct {
	Message string `json:"message"`

	// 필드 출력 안함
	Author string `json:"-"`

	// 값이 비어 있을 때 필드 출력 금지
	Date string `json:",omitempty"`

	// 출력을 문자열로 변환하고 이름을 id로 바꾼다
	ID int `json:"id,string"`
}

func main() {
	port := 8888
	http.HandleFunc("/helloworld", helloworldHandler)
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloworldHandler(w http.ResponseWriter, r *http.Request) {

	response := helloworldResponse{Message: fmt.Sprintf("helloworld : %v", time.Now())}

	// 1. 버퍼에 한번 쓴 후 보내기
	//data, err := json.Marshal(response)
	// if err != nil {
	// 	panic("Oops")
	// }

	// fmt.Fprintf(w, string(data))

	// 2. 바로 스트림에 쓰기
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)

}
