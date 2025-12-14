package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "inventaris",
	Short: "Aplikasi Inventaris Sederhana dengan Golang CLI",
}

func Execute(){
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}