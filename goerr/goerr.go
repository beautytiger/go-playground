package main

import (
	"fmt"
	"net/http"
)

func main() {
	testRepo := http.Server{}

	errChan := make(chan error)
	go func() {
		fmt.Println("starting server")
		errChan <- testRepo.ListenAndServe()
		fmt.Println("closing server")
	}()

	fmt.Println("stopping server")
	if err := testRepo.Close(); err != nil {
		fmt.Printf("Failed to close test repo server: %s", err)
	}
	fmt.Println("rev error")
	if err := <- errChan; err != nil && err != http.ErrServerClosed {
		fmt.Printf("Failed to run test repo server: %s", err)
	}
}
