package main

import (
	"crypto/md5"
	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/util"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
	"strings"
	"time"
)

func main() {

	h := md5.New()
	content := strings.NewReader(fmt.Sprintf(constant.DeviceMobileConfig, "222222"))
	content.WriteTo(h)

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
	md5s := base64.StdEncoding.EncodeToString(h.Sum(nil))
	//resp.HTTPRequest.Header.Set("Content-MD5", md5s)
	url, err := resp.Presign(15 * time.Minute)
	if err != nil {
		fmt.Println("error presigning request:", err)
		return
	}
	fmt.Println(111, url)

	req, err := http.NewRequest("PUT", url, strings.NewReader(fmt.Sprintf(constant.DeviceMobileConfig, "22221122")))
	if err != nil {
		fmt.Println("error creating request", url)
		return
	}
	req.Header.Set("Content-MD5", md5s)

	defClient, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("http.DefaultClient err", err)
		return
	}
	fmt.Println(defClient)
}
