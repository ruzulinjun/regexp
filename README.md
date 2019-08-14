# regexp

Serve embedded files from [golang/go/src/regexp](https://github.com/golang/go/tree/master/src/regexp).

[GoDoc](http://godoc.org/github.com/ruzulinjun/regexp)

### Installation

Install with

    $ go get github.com/ruzulinjun/regexp

### Provide possession quantifier

Golang package regexp dose not provide possession quantifier.

use

    func BenchmarkMain(b \*testing.B) {
    	re := regexp.MustCompile("^[^可爱]*爱爱|^[^不怕约]*约")
    	for i := 0; i < b.N; i++ {
    		re.FindString("这篇文档中，我使用这句话用来测试正则匹配的效率，查看性能提升情况。")
    	}
    }

use golang package regexp: BenchmarkMain-4   	  200000	     11675 ns/op

This package provides star and plus possession quantifier.

use 

    func BenchmarkMain(b \*testing.B) {
    	re := regexp.MustCompile("^[^可爱]*+爱爱|^[^不怕约]*+约")
    	for i := 0; i < b.N; i++ {
    		re.FindString("这篇文档中，我使用这句话用来测试正则匹配的效率，查看性能提升情况。")
    	}
    }

use ruzulinjun/regexp: BenchmarkMain-4   	200000	      9430 ns/op

Performance improves 19.2%.

### For scenes using Chinese, Optimize performance
English has only 26 letters, while Chinese has thousands daily characters.

Based on this scenes, this package optimizes performance.

use

    func BenchmarkMain(b \*testing.B) {
    	re := regexp.MustCompile("^[^可]*爱爱")
    	for i := 0; i < b.N; i++ {
    		re.FindString("这篇文档中，我使用这句话用来测试正则匹配的效率，查看性能提升情况。")
    	}
    }

use golang package regexp: BenchmarkMain-4   	  500000	      2806 ns/op

use ruzulinjun/regexp: BenchmarkMain-4   	10000000	       178 ns/op 

Performance improves 93.7%.

### the use of this package is the same as golang package regexp
