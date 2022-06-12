package handler

import (
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type IpaInstallHandler struct {
}

func NewIpaInstallHandler() *IpaInstallHandler {
	return &IpaInstallHandler{}
}

type getInstallIpaPlistArgs struct {
	FileName         string `form:"file_name" validate:"required"`
	Title            string `form:"title" validate:"required"`
	BundleIdentifier string `form:"bundle_identifier" validate:"required"`
	BundleVersion    string `form:"bundle_version" validate:"required"`
}

func (p *getInstallIpaPlistArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *IpaInstallHandler) GetInstallIpaPlistFile(w http.ResponseWriter, r *http.Request) {
	args := getInstallIpaPlistArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	plistFile := fmt.Sprintf(plistTemplate, args.FileName, args.FileName, args.FileName, args.Title, args.Title, args.BundleIdentifier, args.BundleVersion)

	w.Header().Add("Content-Type", "text/xml;charset=UTF-8")
	w.Write([]byte(plistFile))
}

var plistTemplate = `<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0"> 
	<dict> 
		<key>items</key> 
		<array> 
			<dict> 
				<key>assets</key> 
				<array> 
					<dict> 
						<key>kind</key> 
						<string>software-package</string> 
						<key>url</key> 
						<string>http://127.0.0.1:1995/%s.ipa</string> 
					</dict> 
					<dict> 
						<key>kind</key> 
						<string>display-image</string> 
						<key>url</key> 
						<string>http://127.0.0.1:1995/%s.png</string> 
						<key>needs-shine</key> 
						<true /> 
					</dict> 
					<dict> 
						<key>kind</key> 
						<string>full-size-image</string> 
						<key>url</key> 
						<string>http://127.0.0.1:1995/%s.png</string> 
						<key>needs-shine</key> 
						<true /> 
					</dict> 
				</array> 
				<key>metadata</key> 
				<dict> 
					<key>kind</key> 
					<string>software</string> 
					<key>subtitle</key> 
					<string>%s</string> 
					<key>title</key> 
					<string>%s</string> 
					<key>bundle-identifier</key> 
					<string>%s</string> 
					<key>bundle-version</key> 
					<string>%s</string> 
				</dict> 
			</dict> 
		</array> 
	</dict> 
</plist>`
