package clients

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"log"
	"net"
	"net/http"
	"os"
	"parser/entities"
	"runtime"
	"strings"
	"sync/atomic"
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

func BulkLemma(lemmaChan chan entities.Lemma, es *elasticsearch.Client) {

	var countSuccessful uint64

	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         "lemma",          // The default index name
		Client:        es,               // The Elasticsearch client
		NumWorkers:    runtime.NumCPU(), // The number of worker goroutines
		FlushBytes:    5000000,          // The flush threshold in bytes
		FlushInterval: 30 * time.Second, // The periodic flush interval
	})
	if err != nil {
		log.Fatalf("Error creating the indexer: %s", err)
	}
	start := time.Now().UTC()
	for a := range lemmaChan {
		// Prepare the data payload: encode article to JSON
		//
		data, err := json.Marshal(a)
		if err != nil {
			log.Fatalf("Cannot encode article %d: %s", a.ID, err)
		}

		// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
		//
		// Add an item to the BulkIndexer
		//
		err = bi.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				// Action field configures the operation to perform (index, create, delete, update)
				Action: "index",

				// DocumentID is the (optional) document ID
				DocumentID: a.ID,

				// Body is an `io.Reader` with the payload
				Body: bytes.NewReader(data),

				// OnSuccess is called for each successful operation
				OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
					atomic.AddUint64(&countSuccessful, 1)
				},

				// OnFailure is called for each failed operation
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						log.Printf("ERROR: %s", err)
					} else {
						log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
					}
				},
			},
		)
		if err != nil {
			log.Fatalf("Unexpected error: %s", err)
		}
	}

	if err := bi.Close(context.Background()); err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}

	biStats := bi.Stats()

	// Report the results: number of indexed docs, number of errors, duration, indexing rate
	//
	log.Println(strings.Repeat("▔", 65))

	dur := time.Since(start)

	if biStats.NumFailed > 0 {
		log.Fatalf(
			"Indexed [%s] documents with [%s] errors in %s (%s docs/sec)",
			int64(biStats.NumFlushed),
			int64(biStats.NumFailed),
			dur.Truncate(time.Millisecond),
			int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed)),
		)
	} else {
		log.Printf(
			"Sucessfuly indexed [%s] documents in %s (%s docs/sec)",
			int64(biStats.NumFlushed),
			dur.Truncate(time.Millisecond),
			int64(1000.0/float64(dur/time.Millisecond)*float64(biStats.NumFlushed)),
		)
	}

}
