package main

import (
	"fmt"
	"sync"
	"time"
)

func main (){
	fmt.Println("start")
	m := Manager{}
	err := m.initialize()
	fmt.Printf("final: %v\n", err)
}

type Manager struct {
	Version string
	Timeout time.Duration
	once sync.Once
}

func (m *Manager) initialize() (err error) {
	errChan := make(chan error, 1)
	defer close(errChan)
	//var err error
	m.once.Do(func() {
		cnf, err := phase1()
		fmt.Printf("phase1: %v", err)
		if err != nil {
			err = fmt.Errorf("failed to get Kubernetes config: %v\n", err)
			fmt.Printf("phase1: %v", err)
			errChan <- err
			return
		}

		err = phase2(cnf)
		if err != nil {
			err = fmt.Errorf("failed to create manager client: %v", err)
			errChan <- err
			return
		}
		if m.Version == "" {
			m.Version = "DefaultVersion"
		}
		if m.Timeout <= 0 {
			m.Timeout = 0
		}
	})
	err = <- errChan
	fmt.Printf("returned: %v\n", err)
	return err
}

func phase1() (string, error) {
	//return "", fmt.Errorf("phase1: %d", 1)
	return "", nil
}

func phase2(a string) error {
	if a == "" {
		return fmt.Errorf("phase2: %d", 2)
	}
	return fmt.Errorf("phase2: %d", 2)
}
