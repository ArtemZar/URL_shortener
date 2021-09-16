package main

import (
	"net/http"

	h "github.com/ArtemZar/URL_shortener/api/handler"
)

func main() {
	h.HandlesFunc()

	http.ListenAndServe("localhost:8080", nil)
}
