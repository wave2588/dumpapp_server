package util

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"dumpapp_server/pkg/common/util"
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

var (
	node, nodeErr = snowflake.NewNode(1)
	u             = uuid.New()
)

func MustGenerateID(ctx context.Context) int64 {
	util.PanicIf(nodeErr)
	return node.Generate().Int64()
}

func MustGenerateUUID() string {
	return u.String()
}

func MustGenerateCaptcha(ctx context.Context) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}

func MustGenerateCode(ctx context.Context, l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return cast.ToString(result)
}

func MustGenerateAppCDKEY() string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return fmt.Sprintf("DP%s", cast.ToString(result))
}
