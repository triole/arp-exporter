# ARP Exporter

## Help

```go mdox-exec="r -h"
Usage: arp-exporter [flags]

arp table exporter, supports prometheus metrics or json format

Flags:
  -h, --help                      Show context-sensitive help.
  -i, --info=STRING               look up and display mac address vendor
                                  information
  -l, --list                      list all vendors in db
  -c, --config=STRING             configuration file
  -e, --enable-vendors            enable displaying mac vendors
  -s, --server                    run web server
  -b, --bind=":9100"              bind to
      --log-file="/dev/stdout"    log file
      --log-level="info"          log level
      --log-no-colors             disable output colours, print plain text
      --log-json                  enable json log, instead of text one
  -n, --dry-run                   dry run, just print operations that would run
  -V, --version-flag              display version
```

## Vendors

Vendor database taken from [maclookup.app](https://maclookup.app/downloads/json-database).
