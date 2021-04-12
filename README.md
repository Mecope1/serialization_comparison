# A comparison and quick discussion about serialization choices.

Currently a work in process. I plan on adding more methods of serializing data and adding a link to an article that will talk about this topic more in-depth.

To run the tests, type in `go test -bench=.` while in the main project directory. 

Here are the results when run on my ubuntu machine.

```shell script
goos: linux
goarch: amd64
pkg: serialization_comparison
BenchmarkMsgpackEnc-4       	 2100422	       559 ns/op
BenchmarkMsgpackDec-4       	 1750468	       698 ns/op
BenchmarkJsonEnc-4          	 2029009	       592 ns/op
BenchmarkJsonDec-4          	  667212	      1620 ns/op
BenchmarkMsgpackEncLong-4   	  768330	      1908 ns/op
BenchmarkMsgpackDecLong-4   	  686906	      1848 ns/op
BenchmarkJsonEncLong-4      	  164044	      7451 ns/op
BenchmarkJsonDecLong-4      	   39748	     29106 ns/op
PASS
```
