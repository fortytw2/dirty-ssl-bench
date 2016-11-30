nginx

fortytw2@trainman ~ % echo "GET https://localhost:8081/" | vegeta attack -duration=10s -insecure | tee results.bin | vegeta report
Requests      [total, rate]            500, 50.10
Duration      [total, attack, wait]    9.980737837s, 9.979999933s, 737.904µs
Latencies     [mean, 50, 95, 99, max]  700.119µs, 702.379µs, 748.835µs, 825.008µs, 18.068072ms
Bytes In      [total, mean]            6000, 12.00
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:500
Error Set:

fortytw2@trainman ~ % echo "GET https://localhost:8081/" | vegeta attack -duration=30s -insecure | tee results.bin | vegeta report


Requests      [total, rate]            1500, 50.03
Duration      [total, attack, wait]    29.98051442s, 29.979999935s, 514.485µs
Latencies     [mean, 50, 95, 99, max]  682.338µs, 686.243µs, 751.002µs, 1.677588ms, 18.316751ms
Bytes In      [total, mean]            18000, 12.00
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:1500
Error Set:



go


fortytw2@trainman ~ % echo "GET https://localhost:8080/" | vegeta attack -duration=10s -insecure | tee results.bin | vegeta report


Requests      [total, rate]            500, 50.10
Duration      [total, attack, wait]    9.980341349s, 9.979999949s, 341.4µs
Latencies     [mean, 50, 95, 99, max]  523.136µs, 380.147µs, 601.988µs, 1.007531ms, 34.351731ms
Bytes In      [total, mean]            6000, 12.00
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:500
Error Set:
fortytw2@trainman ~ % echo "GET https://localhost:8080/" | vegeta attack -duration=10s -insecure | tee results.bin | vegeta report


Requests      [total, rate]            500, 50.10
Duration      [total, attack, wait]    9.980341349s, 9.979999949s, 341.4µs
Latencies     [mean, 50, 95, 99, max]  523.136µs, 380.147µs, 601.988µs, 1.007531ms, 34.351731ms
Bytes In      [total, mean]            6000, 12.00
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:500
Error Set:


fortytw2@trainman ~ % echo "GET https://localhost:8080/" | vegeta attack -duration=30s -insecure | tee results.bin | vegeta report


Requests      [total, rate]            1500, 50.03
Duration      [total, attack, wait]    29.980323182s, 29.979999924s, 323.258µs
Latencies     [mean, 50, 95, 99, max]  491.133µs, 434.315µs, 735.471µs, 1.18063ms, 19.708664ms
Bytes In      [total, mean]            18000, 12.00
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:1500
Error Set:
