package main

import (
	"fmt"
	"os"
	"net/http"
)

func main() {
	color := os.Getenv("COLOR")

	if color == "" {
		color = "unspecified"
	}

	hostname, _ := os.Hostname()

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		request.Body.Close()

		response.WriteHeader(http.StatusOK)
		fmt.Fprintln(response, color, hostname)
	})

	http.ListenAndServe(":8080", nil)
}