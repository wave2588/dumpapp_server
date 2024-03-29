package handler

import (
	"fmt"
	"net/http"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/datatype"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/config"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	impl2 "dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	util3 "dumpapp_server/pkg/util"
	"github.com/go-playground/validator/v10"
)

type FileHandler struct {
	lingshulianCtl controller.LingshulianController
	fileCtl        controller.FileController
	fileDAO        dao.FileDAO
}

func NewFileHandler() *FileHandler {
	return &FileHandler{
		lingshulianCtl: impl.DefaultLingshulianController,
		fileCtl:        impl.DefaultFileController,
		fileDAO:        impl2.DefaultFileDAO,
	}
}

type postPlistFileArgs struct {
	AppIcon     string `json:"app_icon" validate:"required"`
	AppName     string `json:"app_name" validate:"required"`
	AppVersion  string `json:"app_version" validate:"required"`
	AppURL      string `json:"app_url" validate:"required"`
	AppBundleID string `json:"app_bundle_id" validate:"required"`
}

func (p *postPlistFileArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

func (h *FileHandler) CreatePlistFile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := &postPlistFileArgs{}
	util.PanicIf(util.JSONArgs(r, args))

	plistID := util3.MustGenerateID(ctx)
	token := fmt.Sprintf("%d.plist", plistID)

	bucket := config.DumpConfig.AppConfig.LingshulianMemberSignIpaBucket

	util.PanicIf(h.fileDAO.Insert(ctx, &models.File{
		Token:    token,
		IsDelete: false,
		BizExt: datatype.FileBizExt{
			Bucket: bucket,
		},
	}))

	util.PanicIf(h.fileCtl.PutFileToLocal(ctx, h.fileCtl.GetPlistFolderPath(ctx), token, []byte(fmt.Sprintf(constant.PlistFileConfig, args.AppURL, args.AppIcon, args.AppIcon, args.AppBundleID, args.AppVersion, args.AppName))))

	util.RenderJSON(w, map[string]interface{}{
		"plist_url": h.fileCtl.GetPlistFileURL(ctx, token),
	})
}

func (h *FileHandler) CheckPlistFileExist(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	key := util.URLParam(r, "key")

	util.RenderJSON(w, map[string]bool{
		"is_exist": h.fileCtl.CheckPlistFileExist(ctx, key),
	})
}
