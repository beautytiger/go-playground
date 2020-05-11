package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

var PORT = 8080

// apt install nmap
// ncat localhost 8080
func main() {
	server, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		panic(err)
	}
	for {
		client, err := server.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s got con from %s\n", client.LocalAddr().String(), client.RemoteAddr().String())
		b := bufio.NewReader(client)
		for {
			line, err := b.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					fmt.Printf("%s client close conn to the server\n", client.RemoteAddr().String())
					break
				}
				panic(err)
			}
			l, err := client.Write([]byte(fmt.Sprintf("got data from client: %s", line)))
			if err != nil {
				panic(err)
			}
			fmt.Printf("write %d bytes to clinet\n", l)
		}
	}
}
