package impl

import (
	"context"
	"fmt"

	"dumpapp_server/pkg/common/constant"
	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/controller"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"github.com/pkg/errors"
)

type CertificatePriceController struct {
	accountDAO dao.AccountDAO
}

var DefaultCertificatePriceController controller.CertificatePriceController

func init() {
	DefaultCertificatePriceController = NewCertificatePriceController()
}

func NewCertificatePriceController() *CertificatePriceController {
	return &CertificatePriceController{
		accountDAO: impl.DefaultAccountDAO,
	}
}

type price struct {
	ID          int64                                    // id
	Price       func(accountRole enum.AccountRole) int64 // 价格
	Title       string                                   // 标题
	Description string                                   // 描述
}

var prices = []*price{
	{
		ID: constant.CertificateIDL1,
		Price: func(accountRole enum.AccountRole) int64 {
			if accountRole == enum.AccountRoleAgent {
				return constant.CertificatePriceAgentL1
			}
			return constant.CertificatePriceL1
		},
		Title:       constant.CertificateTitleL1,
		Description: constant.CertificateDescriptionL1,
	},
	{
		ID: constant.CertificateIDL2,
		Price: func(accountRole enum.AccountRole) int64 {
			if accountRole == enum.AccountRoleAgent {
				return constant.CertificatePriceAgentL2
			}
			return constant.CertificatePriceL2
		},
		Title:       constant.CertificateTitleL2,
		Description: constant.CertificateDescriptionL2,
	},
	{
		ID: constant.CertificateIDL3,
		Price: func(accountRole enum.AccountRole) int64 {
			if accountRole == enum.AccountRoleAgent {
				return constant.CertificatePriceAgentL3
			}
			return constant.CertificatePriceL3
		},
		Title:       constant.CertificateTitleL3,
		Description: constant.CertificateDescriptionL3,
	},
}

func (c *CertificatePriceController) BatchGetPrices(ctx context.Context, memberIDs []int64) (map[int64][]*controller.CertificatePriceInfo, error) {
	accountMap, err := c.accountDAO.BatchGet(ctx, memberIDs)
	if err != nil {
		return nil, err
	}
	result := make(map[int64][]*controller.CertificatePriceInfo)
	for _, memberID := range memberIDs {
		accountRole := enum.AccountRoleNone
		if account, ok := accountMap[memberID]; ok {
			accountRole = account.Role
		}
		resultPrice := make([]*controller.CertificatePriceInfo, 0)
		for _, p := range prices {
			resultPrice = append(resultPrice, &controller.CertificatePriceInfo{
				ID:          p.ID,
				Price:       p.Price(accountRole),
				Title:       p.Title,
				Description: p.Description,
			})
		}
		result[memberID] = resultPrice
	}
	return result, nil
}

func (c *CertificatePriceController) GetPrices(ctx context.Context, memberID int64) ([]*controller.CertificatePriceInfo, error) {
	priceMap, err := c.BatchGetPrices(ctx, []int64{memberID})
	if err != nil {
		return nil, err
	}
	return priceMap[memberID], nil
}

func (c *CertificatePriceController) GetPriceByID(ctx context.Context, memberID, priceID int64) (*controller.CertificatePriceInfo, error) {
	prices, err := c.GetPrices(ctx, memberID)
	if err != nil {
		return nil, err
	}
	for _, info := range prices {
		if info.ID == priceID {
			return info, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("获取价格失败 memberID: %d  priceID: %d", memberID, priceID)) /// 理论上不会走到这
}
