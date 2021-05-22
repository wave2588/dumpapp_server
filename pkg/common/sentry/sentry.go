package sentry

import (
	"log"
	"time"

	"dumpapp_server/pkg/config"
	"github.com/getsentry/sentry-go"
)

func init() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: config.DumpConfig.AppConfig.SentryDSN,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)
}

func RavenCaptureError(err error) {
	sentry.CaptureException(err)
}
