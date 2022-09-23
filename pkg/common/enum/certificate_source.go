//go:generate enumer -type=CertificateSource -json -sql -transform=snake -trimprefix=CertificateSource
// go get github.com/dmarkham/enumer
package enum

type CertificateSource int

const (
	CertificateSourceV1 CertificateSource = iota + 1
	CertificateSourceV2
	CertificateSourceV3
)
