package cmd

import (
	"fmt"
	"log"

	"github.com/dustin/go-humanize"
	flag "github.com/spf13/pflag"
	"github.com/tatthien/cert/pkg/cert"
)

var host string

func init() {
	flag.StringVar(&host, "host", "", "hostname to check the certificate")
}

// Execute shows the certificate information from the provided hostname.
func Execute() {
	flag.Parse()

	cert, err := cert.GetCertificate(host)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(host)
	fmt.Printf("Issued by: %s\n", cert.Issuer.CommonName)
	fmt.Printf("Expires: %v (%s)\n", cert.Expires, humanize.Time(cert.Expires))
}
