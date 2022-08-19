package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"dumpapp_server/pkg/common/util"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/url"
	"time"
)

func main() {
	//accessID := "e36acdae35255f5f6169edf36ac318f7"
	//accessKey := "28e04c43963e3dbd6d6e38569d8b7eb3ed4212220b3b2c74f87c7b1ad5bfe786"
	accessID := "8080919c91fbd53df5404ab3645f5d17"
	accessKey := "b91d939a5c59ce7b1b205adf892af4fb1c3ba8c6e5155980538878df37cd96e6"

	URL := "https://api.lingshulian.com/api/auth/secret"

	bodyInfo := map[string]interface{}{
		"ttl": 900,
		"policy": []string{
			"full_control",
		},
		"bucket_name": "membersignipa",
	}
	model := "POST"

	urlInfo, err := url.Parse(URL)
	util.PanicIf(err)

	host := urlInfo.Host
	urlPath := urlInfo.Path

	bodyData, err := json.Marshal(bodyInfo)
	util.PanicIf(err)

	body := string(bodyData)

	expiryTo := time.Now().Unix() + 900

	signAccessSecret := fmt.Sprintf("%s-%s", accessID, accessKey)
	fmt.Println("signAccessSecret-->; ", signAccessSecret)

	signString := fmt.Sprintf("%s\n%s\n%s\n%s\n%d", model, host, urlPath, body, expiryTo)
	fmt.Println("signString--->: ", signString)

	sign := fmt.Sprintf("%s-%d-%s", accessID, expiryTo, Hmac(signString, signAccessSecret))
	fmt.Println("sign-->: ", sign)

	fmt.Println("uY2J8EUbRhmXwtaSentrqiXGzWM=")
	fmt.Println(Hmac("key", "data"))

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	// Create S3 service client
	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("myBucket"),
		Key:    aws.String("myKey"),
	})
	urlStr, err := req.Presign(15 * time.Minute)

}

func Hmac(key, data string) string {
	mac := hmac.New(sha1.New, []byte(data))
	mac.Write([]byte(key))
	res := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return res
}
