# Gin Web Framework

<img align="right" width="159px" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">

[![Build Status](https://github.com/gin-gonic/gin/workflows/Run%20Tests/badge.svg?branch=master)](https://github.com/gin-gonic/gin/actions?query=branch%3Amaster)
[![codecov](https://codecov.io/gh/gin-gonic/gin/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-gonic/gin)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-gonic/gin)](https://goreportcard.com/report/github.com/gin-gonic/gin)
[![Go Reference](https://pkg.go.dev/badge/github.com/gin-gonic/gin?status.svg)](https://pkg.go.dev/github.com/gin-gonic/gin?tab=doc)
[![Sourcegraph](https://sourcegraph.com/github.com/gin-gonic/gin/-/badge.svg)](https://sourcegraph.com/github.com/gin-gonic/gin?badge)
[![Open Source Helpers](https://www.codetriage.com/gin-gonic/gin/badges/users.svg)](https://www.codetriage.com/gin-gonic/gin)
[![Release](https://img.shields.io/github/release/gin-gonic/gin.svg?style=flat-square)](https://github.com/gin-gonic/gin/releases)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/gin-gonic/gin)](https://www.tickgit.com/browse?repo=github.com/gin-gonic/gin)

Gin is a web framework written in [Go](https://go.dev/). It features a martini-like API with performance that is up to 40 times faster thanks to [httprouter](https://github.com/julienschmidt/httprouter).
If you need performance and good productivity, you will love Gin.

**Gin's key features are:**

- Zero allocation router
- Speed
- Middleware support
- Crash-free
- JSON validation
- Route grouping
- Error management
- Built-in rendering
- Extensible

## Getting started

### Prerequisites

Gin requires [Go](https://go.dev/) version [1.22](https://go.dev/doc/devel/release#go1.22.0) or above.

### Getting Gin

With [Go's module support](https://go.dev/wiki/Modules#how-to-use-modules), `go [build|run|test]` automatically fetches the necessary dependencies when you add the import in your code:

```sh
import "github.com/gin-gonic/gin"
```

Alternatively, use `go get`:

```sh
go get -u github.com/gin-gonic/gin
```

### Running Gin

A basic example:

```go
package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

To run the code, use the `go run` command, like:

```sh
$ go run example.go
```

Then visit [`0.0.0.0:8080/ping`](http://0.0.0.0:8080/ping) in your browser to see the response!

### See more examples

#### Quick Start

Learn and practice with the [Gin Quick Start](docs/doc.md), which includes API examples and builds tag.

#### Examples

A number of ready-to-run examples demonstrating various use cases of Gin are available in the [Gin examples](https://github.com/gin-gonic/examples) repository.

## Documentation

See the [API documentation on go.dev](https://pkg.go.dev/github.com/gin-gonic/gin).

The documentation is also available on [gin-gonic.com](https://gin-gonic.com) in several languages:

- [English](https://gin-gonic.com/docs/)
- [简体中文](https://gin-gonic.com/zh-cn/docs/)
- [繁體中文](https://gin-gonic.com/zh-tw/docs/)
- [日本語](https://gin-gonic.com/ja/docs/)
- [Español](https://gin-gonic.com/es/docs/)
- [한국어](https://gin-gonic.com/ko-kr/docs/)
- [Turkish](https://gin-gonic.com/tr/docs/)
- [Persian](https://gin-gonic.com/fa/docs/)
- [Português](https://gin-gonic.com/pt/docs/)

### Articles

- [Tutorial: Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)