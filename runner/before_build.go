package runner

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os/exec"
)

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func beforeBuildExecuteBat() (string, bool) {
	if preBuildCmd() == "" {
		buildLog("No before_build.bat found")
		return "", false
	}

	buildCmdStr := fmt.Sprintf("cmd /C %s", preBuildCmd())

	buildLog("Before Building... " + buildCmdStr)

	// cmd := exec.Command("go", "build", "-o", buildPath(), root())
	cmd := exec.Command("cmd", "/C", preBuildCmd())
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	utf8, _ := GbkToUtf8([]byte(out.String()))
	buildLog(string(utf8))
	if stderr.String() != "" {
		gbkToUtf8, _ := GbkToUtf8([]byte(stderr.String()))
		buildLog(string(gbkToUtf8))
	}
	if err != nil {
		fatal(err)
	}
	buildLog("Before Building... " + buildCmdStr + " done")
	if err != nil {
		return out.String(), false
	}

	return "", true
}
