package cert

import (
	"crypto/tls"
	"crypto/x509/pkix"
	"fmt"
	"net"
	"time"
)

var (
	dialer = &net.Dialer{Timeout: 5 * time.Second}
)

type CertChecker struct {
	Hostname string
}

type Cert struct {
	Issuer  pkix.Name
	Expires time.Time
}

func New(hostname string) *CertChecker {
	return &CertChecker{
		Hostname: hostname,
	}
}

// GetCertificate get the certificate information from the provided hostname.
func (ck *CertChecker) GetCertificate() (Cert, error) {
	addr := fmt.Sprintf("%s:%d", ck.Hostname, 443)

	conn, err := tls.DialWithDialer(dialer, "tcp", addr, nil)

	if err != nil {
		return Cert{}, err
	}

	defer conn.Close()

	err = conn.VerifyHostname(ck.Hostname)

	if err != nil {
		return Cert{}, err
	}

	c := conn.ConnectionState().PeerCertificates[0]

	return Cert{
		Issuer:  c.Issuer,
		Expires: c.NotAfter,
	}, nil
}
