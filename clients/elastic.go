package clients

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func CreateElasticClient() (*elasticsearch.Client, error) {
	URL, ok := os.LookupEnv("ELASTIC_URL")
	if !ok || len(URL) == 0 {
		log.Fatalf("elastic_url: environment variable not declared: ELASTIC_URL")
	}

	PORT, ok := os.LookupEnv("ELASTIC_PORT")
	if !ok || len(PORT) == 0 {
		log.Fatalf("elastic_port: environment variable not declared: ELASTIC_PORT")
	}

	cfg := elasticsearch.Config{
		Addresses: []string{
			URL + ":" + PORT,
		},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   100,
			ResponseHeaderTimeout: time.Second * 10,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}

	return elasticsearch.NewClient(cfg)
}

func IndexLemmaData(documentID string, body []byte, es *elasticsearch.Client) {
	// Set up the request object.
	req := esapi.IndexRequest{
		Index:      "lemma",
		DocumentID: documentID,
		Body:       strings.NewReader(string(body)),
		Refresh:    "true",
	}

	//// Perform the request with the client.
	res, err := req.Do(context.Background(), es)
	//log.Println(res)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document ID=%d", res.Status())
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			//log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
}
