package impl

import (
	"context"
	"dumpapp_server/pkg/common/util"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type LingshulianController struct {
}

var DefaultLingshulianController *LingshulianController

func init() {
	DefaultLingshulianController = NewLingshulianController()
}

func NewLingshulianController() *LingshulianController {
	//secretId：8080919c91fbd53df5404ab3645f5d17
	//secretKey：b91d939a5c59ce7b1b205adf892af4fb1c3ba8c6e5155980538878df37cd96e6

	return &LingshulianController{}
}

func (c *LingshulianController) PutMemberSignIpa(ctx context.Context, name string, data string) error {
	//reader := strings.NewReader(data)
	//_, err := c.signIpaClient.Object.Put(ctx, name, reader, nil)
	//return err
	return nil
}

func (c *LingshulianController) GetMemberSignIpa(ctx context.Context, ipaToken string) (string, error) {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				"8080919c91fbd53df5404ab3645f5d17",
				"b91d939a5c59ce7b1b205adf892af4fb1c3ba8c6e5155980538878df37cd96e6",
				"",
			),
		),
		config.WithRegion("us-east-1"),
		config.WithEC2IMDSEndpoint("https://s3-us-east-1.ossfiles.com/"),
	)
	if err != nil {
		return "", err
	}

	client := s3.NewFromConfig(cfg)

	res, err := client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: util.StringPtr("membersignipa"),
		Key:    util.StringPtr(ipaToken),
	})
	if err != nil {
		return "", err
	}

	fmt.Println(res.ContentType)

	return "", nil
}

func (c *LingshulianController) DeleteMemberSignIpa(ctx context.Context, tokenPath string) error {
	return nil
}

func HttpRequest(method, endpoint string, header, values map[string]string, timeout time.Duration) ([]byte, error) {
	data := url.Values{}
	for key, value := range values {
		data.Set(key, value)
	}
	client := &http.Client{}
	if timeout != 0 {
		client.Timeout = timeout
	}
	r, err := http.NewRequest(method, endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		r.Header.Add(key, value)
	}
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
