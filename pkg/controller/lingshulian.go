package controller

import "context"

type LingshulianController interface {

	/// sign_ipa bucket
	PutMemberSignIpa(ctx context.Context, name string, data string) error
	GetMemberSignIpa(ctx context.Context, ipaToken string) (string, error)
	DeleteMemberSignIpa(ctx context.Context, tokenPath string) error
}
