package storage

import "log"

const (
	WithElastic uint32 = 1 << iota
	WithLogger
	WithTimer
)

func CreateLemmaStorage(storageType uint32, logger *log.Logger) (s LemmaStorageInterface) {
	if storageType&WithElastic > 0 {
		ss, err := NewElasticStorage()
		if err != nil {
			log.Fatalf("Error creating the client: %s", err)
		}
		s = ss
	}
	if storageType&WithLogger > 0 {
		ss, err := NewLoggerStorage(s, logger)
		if err != nil {
			log.Fatalf("Error creating the client: %s", err)
		}
		s = ss
	}
	if storageType&WithTimer > 0 {
		ss, err := NewTimerStorage(s, logger)
		if err != nil {
			log.Fatalf("Error creating the client: %s", err)
		}
		s = ss
	}

	return s
}
