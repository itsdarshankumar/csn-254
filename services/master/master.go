package master

import (
	"fmt"
	"os"

	"github.com/itsdarshankumar/CSN-254/services/imagemaker"
	"github.com/spf13/cobra"
)

func DeployerMaster() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "deploy your source code",
		Run:   Deployer,
	}
	// Add a "path" flag to the "run" command
	cmd.Flags().String("name", "", "the name to the application")
	// Add a "path" flag to the "run" command
	cmd.Flags().String("path", "", "the path to the application")

	return cmd
}

func Deployer(cmd *cobra.Command, args []string) {
	path, err := cmd.Flags().GetString("path")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	imagemaker.Imagemaker(path, name)
}
