package constant

const (
	CertificatePriceL1 int64 = 35
	CertificatePriceL2 int64 = 78
	CertificatePriceL3 int64 = 97
)

type CertificatePriceInfo struct {
	ID          int64  `json:"id,string"`
	Price       int64  `json:"price"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func GetCertificatePrices() []*CertificatePriceInfo {
	return []*CertificatePriceInfo{
		{
			ID:          1,
			Price:       CertificatePriceL1,
			Title:       "普通版",
			Description: "理论 1 年。",
		},
		{
			ID:          2,
			Price:       CertificatePriceL2,
			Title:       "稳定版",
			Description: "理论 1 年，售后半年，掉了补 6 次。",
		},
		{
			ID:          3,
			Price:       CertificatePriceL3,
			Title:       "豪华版",
			Description: "理论 1 年，售后 1 年，掉了补 12 次。",
		},
	}
}
