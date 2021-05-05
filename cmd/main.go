package main

import (
	"fmt"
	"net/http"

	"github.com/henzai/ranwei_functions/handler"
)

func main() {
	http.HandleFunc("/save", handler.SaveItem)
	fmt.Println("server start")
	http.ListenAndServe(":8080", nil)
}
