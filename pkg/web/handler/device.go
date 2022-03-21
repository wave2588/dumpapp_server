package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"dumpapp_server/pkg/common/constant"
	errors2 "dumpapp_server/pkg/common/errors"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/errors"
	util2 "dumpapp_server/pkg/util"
	controller2 "dumpapp_server/pkg/web/controller"
	impl3 "dumpapp_server/pkg/web/controller/impl"
	xj "github.com/basgys/goxml2json"
	"github.com/go-playground/validator/v10"
	pkgErr "github.com/pkg/errors"
	"github.com/skip2/go-qrcode"
)

type DeviceHandler struct {
	accountDAO            dao.AccountDAO
	memberDeivceDAO       dao.MemberDeviceDAO
	memberIDEncryptionCtl controller.MemberIDEncryptionController
	signWebCtl            controller2.SignMobileconfigWebController
}

func NewDeviceHandler() *DeviceHandler {
	return &DeviceHandler{
		accountDAO:            impl.DefaultAccountDAO,
		memberDeivceDAO:       impl.DefaultMemberDeviceDAO,
		memberIDEncryptionCtl: impl2.DefaultMemberIDEncryptionController,
		signWebCtl:            impl3.DefaultSignMobileconfigWebController,
	}
}

/// 获取下载描述文件二维码
func (h *DeviceHandler) GetMobileConfigQRCode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	loginID := mustGetLoginID(ctx)
	code, err := h.memberIDEncryptionCtl.GetCodeByMemberID(ctx, loginID)
	util.PanicIf(err)

	url := fmt.Sprintf("%s/device/config/file?code=%s", constant.HOST, code)
	q, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		return
	}
	png, err := q.PNG(256)
	if err != nil {
		return
	}
	util.RenderJSON(w, map[string]interface{}{
		"image_base64": base64.StdEncoding.EncodeToString(png),
		"url":          url,
	})
	// w.Header().Set("Content-Type", "image/png")
	// w.Header().Set("Content-Length", fmt.Sprintf("%d", len(png)))
	// w.Write(png)
}

type getMobileConfigFileArgs struct {
	Code string `form:"code" validate:"required"`
}

func (p *getMobileConfigFileArgs) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return errors.UnproccessableError(fmt.Sprintf("参数校验失败: %s", err.Error()))
	}
	return nil
}

/// 实际获取描述文件接口
func (h *DeviceHandler) GetMobileConfigFile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	args := getMobileConfigFileArgs{}
	util.PanicIf(formDecoder.Decode(&args, r.URL.Query()))
	util.PanicIf(args.Validate())

	content, err := h.signWebCtl.Sign(ctx, args.Code)
	util.PanicIf(err)

	w.Header().Add("Content-Type", "application/x-apple-aspen-config; chatset=utf-8")
	w.Header().Add("Content-Disposition", "attachment;filename=\"1.mobileconfig\"")
	w.Write(content)
}

type device struct {
	Device *deviceInfo `json:"dict" validate:"required"`
}

type deviceInfo struct {
	Keys   []string `json:"key" validate:"required"`
	Values []string `json:"string" validate:"required"`
}

func (h *DeviceHandler) Bind(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	code := util.URLParam(r, "code")
	memberID, err := h.memberIDEncryptionCtl.GetMemberIDByCode(ctx, code)
	util.PanicIf(err)

	s, err := ioutil.ReadAll(r.Body)
	util.PanicIf(err)
	content := string(s)
	index := strings.Index(content, "<dict>")
	index2 := strings.Index(content, "</dict>")
	/// 验证是否找对了 dict 的下标
	if index == 0 || index2 == 0 {
		panic(errors.UnproccessableError("设备认证失败。"))
	}
	c := content[index : index2+len("</dict>")]

	de, err := xj.Convert(strings.NewReader(c))
	util.PanicIf(err)

	args := &device{}
	util.PanicIf(json.Unmarshal(de.Bytes(), args))

	if args.Device == nil {
		panic(errors.UnproccessableError("解析 device 失败。"))
	}

	keys := args.Device.Keys
	values := args.Device.Values

	/// 验证 keys 和 values 长度是否一致
	if len(keys) != len(values) {
		panic(errors.UnproccessableError("device key value 不一致"))
	}

	id := util2.MustGenerateID(ctx)
	memberDevice := &models.MemberDevice{
		ID:       id,
		MemberID: memberID,
	}
	for idx, key := range args.Device.Keys {
		value := values[idx]
		switch key {
		case "IMEI":
			memberDevice.Imei = value
		case "PRODUCT":
			memberDevice.Product = value
		case "UDID":
			memberDevice.Udid = value
		case "VERSION":
			memberDevice.Version = value
		default:
			panic(errors.UnproccessableError(fmt.Sprintf("发现未识别的 key。 key = %s", key)))
		}
	}

	md, err := h.memberDeivceDAO.GetByUdid(ctx, memberDevice.Udid)
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		util.PanicIf(err)
	}
	if md != nil {
		account, err := h.accountDAO.Get(ctx, md.MemberID)
		util.PanicIf(err)
		w.Header().Set("Location", fmt.Sprintf("https://dumpapp.com/view_udid?udid=%s&product=%s&version=%s&exist_email=%s", memberDevice.Udid, memberDevice.Product, memberDevice.Version, account.Email))
		w.WriteHeader(301)
		return
	}
	util.PanicIf(h.memberDeivceDAO.Insert(ctx, memberDevice))

	w.Header().Set("Location", fmt.Sprintf("https://dumpapp.com/view_udid?udid=%s&product=%s&version=%s", memberDevice.Udid, memberDevice.Product, memberDevice.Version))
	w.WriteHeader(301)
}
