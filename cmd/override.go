package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var go_up = &cobra.Command{
	Use: "goup",
	Short: "A simple Golang manager",
	Long: "goup is a simple tool for managing Golang installs similar to the rustup utility provided by rust-lang",
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := go_up.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}