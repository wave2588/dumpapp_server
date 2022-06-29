package install_app

import "context"

type ALiPayInstallAppController interface {
	GetPayURLByInstallApp(ctx context.Context, number int64, contactWay string) (int64, string, error)
	AliPayCallbackOrder(ctx context.Context, orderID int64) error
}
