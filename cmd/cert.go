package cmd

import (
	"fmt"
	"log"
	"time"

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

	expires, err := getLocalTime(cert.Expires)

	if err != nil {
		log.Fatal("cannot load time location")
	}

	fmt.Printf("Expires: %v (%s)\n", expires.Format("Monday, 02 January 2006 at 15:04:05"), humanize.Time(cert.Expires))
}

func getLocalTime(t time.Time) (time.Time, error) {
	loc, err := time.LoadLocation("Local")

	if err != nil {
		return t, err
	}

	return t.In(loc), nil
}
