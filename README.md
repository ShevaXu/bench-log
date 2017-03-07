# Benchmarking pkg/log

Attempt to improve logging performance by avoid unnecessary mutex operations (d25504f).

```
$ go test -bench="."
```

**Result:** 

*(Note that each run might be slightly different.)*

```
BenchmarkMutex-8                      	100000000	        17.2 ns/op
BenchmarkLog-8                        	 3000000	       432 ns/op
BenchmarkNewLog-8                     	 3000000	       425 ns/op
BenchmarkLogWithFileLine-8            	 1000000	      1047 ns/op
BenchmarkNewLogWithFileLine-8         	 1000000	      1085 ns/op
BenchmarkLogWithFileLineMulti-8       	  200000	      5430 ns/op
BenchmarkNewLogWithFileLineMulti-8    	  300000	      4327 ns/op
BenchmarkLogWithFileLineMulti2-8      	  100000	     11478 ns/op
BenchmarkNewLogWithFileLineMulti2-8   	  200000	      8457 ns/op
```