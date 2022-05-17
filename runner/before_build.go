package runner

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os/exec"
	"strings"
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
	errStr := stderr.String()
	utf8Str := string(utf8)
	buildLog(utf8Str)
	if stderr.String() != "" {
		//gbkToUtf8, _ := GbkToUtf8([]byte(stderr.String()))
		//buildLog(string(gbkToUtf8))
		buildLog(errStr)
		//fatal(errors.New("before Building Failed"))
	}
	if err != nil {
		buildLog("error:" + err.Error())
		fatal(err)
	}
	if strings.Contains(utf8Str, "before build failed") || strings.Contains(errStr, "before build failed") {
		//buildLog("before build failed")
		fatal(errors.New("before build failed"))
	}
	buildLog("Before Building... " + buildCmdStr + " done")
	if err != nil {
		return out.String(), false
	}

	return "", true
}
