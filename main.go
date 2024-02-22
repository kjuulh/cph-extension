package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	cmd := cobra.Command{
		Use: "cph-extension",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("Hello, cph-extension")

			return nil
		},
	}

	cmd.AddCommand(&cobra.Command{
		Use: "one-more-command",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("one more")
			return nil
		},
	})

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
