package main

import (
	"fmt"
	"io"
	"net/http"
)

const port = 8123

func main() {
	// Set up an HTTP server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(debugRequest),
	}

	// Start the server
	fmt.Printf("Listening on port %d...\n", port)
	_ = server.ListenAndServe()
}

func debugRequest(w http.ResponseWriter, r *http.Request) {
	// Print all headers
	fmt.Println("Headers:")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Println(name, value)
		}
	}

	// Print all cookies
	fmt.Println("Cookies:")
	for _, cookie := range r.Cookies() {
		fmt.Println(cookie.Name, cookie.Value)
	}

	// Print all query parameters
	fmt.Println("Query parameters:")
	for name, values := range r.URL.Query() {
		for _, value := range values {
			fmt.Println(name, value)
		}
	}

	// Print the request body
	fmt.Println("Request body:")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	// Print the request method
	fmt.Println("Request method:")
	fmt.Println(r.Method)

	// Print the request URL
	fmt.Println("Request URL:")
	fmt.Println(r.URL)

	// Print the request protocol
	fmt.Println("Request protocol:")
	fmt.Println(r.Proto)

	// Print the request host
	fmt.Println("Request host:")
	fmt.Println(r.Host)

	// Print the request remote address
	fmt.Println("Request remote address:")
	fmt.Println(r.RemoteAddr)

	// Print the request path
	fmt.Println("Request path:")
	fmt.Println(r.URL.Path)

	// Print the request scheme
	fmt.Println("Request scheme:")
	fmt.Println(r.URL.Scheme)

	// Print the request user agent
	fmt.Println("Request user agent:")
	fmt.Println(r.UserAgent())

	// Print the request referer
	fmt.Println("Request referer:")
	fmt.Println(r.Referer())

	// Print the request content length
	fmt.Println("Request content length:")
	fmt.Println(r.ContentLength)

	// Send a response with status code 200
	fmt.Println("Sending response...")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}
