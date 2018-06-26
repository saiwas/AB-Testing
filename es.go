package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/olivere/elastic"
)

// Tweet is a structure used for serializing/deserializing data in Elasticsearch.
type Event struct {
	ID         int
	Name       string
	Attributes map[string]string
}

const mapping = `
{
	"settings":{
		"number_of_shards": 2,
		"number_of_replicas": 0
	},
	"mappings":{
		"golang_event":{
			"properties":{
				"name":{
					"type":"keyword"
				},
				"attributes":{
					"dynamic": true,
					"properties": {}
				}
			}
		}
	}
}`

func main() {
	ctx := context.Background()
	esUrl := "https://search-ab-testing-tools-sandbox-ubpfotq3qfjes6sfsqe3lqr334.us-east-1.es.amazonaws.com/"

	client, err := elastic.NewClient(
		elastic.SetURL(esUrl),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetMaxRetries(5),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	if err != nil {
		// Handle error
		fmt.Print("Error!!!!!!!!")
		panic(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping(esUrl).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := client.ElasticsearchVersion(esUrl)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists("golang_event").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex("golang_event").BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

	count := 10000000
	for i := 0; i < count; i++ {
		attributes := make(map[string]string)
		attributes["devidce"] = "mobile"
		attributes["send_by"] = "golang"
		attributes["test_id"] = string(i)
		golang_event := Event{ID: i, Name: "golang_event name", Attributes: attributes}
		put1, err := client.Index().
			Index("golang_event").
			Type("golang_event").
			Id(string(i)).
			BodyJson(golang_event).
			Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		fmt.Printf("Indexed golang_event %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	}

	// Flush to make sure the documents got written.
	_, err = client.Flush().Index("golang_event").Do(ctx)
	if err != nil {
		panic(err)
	}
}
