package sentry

import (
	"log"
	"time"

	"github.com/getsentry/sentry-go"
)

func init() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://ef12941ca92f43eca079ca38fe70c62c@o1388610.ingest.sentry.io/6711092",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("It works!")
}

func RavenCaptureError(err error) {
	sentry.CaptureException(err)
}
