Dirty SSL Bench
------

Compare Nginx SSL Termination Overhead (within a reverse proxy) vs Go (in process).

### Setup

Run `./generate-keys.sh localhost` to create a a local ssl cert. Copy the keyfiles into
`/etc/nginx/` along with the `nginx.conf` here.

The Go binary is a simple Hello World HTTP server that needs to be recompiled to serve HTTPS.

# Results

All tests here (not very scientific, for sure) were run on a quad core i7, if that is relevant.

### Nginx Results

```
fortytw2@fortytw2 ~ % echo "GET https://localhost:8081/" | vegeta attack -duration=30s -insecure | tee results.bin | vegeta report


Requests      [total, rate]            1500, 50.03
Duration      [total, attack, wait]    29.98051442s, 29.979999935s, 514.485µs
Latencies     [mean, 50, 95, 99, max]  682.338µs, 686.243µs, 751.002µs, 1.677588ms, 18.316751ms
Bytes In      [total, mean]            18000, 12.00
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:1500
Error Set:
```

### Go Results

```
fortytw2@fortytw2 ~ % echo "GET https://localhost:8080/" | vegeta attack -duration=30s -insecure | tee results.bin | vegeta report


Requests      [total, rate]            1500, 50.03
Duration      [total, attack, wait]    29.980323182s, 29.979999924s, 323.258µs
Latencies     [mean, 50, 95, 99, max]  491.133µs, 434.315µs, 735.471µs, 1.18063ms, 19.708664ms
Bytes In      [total, mean]            18000, 12.00
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:1500
Error Set:
```
