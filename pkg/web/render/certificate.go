package render

import (
	"context"

	"dumpapp_server/pkg/common/util"
	"dumpapp_server/pkg/dao"
	"dumpapp_server/pkg/dao/impl"
	"dumpapp_server/pkg/dao/models"
	"dumpapp_server/pkg/http"
	impl2 "dumpapp_server/pkg/http/impl"
	util2 "dumpapp_server/pkg/util"
)

type Certificate struct {
	meta *models.Certificate

	ID        int64 `json:"id,string"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`

	P12IsActive bool `json:"p12_is_active" render:"method=RenderP12IsActive"`
	IsValidate  bool `json:"is_validate" render:"method=RenderIsValidate"`
}

type CertificateRender struct {
	ids           []int64
	loginID       int64
	includeFields []string

	CertificateMap map[int64]*Certificate

	certificateDAO   dao.CertificateDAO
	certificateServe http.CertificateServer
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
		"IsValidate",
	}),
}

func NewCertificateRender(ids []int64, loginID int64, opts ...CertificateOption) *CertificateRender {
	f := &CertificateRender{
		ids:     ids,
		loginID: loginID,

		certificateDAO:   impl.DefaultCertificateDAO,
		certificateServe: impl2.DefaultCertificateServer,
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
		return f.CertificateMap
	}

	f.fetch(ctx)

	err := autoRender(ctx, f, Certificate{}, f.includeFields)
	if err != nil {
		panic(err)
	}

	return f.CertificateMap
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
			meta:      meta,
			ID:        meta.ID,
			CreatedAt: meta.CreatedAt.Unix(),
			UpdatedAt: meta.UpdatedAt.Unix(),
		}
	}
	f.CertificateMap = result
}

func (f *CertificateRender) RenderP12IsActive(ctx context.Context) {
	batch := util2.NewBatch(ctx)
	isActiveMap := make(map[int64]bool)
	for _, certificate := range f.CertificateMap {
		batch.Append(func(cer *Certificate) util2.FutureFunc {
			return func() error {
				response, err := f.certificateServe.CheckP12File(ctx, cer.meta.P12FileDate, "1")
				if err != nil {
					return err
				}
				isActiveMap[cer.ID] = response.Data
				return nil
			}
		}(certificate))
	}
	batch.Wait()

	for _, certificate := range f.CertificateMap {
		certificate.P12IsActive = isActiveMap[certificate.ID]
	}
}

func (f *CertificateRender) RenderIsValidate(ctx context.Context) {
	batch := util2.NewBatch(ctx)
	isValidateMap := make(map[int64]bool)
	for _, certificate := range f.CertificateMap {
		batch.Append(func(cer *Certificate) util2.FutureFunc {
			return func() error {
				if cer.meta.UdidBatchNo == "" {
					return nil
				}
				response, err := f.certificateServe.CheckCerByUDIDBatchNo(ctx, cer.meta.UdidBatchNo)
				if err != nil {
					return err
				}
				isValidateMap[cer.ID] = response.Data
				return nil
			}
		}(certificate))
	}
	batch.Wait()

	for _, certificate := range f.CertificateMap {
		certificate.IsValidate = isValidateMap[certificate.ID]
	}
}
