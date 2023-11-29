package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"

	opensearchapi "github.com/opensearch-project/opensearch-go/opensearchapi"

	opensearch "github.com/opensearch-project/opensearch-go"
)

func main() {
	client, err := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Addresses: []string{"https://localhost:9200"},
		Username:  "admin", // For testing only. Don't store credentials in code.
		Password:  "admin",
	})
	if err != nil {
		fmt.Println("cannot initialize", err)
		os.Exit(1)
	}
	settings := strings.NewReader(`{
		'settings': {
			'index': {
					 'number_of_shards': 1,
					 'number_of_replicas': 2
					 }
				 }
		}`)

	// Create an index with non-default settings.
	res := opensearchapi.IndicesCreateRequest{
		Index: "books",
		Body:  settings,
	}
	fmt.Println(res)
	index := opensearchapi.IndicesGetRequest{Index: []string{"*"}, Human: true}

	indexResponse, err := index.Do(context.Background(), client)
	fmt.Println("Index for a document")
	fmt.Println(indexResponse)
}
