package render

import (
	"context"
	"fmt"
	"time"

	errors2 "dumpapp_server/pkg/common/errors"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
	pkgErr "github.com/pkg/errors"
)

type Member struct {
	meta *models.Account

	ID     int64   `json:"id,string"`
	Email  string  `json:"email"`
	Status string  `json:"status"`
	Phone  *string `json:"phone,omitempty"`

	/// 可下载次数
	DownloadCount int64 `json:"download_count" render:"method=RenderDownloadCount"`
	/// todo: 已经没有这个东西了
	Vip *Vip `json:"vip,omitempty" render:"method=RenderMemberVip"`
	/// 邀请链接
	InviteURL *string `json:"invite_url" render:"method=RenderInviteURL"`
	/// 用户绑定的设备信息
	Devices []*Device `json:"devices" render:"method=RenderDevices"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`

	/// Admin 相关
	Admin *Admin `json:"admin,omitempty" render:"method=RenderAdmin"`
}

type Device struct {
	UDID string `json:"udid"`
}

type Vip struct {
	IsVip   bool   `json:"is_vip"`
	StartAt *int64 `json:"start_at,omitempty"`
	EndAt   *int64 `json:"end_at,omitempty"`
}

type MemberRender struct {
	ids           []int64
	loginID       int64
	includeFields []string

	memberMap map[int64]*Member

	accountDAO              dao.AccountDAO
	memberVipDAO            dao.MemberVipDAO
	memberDownloadNumberDAO dao.MemberDownloadNumberDAO
	memberInviteCodeDAO     dao.MemberInviteCodeDAO
	memberDeviceDAO         dao.MemberDeviceDAO
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

var MemberAdminRenderFields = []MemberOption{
	MemberIncludes([]string{
		"DownloadCount",
		"Vip",
		"Admin",
	}),
}

var MemberDefaultRenderFields = []MemberOption{
	MemberIncludes([]string{
		"DownloadCount",
		"Vip",
		"InviteURL",
		"Devices",
	}),
}

func NewMemberRender(ids []int64, loginID int64, opts ...MemberOption) *MemberRender {
	f := &MemberRender{
		ids:     ids,
		loginID: loginID,

		accountDAO:              impl.DefaultAccountDAO,
		memberVipDAO:            impl.DefaultMemberVipDAO,
		memberDownloadNumberDAO: impl.DefaultMemberDownloadNumberDAO,
		memberInviteCodeDAO:     impl.DefaultMemberInviteCodeDAO,
		memberDeviceDAO:         impl.DefaultMemberDeviceDAO,
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
			Phone:     util.StringPtr(account.Phone),
			CreatedAt: account.CreatedAt.Unix(),
			UpdatedAt: account.UpdatedAt.Unix(),
		}
	}
	f.memberMap = res
}

func (f *MemberRender) RenderMemberVip(ctx context.Context) {
	memberVipMap, err := f.memberVipDAO.BatchGet(ctx, f.ids)
	util.PanicIf(err)
	for _, member := range f.memberMap {
		if v, ok := memberVipMap[member.ID]; ok {
			now := time.Now().Unix()
			if v.StartAt.Unix() < now && v.EndAt.Unix() > now {
				member.Vip = &Vip{
					IsVip:   true,
					StartAt: util2.Int64Ptr(v.StartAt.Unix()),
					EndAt:   util2.Int64Ptr(v.EndAt.Unix()),
				}
				continue
			}
		}
		member.Vip = &Vip{
			IsVip: false,
		}
	}
}

func (f *MemberRender) RenderDownloadCount(ctx context.Context) {
	countMap, err := f.memberDownloadNumberDAO.BatchGetMemberNormalCount(ctx, f.ids)
	util.PanicIf(err)
	for _, member := range f.memberMap {
		member.DownloadCount = countMap[member.ID]
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
		inviteCodeString = util2.MustGenerateCode(ctx, 6)
		util.PanicIf(f.memberInviteCodeDAO.Insert(ctx, &models.MemberInviteCode{
			MemberID: f.loginID,
			Code:     inviteCodeString,
		}))
	}

	for _, member := range f.memberMap {
		if member.ID == f.loginID {
			member.InviteURL = util.StringPtr(fmt.Sprintf("https://www.dumpapp.com/register?invite_code=%s", inviteCodeString))
		}
	}
}

func (f *MemberRender) RenderDevices(ctx context.Context) {
	_, ok := f.memberMap[f.loginID]
	if !ok {
		return
	}

	memberDeviceMap, err := f.memberDeviceDAO.BatchGetByMemberIDs(ctx, []int64{f.loginID})
	util.PanicIf(err)

	devices, ok := memberDeviceMap[f.loginID]
	if !ok {
		return
	}
	result := make([]*Device, 0)
	for _, device := range devices {
		result = append(result, &Device{UDID: device.Udid})
	}
	f.memberMap[f.loginID].Devices = result
}

func (f *MemberRender) RenderAdmin(ctx context.Context) {
	adminMap := NewAdminRender(f.ids, f.loginID).RenderMap(ctx)
	for _, member := range f.memberMap {
		member.Admin = adminMap[member.ID]
	}
}
