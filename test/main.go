package main

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
)

func main() {
	//log.Error("11", "12312321")

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://127.0.0.1:9200/",
		},
		APIKey: "",
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	// API Key should have cluster monitoring rights
	infores, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	fmt.Println(infores)
}
