package main

import (
	"context"
	"fmt"
	"time"

	"github.com/olivere/elastic"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Obtain a client and connect to the default Elasticsearch installation
	// on 127.0.0.1:9200. Of course, you can configure your client to connect
	// to other hosts and configure it in various other ways.
	client, err := elastic.NewClient()
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
	fmt.Printf("elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}
