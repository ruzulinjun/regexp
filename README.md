# regexp

Serve embedded files from [golang/go/src/regexp](https://github.com/golang/go/src/regexp).

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

     re := regexp.MustCompile(`[^可]+爱`)
     begin := time.Now()
     for i := 0; i <= 1000000; i++ {
         re.FindString(`这篇文档中，我使用这句话用来测试正则匹配的效率，查看性能提升情况。`)
     }
     fmt.Println(time.Since(begin))

golang/go/src/regexp uses 3.941630471s.
ruzulinjun/regexp uses 175.001135ms.
Performance improves 95.6%.

### the use of this package is same as golang/go/src/regexp
