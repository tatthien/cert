package cert

import (
	"crypto/tls"
	"crypto/x509/pkix"
	"fmt"
	"time"
)

type Cert struct {
	Issuer  pkix.Name
	Expires time.Time
}

// GetCertificate get the certificate information from the provided hostname.
func GetCertificate(hostname string) (Cert, error) {
	addr := fmt.Sprintf("%s:%d", hostname, 443)

	conn, err := tls.Dial("tcp", addr, &tls.Config{
		InsecureSkipVerify: true,
	})

	if err != nil {
		return Cert{}, err
	}

	defer conn.Close()

	cert := conn.ConnectionState().PeerCertificates[0]
	return Cert{
		Issuer:  cert.Issuer,
		Expires: cert.NotAfter,
	}, nil
}
