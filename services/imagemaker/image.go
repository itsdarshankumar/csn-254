package imagemaker

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
)

func Imagemaker(path string, appName string) {
	// create a new spinner with custom settings
	s := spinner.New(spinner.CharSets[20], 100*time.Millisecond)
	s.Suffix = " Building OCI Image..."
	s.Start()
	defer s.Stop()

	// build the OCI Image
	command := "~/pack build itsdarshankumar/" + appName + " --path " + path + " --buildpack paketo-buildpacks/nodejs --builder paketobuildpacks/builder:base --publish"
	cmd := exec.Command("zsh","-c",command)

	// Capture the output
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	// print success message
	fmt.Println("Image successfully built.....")
}
