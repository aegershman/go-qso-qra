# go-qso-qra

FCC amateur radio callsign search from your cli.

- Caches results in `~/.ham_cache`.
- Provides a Google Maps link to the callsign's address.
- Uses APIs from `hamdb` (default) or `callook` using `-api {hamdb, callook}` flag.

## usage

```sh
qso wb6acu
```

example output

```txt
Callsign: WB6ACU
Class: E
Expires: 08/27/2031
Status: A
Grid: DM04sc
Latitude: 34.0986995
Longitude: -118.4198427
First Name: JOSEPH
Middle Initial: F
Last Name: WALSH
Suffix:
Address 1: 1501 Summitridge Dr
Address 2: Beverly Hills
State: CA
Zip: 90210
Country: United States
Google Maps Link: https://www.google.com/maps?q=34.0986995,-118.4198427
```

## install

```sh
brew install aegershman/tap/qso
```
