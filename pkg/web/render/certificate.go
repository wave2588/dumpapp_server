package render

import (
	"context"

	"dumpapp_server/pkg/common/enum"
	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/controller"
	impl3 "dumpapp_server/pkg/controller/impl"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	util2 "dumpapp_server/pkg/util"
)

type Certificate struct {
	Meta *models.CertificateV2 `json:"-"`

	ID        int64 `json:"id,string"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`

	P12IsActive bool `json:"p12_is_active" render:"method=RenderP12IsActive"`
}

type CertificateRender struct {
	ids           []int64
	loginID       int64
	includeFields []string

	certificateMap map[int64]*Certificate

	certificateDAO   dao.CertificateV2DAO
	certificateV1Ctl controller.CertificateController
	certificateV2Ctl controller.CertificateController
}

type CertificateOption func(*CertificateRender)

func CertificateIncludes(fields []string) CertificateOption {
	return func(render *CertificateRender) {
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

var CertificateDefaultRenderFields = []CertificateOption{
	CertificateIncludes([]string{
		"P12IsActive",
	}),
}

func NewCertificateRender(ids []int64, loginID int64, opts ...CertificateOption) *CertificateRender {
	f := &CertificateRender{
		ids:     ids,
		loginID: loginID,

		certificateDAO:   impl.DefaultCertificateV2DAO,
		certificateV1Ctl: impl3.DefaultCertificateV1Controller,
		certificateV2Ctl: impl3.DefaultCertificateV2Controller,
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func (f *CertificateRender) RenderSlice(ctx context.Context) []*Certificate {
	tMap := f.RenderMap(ctx)
	tSlice := make([]*Certificate, len(f.ids))
	for i, id := range f.ids {
		tSlice[i] = tMap[id]
	}
	return tSlice
}

func (f *CertificateRender) RenderMap(ctx context.Context) map[int64]*Certificate {
	if len(f.ids) == 0 {
		return f.certificateMap
	}

	f.fetch(ctx)

	err := autoRender(ctx, f, Certificate{}, f.includeFields)
	if err != nil {
		panic(err)
	}

	return f.certificateMap
}

func (f *CertificateRender) fetch(ctx context.Context) {
	cerMap, err := f.certificateDAO.BatchGet(ctx, f.ids)
	util.PanicIf(err)

	result := make(map[int64]*Certificate)
	for _, id := range f.ids {
		meta, ok := cerMap[id]
		if !ok {
			continue
		}
		result[meta.ID] = &Certificate{
			Meta:      meta,
			ID:        meta.ID,
			CreatedAt: meta.CreatedAt.Unix(),
			UpdatedAt: meta.UpdatedAt.Unix(),
		}
	}
	f.certificateMap = result
}

func (f *CertificateRender) RenderP12IsActive(ctx context.Context) {
	isActiveMap := make(map[int64]bool)
	batch := util2.NewBatch(ctx)
	for _, certificate := range f.certificateMap {
		batch.Append(func(cer *Certificate) util2.FutureFunc {
			return func() error {
				switch cer.Meta.Source {
				case enum.CertificateSourceV1:
					response, err := f.certificateV1Ctl.CheckCerIsActive(ctx, cer.ID)
					if err != nil {
						return err
					}
					isActiveMap[cer.ID] = response
				case enum.CertificateSourceV2:
					response, err := f.certificateV2Ctl.CheckCerIsActive(ctx, cer.ID)
					if err != nil {
						return err
					}
					isActiveMap[cer.ID] = response
				}
				return nil
			}
		}(certificate))
	}
	batch.Wait()

	for _, certificate := range f.certificateMap {
		certificate.P12IsActive = isActiveMap[certificate.ID]
	}
}
