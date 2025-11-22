package main

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "github-activity",
		Short: "A simple github events fetcher CLI",
	}
	
	var usernameCmd = &cobra.Command{
		Use: "username [github-username]",
		Short: "Fetch a user's events",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string){
			if err := fetchUserEvent(args[0]); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		},
	}

	rootCmd.AddCommand(usernameCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
