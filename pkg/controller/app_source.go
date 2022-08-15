package controller

import "context"

type AppSourceController interface {
	Insert(ctx context.Context, loginID int64, URL string) (int64, error)
	BatchGetAppSourceInfo(ctx context.Context, URLs []string) (map[string]*AppSourceInfo, error)
}

type AppSourceInfo struct {
	Name       string              `json:"name"`
	Message    string              `json:"message"`
	Identifier string              `json:"identifier"`
	SourceURL  string              `json:"sourceURL"`
	Sourceicon string              `json:"sourceicon"`
	PayURL     string              `json:"payURL"`
	UnlockURL  string              `json:"unlockURL"`
	Apps       []*AppSourceAppInfo `json:"apps"`

	/// 如果此字段不为空，说明是加密数据
	Appstore *string `json:"appstore,omitempty"`
}

type AppSourceAppInfo struct {
	Name               string      `json:"name"`
	Version            string      `json:"version"`
	VersionDate        string      `json:"versionDate"`
	VersionDescription string      `json:"versionDescription"`
	Lock               interface{} `json:"lock"`
	DownloadURL        string      `json:"downloadURL"`
	IsLanZouCloud      string      `json:"isLanZouCloud"`
	IconURL            string      `json:"iconURL"`
	TintColor          string      `json:"tintColor"`
	Size               string      `json:"size"`
}
