package render

import (
	"context"
	"fmt"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	errors2 "dumpapp_server/pkg/common/errors"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	impl2 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
	pkgErr "github.com/pkg/errors"
	"github.com/spf13/cast"
)

type Member struct {
	meta *models.Account

	ID     int64            `json:"id,string"`
	Email  string           `json:"email"`
	Status string           `json:"status"`
	Role   enum.AccountRole `json:"role"`

	Phone    *string `json:"phone,omitempty" render:"method=RenderPhone"`
	Password *string `json:"password,omitempty" render:"method=RenderPassword"`

	PayCount *int64 `json:"pay_count,omitempty" render:"method=RenderPayCount"`

	/// 所有价格明细
	DispenseInfo    *DispenseInfo    `json:"dispense_info" render:"method=RenderDispenseInfo"`       /// 分发券价格
	IpaInfo         *IpaInfo         `json:"ipa_info" render:"method=RenderIpaInfo"`                 /// ipa 包价格
	CertificateInfo *CertificateInfo `json:"certificate_info" render:"method=RenderCertificateInfo"` /// 证书价格

	/// 邀请链接
	InviteURL *string `json:"invite_url,omitempty" render:"method=RenderInviteURL"`
	/// 用户绑定的设备信息
	Devices []*Device `json:"devices,omitempty" render:"method=RenderDevices"`

	Token *string `json:"token,omitempty" render:"method=RenderToken"`

	/// 分享信息
	ShareInfo *ShareInfo `json:"share_info"`

	/// 充值活动
	PayCampaign *PayCampaign `json:"pay_campaign"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`

	/// Admin 相关
	Admin *Admin `json:"admin,omitempty" render:"method=RenderAdmin"`
}

type PayCampaign struct {
	Description string `json:"description"`
}

type DispenseInfo struct {
	Count int64  `json:"count"`
	Rule  string `json:"rule"`
	Ratio int64  `json:"ratio"`
}

type IpaInfo struct {
	Price int64 `json:"price"`
}

type CertificateInfo struct {
	Prices []*controller.CertificatePriceInfo `json:"prices"`
}

type MemberRender struct {
	ids           []int64
	loginID       int64
	includeFields []string

	memberMap map[int64]*Member

	accountDAO            dao.AccountDAO
	memberPayCountDAO     dao.MemberPayCountDAO
	memberInviteCodeDAO   dao.MemberInviteCodeDAO
	memberDeviceDAO       dao.MemberDeviceDAO
	memberIDEncryptionDAO dao.MemberIDEncryptionDAO
	dispenseCountDAO      dao.DispenseCountDAO

	certificatePriceCtl controller.CertificatePriceController
	memberPayOrderCtl   controller.MemberPayOrderController
}

type MemberOption func(*MemberRender)

func MemberIncludes(fields []string) MemberOption {
	return func(render *MemberRender) {
		fields = append(fields, defaultFields...)
		uniqFields := make([]string, 0)
		fieldSet := util2.NewSet()
		for _, field := range fields {
			if fieldSet.Exists(field) {
				continue
			}
			fieldSet.Add(field)
			uniqFields = append(uniqFields, field)
		}
		render.includeFields = uniqFields
	}
}

var DefaultFields = []string{
	"PayCount",
	"DispenseInfo",
	"InviteURL",
	"IpaInfo",
	"CertificateInfo",
}

var MemberAdminRenderFields = []MemberOption{
	MemberIncludes(append(DefaultFields, []string{"Admin", "Password", "Phone", "Token"}...)),
}

var MemberDefaultRenderFields = []MemberOption{
	MemberIncludes(DefaultFields),
}

var MemberSelfRenderFields = []MemberOption{
	MemberIncludes(append(DefaultFields, []string{"Token", "Phone"}...)),
}

func NewMemberRender(ids []int64, loginID int64, opts ...MemberOption) *MemberRender {
	f := &MemberRender{
		ids:     ids,
		loginID: loginID,

		accountDAO:            impl.DefaultAccountDAO,
		memberPayCountDAO:     impl.DefaultMemberPayCountDAO,
		memberInviteCodeDAO:   impl.DefaultMemberInviteCodeDAO,
		memberDeviceDAO:       impl.DefaultMemberDeviceDAO,
		memberIDEncryptionDAO: impl.DefaultMemberIDEncryptionDAO,
		dispenseCountDAO:      impl.DefaultDispenseCountDAO,

		certificatePriceCtl: impl2.DefaultCertificatePriceController,
		memberPayOrderCtl:   impl2.DefaultMemberPayOrderController,
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func (f *MemberRender) RenderSlice(ctx context.Context) []*Member {
	tMap := f.RenderMap(ctx)
	tSlice := make([]*Member, len(f.ids))
	for i, id := range f.ids {
		tSlice[i] = tMap[id]
	}
	return tSlice
}

func (f *MemberRender) RenderMap(ctx context.Context) map[int64]*Member {
	if len(f.ids) == 0 {
		return f.memberMap
	}

	f.fetch(ctx)

	err := autoRender(ctx, f, Member{}, f.includeFields)
	if err != nil {
		panic(err)
	}

	return f.memberMap
}

func (f *MemberRender) fetch(ctx context.Context) {
	accountMap, err := f.accountDAO.BatchGet(ctx, f.ids)
	util.PanicIf(err)

	res := make(map[int64]*Member)
	for _, account := range accountMap {
		res[account.ID] = &Member{
			meta:      account,
			ID:        account.ID,
			Email:     account.Email,
			Status:    "normal",
			Role:      account.Role,
			CreatedAt: account.CreatedAt.Unix(),
			UpdatedAt: account.UpdatedAt.Unix(),
			PayCampaign: &PayCampaign{
				Description: f.memberPayOrderCtl.GetPayCampaignDescription(),
			},
		}
	}
	f.memberMap = res

	f.RenderShareInfo(ctx)
}

func (f *MemberRender) RenderPhone(ctx context.Context) {
	for _, member := range f.memberMap {
		if cast.ToString(member.meta.ID) == member.meta.Phone { /// 老用户没有设置手机号
			continue
		}
		member.Phone = util.StringPtr(member.meta.Phone)
	}
}

func (f *MemberRender) RenderPassword(ctx context.Context) {
	for _, member := range f.memberMap {
		member.Password = util.StringPtr(member.meta.Password)
	}
}

func (f *MemberRender) RenderPayCount(ctx context.Context) {
	countMap, err := f.memberPayCountDAO.BatchGetMemberNormalCount(ctx, f.ids)
	util.PanicIf(err)
	for _, member := range f.memberMap {
		member.PayCount = util2.Int64Ptr(countMap[member.ID])
	}
}

func (f *MemberRender) RenderDispenseInfo(ctx context.Context) {
	countMap, err := f.dispenseCountDAO.BatchGetMemberNormalCount(ctx, f.ids)
	util.PanicIf(err)
	for _, member := range f.memberMap {
		member.DispenseInfo = &DispenseInfo{
			Count: countMap[member.ID],
			Rule:  fmt.Sprintf("1D 币兑换 %d 分发劵。\n（分发劵用于针对签名后的 APP生成下载链接，当前 1G 以下 APP上传分发后安装每次消耗 1 分发劵，1G 以上每次消耗 2 分发劵）", constant.DispenseRatioByPayCount),
			Ratio: constant.DispenseRatioByPayCount,
		}
	}
}

func (f *MemberRender) RenderIpaInfo(ctx context.Context) {
	for _, member := range f.memberMap {
		member.IpaInfo = &IpaInfo{
			Price: constant.IpaPrice,
		}
	}
}

func (f *MemberRender) RenderCertificateInfo(ctx context.Context) {
	memberIDs := make([]int64, 0)
	for _, member := range f.memberMap {
		memberIDs = append(memberIDs, member.ID)
	}
	memberIDs = util2.RemoveDuplicates(memberIDs)

	priceMap, err := f.certificatePriceCtl.BatchGetPrices(ctx, memberIDs)
	util.PanicIf(err)

	for _, member := range f.memberMap {
		member.CertificateInfo = &CertificateInfo{
			Prices: priceMap[member.ID],
		}
	}
}

func (f *MemberRender) RenderInviteURL(ctx context.Context) {
	inviteCode, err := f.memberInviteCodeDAO.GetByMemberID(ctx, f.loginID)
	if err != nil && pkgErr.Cause(err) != errors2.ErrNotFound {
		util.PanicIf(err)
	}

	inviteCodeString := ""

	/// 如果邀请码已经存在, 则直接取出即可
	if inviteCode != nil {
		inviteCodeString = inviteCode.Code
	}

	/// 如果没有邀请码则生成邀请码
	if inviteCode == nil {
		inviteCodeString = util2.MustGenerateCode(ctx, 10)
		util.PanicIf(f.memberInviteCodeDAO.Insert(ctx, &models.MemberInviteCode{
			MemberID: f.loginID,
			Code:     inviteCodeString,
		}))
	}

	for _, member := range f.memberMap {
		if member.ID == f.loginID {
			member.InviteURL = util.StringPtr(fmt.Sprintf(constant.InviteURL, inviteCodeString))
		}
	}
}

func (f *MemberRender) RenderDevices(ctx context.Context) {
	/// 非主用户不加载
	_, ok := f.memberMap[f.loginID]
	if !ok {
		return
	}

	/// 获取用户所有设备
	memberDeviceMap, err := f.memberDeviceDAO.BatchGetByMemberIDs(ctx, []int64{f.loginID})
	util.PanicIf(err)

	devices, ok := memberDeviceMap[f.loginID]
	if !ok {
		return
	}

	/// 获取所有设备 id
	deviceIDs := make([]int64, 0)
	for _, device := range devices {
		deviceIDs = append(deviceIDs, device.ID)
	}

	deviceMap := NewDeviceRender(deviceIDs, f.loginID, DeviceDefaultRenderFields...).RenderMap(ctx)
	deviceResult := make([]*Device, 0)
	for _, deviceID := range deviceIDs {
		device, ok := deviceMap[deviceID]
		if !ok {
			continue
		}
		deviceResult = append(deviceResult, device)
	}
	f.memberMap[f.loginID].Devices = deviceResult
}

func (f *MemberRender) RenderShareInfo(ctx context.Context) {
	for _, member := range f.memberMap {
		member.ShareInfo = MustRenderShareInfo(ctx, member.ID)
	}
}

func (f *MemberRender) RenderAdmin(ctx context.Context) {
	adminMap := NewAdminRender(f.ids, f.loginID).RenderMap(ctx)
	for _, member := range f.memberMap {
		member.Admin = adminMap[member.ID]
	}
}

func (f *MemberRender) RenderToken(ctx context.Context) {
	data, err := f.memberIDEncryptionDAO.BatchGetByMemberID(ctx, f.ids)
	util.PanicIf(err)
	for _, member := range f.memberMap {
		d, ok := data[member.meta.ID]
		if !ok {
			continue
		}
		member.Token = util.StringPtr(d.Code)
	}
}
