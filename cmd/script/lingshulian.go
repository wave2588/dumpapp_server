package main

import (
	"bytes"
	"dumpapp_server/pkg/common/util"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	content, err := ioutil.ReadFile("/Users/wave/Documents/Private/dump/dumpapp_server/cmd/script/test.ipa")
	util.PanicIf(err)

	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(
			"8080919c91fbd53df5404ab3645f5d17",
			"b91d939a5c59ce7b1b205adf892af4fb1c3ba8c6e5155980538878df37cd96e6",
			"",
		),
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("s3-us-east-1.ossfiles.com"),
	})
	util.PanicIf(err)

	// Create S3 service client
	svc := s3.New(sess)

	resp, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String("membersignipa"),
		Key:    aws.String("1122223.test"),
	})
	//resp.HTTPRequest.Header.Set("Content-MD5", md5s)
	url, err := resp.Presign(15 * time.Minute)
	if err != nil {
		fmt.Println("error presigning request:", err)
		return
	}
	fmt.Println(111, url)

	//req, err := http.NewRequest("PUT", url, strings.NewReader(fmt.Sprintf(constant.DeviceMobileConfig, "22221122")))
	req, err := http.NewRequest("PUT", url, bytes.NewReader(content))
	if err != nil {
		fmt.Println("error creating request", url)
		return
	}

	defClient, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("http.DefaultClient err", err)
		return
	}
	ss, _ := json.Marshal(defClient.Body)
	fmt.Println(string(ss))
}
