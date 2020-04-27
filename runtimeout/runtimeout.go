package main

import (
	"errors"
	"fmt"
	"time"
)
const DiscoveryRetryInterval = 5 * time.Second
type Config struct {
	Kind string `json:"kind,omitempty"`
}

// fetchKubeConfigWithTimeout tries to run fetchKubeConfigFunc on every DiscoveryRetryInterval, but until discoveryTimeout is reached
func fetchKubeConfigWithTimeout(apiEndpoint string) (*Config, error) {
	discoveryTimeout := 50 * time.Second
	timeout := time.After(50 * time.Second)
	for {
		fmt.Printf("[discovery] Trying to connect to API Server %q\n", apiEndpoint)
		fmt.Printf("[discovery] Successfully established connection with API Server %q\n", apiEndpoint)

		select {
		case <- time.After(DiscoveryRetryInterval):
			continue
		case <- timeout:
			err := errors.New(fmt.Sprintf("abort connecting to API servers after timeout of %v", discoveryTimeout))
			fmt.Printf("[discovery] %v\n", err)
			return nil, err
		}
	}
}

func main() {
	conf, err := fetchKubeConfigWithTimeout("hello world")
	fmt.Printf("%v, %v", conf, err)
}
