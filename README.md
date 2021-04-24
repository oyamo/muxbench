# BenchMark

Benchmarking gorrilla mux and golang http against fiber

## Simple Hello World Response
```
oyamo@ultrabook:~$ wrk -t12 -c400 -d30s http://127.0.0.1:3000/
Running 30s test @ http://127.0.0.1:3000/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.27ms    4.60ms 126.35ms   89.61%
    Req/Sec     9.30k     2.65k   19.92k    62.05%
  3331129 requests in 30.09s, 238.26MB read
Requests/sec: 110689.49
Transfer/sec:      7.92MB

```