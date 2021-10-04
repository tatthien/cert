# cert
Retrieve SSL certificate information from provided hostname.

## Why

I just simply want to retrieve a website's SSL certificate information in my terminal.

## Install

```
$ go get github.com/tatthien/cert
```

For non-Go users

```
$ curl -sf https://gobinaries.com/tatthien/cert | sh
```

## Usage

```shell
# Simply run cert --host=<website_without_http>
$ cert --host=12bit.vn

# Example output
> 12bit.vn
This certificate is valid
Issued by: GTS CA 1D4
 ├── Org: Google Trust Services LLC
 └── Country: US
Expires: Saturday, 25-Dec-21 02:45:49 UTC (2 months from now)

> github.com
This certificate is valid
Issued by: DigiCert High Assurance TLS Hybrid ECC SHA256 2020 CA1
 ├── Org: DigiCert, Inc.
 └── Country: US
Expires: Wednesday, 30-Mar-22 23:59:59 UTC (5 months from now)

> expired.badssl.com
x509: certificate has expired or is not yet valid: current time 2021-10-04T21:42:36+07:00 is after 2015-04-12T23:59:59Z
```

## License

MIT © Thien Nguyen
