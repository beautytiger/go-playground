package main

//https://github.com/golang/go/issues/24652
import (
	"crypto/x509"
	"fmt"
	"log"
)

func main() {
	certs, err := x509.SystemCertPool()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Num System Certs: %d\n", len(certs.Subjects()))
}