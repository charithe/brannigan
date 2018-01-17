[![GoDoc][doc-img]][doc]


Brannigan
=========

Zap + capnslog (Captain's log) = Zapp Brannigan


A shim to redirect [CoreOS capnslog](https://godoc.org/github.com/coreos/pkg/capnslog) output to [Uber Zap](https://godoc.org/go.uber.org/zap). 
This is specially useful when [embedding etcd](https://godoc.org/github.com/coreos/etcd/embed) into your own application.

```
go get github.com/charithe/brannigan
```

Usage
-----

```go
brannigan.RedirectCapnslogToGlobalZapLogger()
// or use your own logger instance
brannigan.RedirectCapnslog(yourZapLogger)
```


[doc-img]: https://godoc.org/github.com/charithe/brannigan?status.svg
[doc]: https://godoc.org/github.com/charithe/brannigan
