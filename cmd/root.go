/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"fmt"
)

func newRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "jobcan",
		Short: "This is jobcan command",
		Long:  `This is jobcan comand long long description`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("hello jobcan")
			return nil
		},
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = newRootCmd()

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


