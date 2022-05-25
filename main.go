package main

import (
	"context"
	"fmt"
	"time"

	"github.com/andyklimenko/elastic-search-demo/config"
	"github.com/olivere/elastic"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(
		elastic.SetBasicAuth(cfg.Elastic.UserName, cfg.Elastic.Password),
		elastic.SetSniff(false),
	)
	if err != nil {
		// Handle error
		panic(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping(cfg.Elastic.Address).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}
