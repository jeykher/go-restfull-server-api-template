package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server deployed on port 5000")
	http.ListenAndServe(":5000", nil)
}
