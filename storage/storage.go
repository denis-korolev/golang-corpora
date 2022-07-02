package storage

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"net"
	"net/http"
	"os"
	"parser/entities"
	"time"
)

type LemmaStorageInterface interface {
	InsertNewLemma(l entities.Lemma) error
}

type LoggerLemmaStorage struct {
	storage LemmaStorageInterface
	logger  *log.Logger
}

func NewLoggerStorage(storage LemmaStorageInterface, logger *log.Logger) (*LoggerLemmaStorage, error) {
	if logger == nil {
		logger = log.Default()
	}

	return &LoggerLemmaStorage{
		storage: storage,
		logger:  logger,
	}, nil
}

func (s *LoggerLemmaStorage) InsertNewLemma(l entities.Lemma) error {
	s.logger.Printf("inset new lemma: %s\n", l.ShortString())
	if s.storage != nil {
		return s.storage.InsertNewLemma(l)
	}
	return nil
}

type TimerLemmaStorage struct {
	storage LemmaStorageInterface
	logger  *log.Logger
}

func NewTimerStorage(storage LemmaStorageInterface, logger *log.Logger) (*TimerLemmaStorage, error) {
	if logger == nil {
		logger = log.Default()
	}

	return &TimerLemmaStorage{
		storage: storage,
		logger:  logger,
	}, nil
}

func (s *TimerLemmaStorage) InsertNewLemma(l entities.Lemma) error {
	t := time.Now()
	var err error
	if s.storage != nil {
		err = s.storage.InsertNewLemma(l)
	}
	s.logger.Printf("insert lemma.id=%s took %s", l.ID, time.Since(t).String())
	return err
}

type ElasticLemmaStorage struct {
	conn    *elasticsearch.Client
	timeout time.Duration
}

func NewElasticStorage() (*ElasticLemmaStorage, error) {
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

	conn, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &ElasticLemmaStorage{
		conn:    conn,
		timeout: 10,
	}, nil
}

func (s *ElasticLemmaStorage) InsertNewLemma(l entities.Lemma) error {
	body, _ := json.Marshal(&l)

	// Set up the request object.
	req := esapi.IndexRequest{
		Index:      "lemma",
		DocumentID: l.ID,
		Body:       bytes.NewReader(body),
		Refresh:    "true",
	}

	// Perform the request with the client.
	ctx, _ := context.WithTimeout(context.Background(), time.Second*s.timeout)
	res, err := req.Do(ctx, s.conn)
	if err != nil {
		return fmt.Errorf("error getting response: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[%s] error indexing document ID=%s", res.Status(), l.ID)
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			return fmt.Errorf("error parsing the response body: %s", err)
		}
	}

	return nil
}
