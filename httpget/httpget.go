package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var paths = []string{
	"http://www.baidu.com",
	"https://www.sogou.com",
	"https://www.google.com",
}

var client = http.Client{Timeout: time.Second*1}


func main() {
	fmt.Println("start request")
	ch := make(chan []byte)
	ctx, cancel := context.WithCancel(context.Background())
	for i := range paths {
		url := paths[i]
		go httpGet(ctx, url, ch)
	}
	loop:
	for {
		select {
		case <- time.After(time.Second * 2):
			cancel()
			break loop
		case data, ok := <- ch:
			fmt.Println(len(data))
			//fmt.Println(string(data))
		}
	}

	fmt.Println("end request")
}

func httpGet(ctx context.Context, url string, ch chan<- []byte) {
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("get", url, err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("get", url, resp.Status)
	if resp.StatusCode == 200 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		//log.Print(resp.Header)
		ch <- bodyBytes
	}
}
