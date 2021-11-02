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

	//openssl smime -sign -in Example.mobileconfig -out SignedVerifyExample.mobileconfig -signer InnovCertificates.pem -certfile root.crt.pem -outform der -nodetach

	path, err := os.Getwd()
	util.PanicIf(err)

	/// 证书文件
	pemPath := fmt.Sprintf("%s/templates/pem/sign.pem", path)

	configURL := strings.ReplaceAll(constant.DeviceMobileConfig, "%s", "https://xxxxx")
	/// 签名前的 mobileconfig
	mobileconfigInPath := fmt.Sprintf("%s/templates/mobileconfig/sign_in.mobileconfig", path)
	util.PanicIf(ioutil.WriteFile(mobileconfigInPath, []byte(configURL), 0644))

	/// 签名后的 mobileconfig
	mobileconfigOutPath := fmt.Sprintf("%s/templates/mobileconfig/sign_out.mobileconfig", path)

	cmdString := fmt.Sprintf("openssl smime -sign -in %s -out %s -signer %s -certfile %s -outform der -nodetach", mobileconfigInPath, mobileconfigOutPath, pemPath, pemPath)
	fmt.Println(cmdString)
	util.PanicIf(cmd(cmdString))

	//openssl smime -sign -in unsigned.mobileconfig -out signed.mobileconfig -signer mbaike.crt -inkey mbaikenopass.key -certfile ca-bundle.pem -outform der -nodetach

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
