package util

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// RenderJSON render data to client in json format.
func RenderJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type ListOutput struct {
	Paging *OutputPaging `json:"paging"`
	Data   interface{}   `json:"data"`
}

// OutputOffset is the
type OutputPaging struct {
	IsEnd    bool   `json:"is_end"`
	IsStart  bool   `json:"is_start"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Totals   int    `json:"totals"`
}

func GenerateOffsetPaging(ctx context.Context, r *http.Request, totals int, currentOffset, limit int) *OutputPaging {
	prevQuery := r.URL.Query()
	if currentOffset > limit {
		prevQuery.Set("offset", strconv.Itoa(currentOffset-limit))
	} else {
		prevQuery.Set("offset", "0")
	}
	prevQuery.Set("limit", strconv.Itoa(limit))

	nextQuery := r.URL.Query()
	nextQuery.Set("offset", strconv.Itoa(currentOffset+limit))
	nextQuery.Set("limit", strconv.Itoa(limit))

	return &OutputPaging{
		IsEnd:    currentOffset+limit >= totals,
		IsStart:  currentOffset == 0,
		Next:     GenerateResourceURL(ctx, r.URL.Path, nextQuery.Encode()),
		Previous: GenerateResourceURL(ctx, r.URL.Path, prevQuery.Encode()),
		Totals:   totals,
	}
}

func GenerateResourceURL(ctx context.Context, path string, query string) string {
	prefix := GetRequestContext(ctx, "request_uri_prefix")
	host := GetRequestContext(ctx, "request_host")
	scheme := GetRequestContext(ctx, "request_protocol")
	return (&url.URL{
		Scheme:   scheme,
		Host:     host,
		Path:     prefix + path,
		RawQuery: query,
	}).String()
}
