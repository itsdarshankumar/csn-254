package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/itsdarshankumar/CSN-254/services/master"
)

func main() {
	// Define the root command
	rootCmd := &cobra.Command{
		Use:   "kubepack",
		Short: "CLI for kubepack",
		Run:   rootFunc,
	}

	// Add subcommands
	rootCmd.AddCommand(master.DeployerMaster())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func rootFunc(cmd *cobra.Command, args []string){
	// function to test the root command
	fmt.Println("welcome to kubepack")
}