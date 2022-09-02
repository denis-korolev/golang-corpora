package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var importLemmaCmd = &cobra.Command{
	Use:   "importLemma",
	Short: "Импорт в эластик данных из файла",
	Long: `Более полное описание:
Ну когда нибудь позже я все тут опишу
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("importLemma called")
	},
}

func init() {
	rootCmd.AddCommand(importLemmaCmd)
}
