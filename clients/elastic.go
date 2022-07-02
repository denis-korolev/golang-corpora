package clients

import (
	"crypto/tls"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"net"
	"net/http"
	"os"
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
