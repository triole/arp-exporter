# ARP Exporter

## Help

```go mdox-exec="r -h"
Usage: arp-exporter [flags]

arp table exporter, supports prometheus metrics or json format

Flags:
  -h, --help                      Show context-sensitive help.
  -f, --arp-table-file=STRING     retrieve arp table from file, default is by
                                  command 'arp -an'
  -m, --mac-info=STRING           look up and display mac address vendor
                                  information
  -l, --list-vendors              list all vendors in db
  -n, --hostname-config=STRING    host name configuration file
  -e, --enable-vendors            enable displaying mac vendors
  -s, --server                    run web server
  -b, --bind=":9100"              where to bind the server to
      --log-file="/dev/stdout"    log file
      --log-level="info"          log level
      --log-no-colors             disable output colours, print plain text
      --log-json                  enable json log, instead of text one
  -V, --version-flag              display version
```

## Hostname Lookup

The tool supports hostname lookup if a hostnames config is provided using `-n`. The hostname config consists of an array that contains a list of device information. These can be the mac address, the device's name and the id of the network interface. If a mac address of the arp table matches a mac address of the hostname config entry, the information will be added to the final output data. This enables you to enrich the arp table or metrics result with metadata like for example the name of the device. A hostname config looks like the following...

```go mdox-exec="tail -n+2 examples/hostnames.yaml"
- name: my_pc
  itf: eth0
  mac: bc:fc:e7:yy:zz:zz
- name: my_tablet
  itf: wifi
  mac: 00:1a:11:aa:bb:cc
- name: another_pc
  itf: wifi
  mac: 78:d6:b2:jj:kk:ll
```

## Vendors

If enabled using `-e`, the arp-exporter also adds vendor information to the output. Vendor database taken from [maclookup.app](https://maclookup.app/downloads/json-database).
