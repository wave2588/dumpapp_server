package controller

import "context"

type AdminDumpOrderController interface {
	Upsert(ctx context.Context, demanderID, ipaID int64, ipaName, ipaVersion, ipaBundleID, ipaAppStoreLink string, isOld bool) error
	Progressed(ctx context.Context, operatorID, ipaID int64, ipaVersion string) error
	Delete(ctx context.Context, operatorID, ipaID int64, ipaVersion string) error
}
