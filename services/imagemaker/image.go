package imagemaker

import (
	"fmt"
	"os/exec"
)

func Imagemaker(path string, appName string) {
	// build the OCI Image
	command := "~/pack build " + appName + " --path " + path + " --buildpack paketo-buildpacks/nodejs --builder paketobuildpacks/builder:base --publish"
	cmd := exec.Command(command)

	// Capture the output
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print the output
	fmt.Println(string(out))
}
