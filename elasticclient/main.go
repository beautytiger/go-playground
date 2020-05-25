package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"net/http"
	"os"
	"sync"
)

func main() {
	// Starting with elastic.v5, you must pass a context to execute each service
	ctx := context.Background()
	println()

	// Obtain a client and connect to the default Elasticsearch installation
	// on 127.0.0.1:9200. Of course you can configure your client to connect
	// to other hosts and configure it in various other ways.
	tp := http.DefaultTransport
	tp.(*http.Transport).MaxIdleConnsPerHost = 100
	httpClient := &http.Client{Transport: tp}
	client, err := elastic.NewClient(
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		elastic.SetHttpClient(httpClient),
	)
	if err != nil {
		// Handle error
		panic(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	// Search with a term query
	termQuery := elastic.NewTermQuery("machine.os", "linux")

	wg := sync.WaitGroup{}
	for i:=0 ; i<5 ; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				searchResult, err := client.Search().
					Index("kibana_sample_data_logs").   // search in index "twitter"
					Query(termQuery).   // specify the query
					//Sort("user", true). // sort by "user" field, ascending
					//From(0).Size(10).   // take documents 0-9
					//Pretty(true).       // pretty print request and response JSON
					Do(ctx)             // execute
				if err != nil {
					// Handle error
					panic(err)
				}
				// searchResult is of type SearchResult and returns hits, suggestions,
				// and all kinds of other information from Elasticsearch.
				fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
				fmt.Println(searchResult.Hits.TotalHits.Value)
			}
		}()
	}
	wg.Wait()
	fmt.Println("all work done")
}