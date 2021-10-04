package cmd

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/tatthien/cert/pkg/cert"
)

// Execute shows the certificate information from the provided hostname.
func Execute() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("Usage: cert hostname_1 hostname_2 hostname_3 ...")
		flag.PrintDefaults()
		os.Exit(1)
	}

	hosts := flag.Args()

	for _, host := range hosts {
		fmt.Printf("> %s\n", host)

		ck := cert.New(host)

		cert, err := ck.GetCertificate()

		if err != nil {
			fmt.Printf("%s\n", err.Error())
		} else {
			fmt.Println("This certificate is valid")
			fmt.Printf("Issued by: %s\n", cert.Issuer.CommonName)
			fmt.Printf(" ├── Org: %s\n", cert.Issuer.Organization[0])
			fmt.Printf(" └── Country: %s\n", cert.Issuer.Country[0])
			fmt.Printf("Expires: %v (%s)\n", cert.Expires.Format(time.RFC850), humanize.Time(cert.Expires))
			fmt.Println()
		}
	}
}
