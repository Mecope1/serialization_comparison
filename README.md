# A comparison and quick discussion about serialization choices.

Currently a work in process. I plan on adding more methods of serializing data and adding a link to an article that will talk about this topic more in-depth.

To run the tests, type in `go test -bench=.` while in the main project directory. 

Here are the results when run on my ubuntu machine.

```shell script
goos: linux
goarch: amd64
pkg: serialization_comparison
BenchmarkMsgpackEnc-4            1839358               685 ns/op
BenchmarkMsgpackDec-4            1231225               969 ns/op
BenchmarkJsonEnc-4               1966341               604 ns/op
BenchmarkJsonDec-4                722858              1724 ns/op
BenchmarkMsgpackEnc_Long-4        770421              1713 ns/op
BenchmarkMsgpackDec_Long-4        381255              3028 ns/op
BenchmarkJsonEnc_Long-4           165088              7829 ns/op
BenchmarkJsonDec_Long-4            40827             29517 ns/op
BenchmarkTlvEnc_Chat-4           1642754               724 ns/op
BenchmarkTlvDec_Chat-4           1000000              1161 ns/op
BenchmarkMsgpackEnc_Chat-4        802418              1536 ns/op
BenchmarkMsgpackDec_Chat-4        769466              1625 ns/op
BenchmarkJsonEnc_Chat-4           913416              1107 ns/op
BenchmarkJsonDec_Chat-4           443504              2422 ns/op
PASS
ok      serialization_comparison        20.326s

```

The ones that are relevant to my chat application, nicknamed [tin-can-communicator](https://github.com/Mecope1/tin-can-communicator),
are the last 12 that have _Chat-4 at the end of their names. Based on the results here, there are a few takeaways: 


* The first is that MessagePack is much faster than Json, especially as the size of the data grows.

* The second is that decoding took much more time than encoding, which I found interesting.

* The final bit is that a type-length-value (TLV) style of protocol does allow someone to squeeze out some 
performance gains over Json or MessagePack. It would require you to maintain the protocol yourself, and to implement 
it for any new languages. Conversely, Json is essentially omnipresent and you could probably find some form of MessagePack
in most languages that are popular. It may be worth it to create your own protocol when you truly need the extra performance
and smaller payload size, but for most applications it would lead to negligible gains and an increase in tech-debt.