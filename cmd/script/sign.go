package main

import (
	"bytes"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/util"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	//ctx := context.Background()

	path, err := os.Getwd()
	util.PanicIf(err)

	/// 证书文件
	serverCrt := fmt.Sprintf("%s/templates/sign/server.crt", path)
	serverKey := fmt.Sprintf("%s/templates/sign/server.key", path)
	caCrt := fmt.Sprintf("%s/templates/sign/ca.crt", path)

	configURL := strings.ReplaceAll(constant.DeviceMobileConfig, "%s", "https://xxxxx")
	/// 签名前的 mobileconfig
	mobileconfigInPath := fmt.Sprintf("%s/templates/mobileconfig/unsign.mobileconfig", path)
	util.PanicIf(ioutil.WriteFile(mobileconfigInPath, []byte(configURL), 0644))

	/// 签名后的 mobileconfig
	mobileconfigOutPath := fmt.Sprintf("%s/templates/mobileconfig/signed.mobileconfig", path)

	cmdString := fmt.Sprintf("openssl smime -sign -in %s -out %s -signer %s -inkey %s  -certfile %s -outform der -nodetach", mobileconfigInPath, mobileconfigOutPath, serverCrt, serverKey, caCrt)
	util.PanicIf(cmd(cmdString))
}

func cmd(input string) error {
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd := exec.Command("/bin/bash", "-c", input)
	stdoutIn, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderrIn, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err = cmd.Start()
	if err != nil {
		return err
	}
	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()
	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()
	err = cmd.Wait()
	if err != nil {
		return err
	}
	if errStdout != nil || errStderr != nil {
		return errors.New(fmt.Sprintf("%s\n%s", errStdout.Error(), errStderr.Error()))
	}
	return nil
}
