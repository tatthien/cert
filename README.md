# cert
Retrieve SSL certificate information from provided hostname.

## Why

I just simply want to retrieve a website's SSL certificate information in my terminal.

## Usage

```shell
$ go get github.com/tatthien/cert

# Simply run cert --host=<website_without_http>
$ cert --host=12bit.vn

# Example output
12bit.vn
Issued by: GTS CA 1D4
Expires: 2021-12-25 02:45:49 +0000 UTC (2 months from now
```

## License

MIT © Thien Nguyen
