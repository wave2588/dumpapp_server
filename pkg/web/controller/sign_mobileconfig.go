package controller

import "context"

type SignMobileconfigWebController interface {
	Sign(ctx context.Context, memberCode string) ([]byte, error)
}
