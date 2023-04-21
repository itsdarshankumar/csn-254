package master

import (
	"fmt"
	"os"
	"strconv"

	"github.com/itsdarshankumar/CSN-254/services/cluster"
	"github.com/itsdarshankumar/CSN-254/services/imagemaker"
	"github.com/spf13/cobra"
)

func DeployerMaster() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "deploy your source code",
		Run:   Deployer,
	}
	// Add a "path" flag to the "deploy" command
	cmd.Flags().String("name", "", "the name to the application")
	// Add a "path" flag to the "deploy" command
	cmd.Flags().String("path", "", "the path to the application")
	// Add a "service port" flag to the "deploy" command
	cmd.Flags().String("port", "", "the port to the application")

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
	portStr, err := cmd.Flags().GetString("port")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	portInt, err := strconv.ParseInt(portStr, 10, 32)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing port '%s': %v\n", portStr, err)
		os.Exit(1)
	}
	port := int32(portInt)
	// call the image maker function
	imagemaker.Imagemaker(path, name)
	// call the cluster function
	cluster.Cluster(name, port)
}
