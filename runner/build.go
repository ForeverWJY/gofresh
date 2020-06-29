package runner

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

func build() (string, bool) {

	buildCmdStr := fmt.Sprintf("go build -o %s %s", buildPath(), buildFile())

	buildLog("Building... "+buildCmdStr)

	// cmd := exec.Command("go", "build", "-o", buildPath(), root())
	cmd := exec.Command("go", "build", "-o", buildPath(), buildFile())

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fatal(err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		fatal(err)
	}

	io.Copy(os.Stdout, stdout)
	errBuf, _ := ioutil.ReadAll(stderr)

	err = cmd.Wait()
	if err != nil {
		return string(errBuf), false
	}

	return "", true
}
