package util

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/snowflake"
)

func MustGenerateID(ctx context.Context) int64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	return node.Generate().Int64()
}

func MustGenerateCaptcha(ctx context.Context) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}
