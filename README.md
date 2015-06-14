go-httpserve
============

go-httpserve is a lightweight go library to
- run one single HTTP server on *a single port* which upgrades to HTTPS when a client tries to connect via HTTP.
- simplify basic HTTP authentication.

Pull requests, bug reporting and feature request (via github issues) are always welcome. :)

## Installation
```
go get github.com/codedust/go-httpserve
```

How to run
----------
If you are unfamiliar with Go, please start by reading
[How to Write Go Code](http://golang.org/doc/code.html).

The best place to get started are the examples in [examples/](examples/).
Simple run the examples with `go run` and visit
[http://localhost:8080/](http://localhost:8080/).

Feel free to ask for help in the issue tracker. ;)


## License
go-httpserve is licensed under the [Mozilla Public License](LICENSE).
