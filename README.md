

./ex_cache_bench_test.go -test.bench ".*" -test.benchtime 3s -test.v -test.benchmem
goos: darwin
goarch: amd64
BenchmarkFreeCacheSet-8           	 5205291	       806 ns/op	     331 B/op	       2 allocs/op
BenchmarkGoCacheSet-8             	 4207192	       926 ns/op	     358 B/op	       4 allocs/op
BenchmarkTTLCacheSet-8            	 1941025	      1855 ns/op	     421 B/op	       5 allocs/op
BenchmarkGCacheSet-8              	 2024690	      2179 ns/op	     572 B/op	       8 allocs/op
BenchmarkFreeCacheGet-8           	 7020132	       815 ns/op	      24 B/op	       1 allocs/op
BenchmarkGoCacheGet-8             	 7618544	       543 ns/op	      24 B/op	       1 allocs/op
BenchmarkTTLCacheGet-8            	 3301683	      1002 ns/op	      24 B/op	       2 allocs/op
BenchmarkGCacheGet-8              	 2465776	      1473 ns/op	     188 B/op	       4 allocs/op
BenchmarkFreeCacheSetParallel-8   	25578692	       428 ns/op	     329 B/op	       3 allocs/op
BenchmarkGoCacheSetParallel-8     	 3775754	       959 ns/op	     388 B/op	       5 allocs/op
BenchmarkTTLCacheSetParallel-8    	 1611339	      2247 ns/op	     380 B/op	       6 allocs/op
BenchmarkGCacheSetParallel-8      	 2138270	      2042 ns/op	     567 B/op	       9 allocs/op
BenchmarkFreeCacheGetParallel-8   	28344583	       225 ns/op	      24 B/op	       2 allocs/op
BenchmarkGoCacheGetParallel-8     	41591349	       230 ns/op	      24 B/op	       2 allocs/op
BenchmarkTTLCacheGetParallel-8    	 4231869	       719 ns/op	      24 B/op	       2 allocs/op
BenchmarkGCacheGetParallel-8      	 3421113	      1816 ns/op	      93 B/op	       3 allocs/op
PASS
ok  	command-line-arguments	305.937s
