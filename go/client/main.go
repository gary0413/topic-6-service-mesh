// client.go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")

	if serverHost == "" || serverPort == "" {
		panic("SERVER_HOST and SERVER_PORT environment variables are required")
		return
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		url := fmt.Sprintf("http://%s:%s/", serverHost, serverPort)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading response: %v", err)
			return
		}

		responseMessage := fmt.Sprintf("Response from server: %s", string(body))
		fmt.Fprintln(w, responseMessage)
	})

	fmt.Println("Client server is running on port 3000...")
	http.ListenAndServe(":3000", nil)
}
