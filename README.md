# regexp

Serve embedded files from [golang/go/src/regexp](https://github.com/golang/go/tree/master/src/regexp).

[GoDoc](http://godoc.org/github.com/ruzulinjun/regexp)

### Installation

Install with

    $ go get github.com/ruzulinjun/regexp/...

### Provide possession quantifier

Golang package regexp dose not provide possession quantifier.

This package provides star and plus possession quantifier.

### For scenes using Chinese, Optimize performance
English has only 26 letters, while Chinese has thousands daily characters.

Based on this scenes, this package optimizes performance.

use

func BenchmarkMain(b \*testing.B) {
	re := regexp.MustCompile("[^可]+爱")
	for i := 0; i < b.N; i++ {
		re.FindString("这篇文档中，我使用这句话用来测试正则匹配的效率，查看性能提升情况。")
	}
}

golang package regexp uses: BenchmarkMain-4   	  500000	      3997 ns/op

ruzulinjun/regexp uses: BenchmarkMain-4   	10000000	       176 ns/op 

Performance improves 95.6%.

### the use of this package is the same as golang package regexp
