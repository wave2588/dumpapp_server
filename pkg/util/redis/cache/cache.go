package cache

import (
	"time"

	"dumpapp_server/pkg/util/redis/store"
)

type MemCache struct {
	store             store.Store
	expire            time.Duration
	fallbackWhenError bool
	limiter           IRateLimiter
}
