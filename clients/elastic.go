package clients

import (
	"crypto/tls"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
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

func NewLemmaBulkIndexer(es *elasticsearch.Client) esutil.BulkIndexer {
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         "lemma",          // The default index name
		Client:        es,               // The Elasticsearch client
		NumWorkers:    runtime.NumCPU(), // The number of worker goroutines
		FlushBytes:    5000000,          // The flush threshold in bytes
		FlushInterval: 30 * time.Second, // The periodic flush interval
	})
	if err != nil {
		log.Fatalf("Error creating LemmaBulk the indexer: %s", err)
	}

	return bi
}
