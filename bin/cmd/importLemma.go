package cmd

import (
	"fmt"
	"log"
	"parser/app/lemma/repository"
	"parser/app/lemma/service"
	"parser/clients"
	"parser/config"
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
		config.CalculatetConfig()

		var wg sync.WaitGroup
		t := time.Now()

		es, err := clients.CreateElasticClient()
		if err != nil {
			log.Fatalf("Error creating the client: %s", err)
		}

		lemmaChan := service.StartImportToChan(&wg)

		fmt.Println("Запускаем горутины для эластика.")
		fmt.Println(time.Now().Sub(t))

		repository.BulkLemma(lemmaChan, es)

		wg.Wait()
		fmt.Println("All goroutines complete.")
		fmt.Println(time.Now().Sub(t))

	},
}

func init() {
	rootCmd.AddCommand(importLemmaCmd)
}
