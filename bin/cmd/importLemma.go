package cmd

import (
	"fmt"
	"log"
	"os"
	"parser/app/lemma/repository"
	"parser/app/lemma/service"
	"parser/clients"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var importLemmaCmd = &cobra.Command{
	Use:   "importLemma",
	Short: "Импорт в эластик данных из файла",
	Long: `Более полное описание:
Ну когда нибудь позже я все тут опишу
`,
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		t := time.Now()

		es, err := clients.CreateElasticClient()
		if err != nil {
			log.Fatalf("Error creating the client: %s", err)
		}
		bi := clients.NewLemmaBulkIndexer(es)

		xmlFile := service.DownloadArchive()
		lemmaChan := service.StartImportToChan(xmlFile, &wg)

		fmt.Println("Запускаем горутины для эластика.")
		fmt.Println(time.Now().Sub(t))

		repository.BulkLemma(lemmaChan, bi)

		wg.Wait()
		fmt.Println("All goroutines complete.")
		fmt.Println(time.Now().Sub(t))

		fmt.Println("Удаляем XML " + xmlFile)
		os.Remove(xmlFile)
		fmt.Println("Удалили")
	},
}

func init() {
	rootCmd.AddCommand(importLemmaCmd)
}
