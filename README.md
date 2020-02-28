# logrus logger adapter
[![Build Status](https://travis-ci.org/clevergo/log-logrus.svg?branch=master)](https://travis-ci.org/clevergo/log-logrus)
[![Coverage Status](https://coveralls.io/repos/github/clevergo/log-logrus/badge.svg?branch=master)](https://coveralls.io/github/clevergo/log-logrus?branch=master)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/clevergo/log-logrus)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/log-logrus)](https://goreportcard.com/report/github.com/clevergo/log-logrus)
[![Release](https://img.shields.io/github/release/clevergo/log-logrus.svg?style=flat-square)](https://github.com/clevergo/log-logrus/releases)

This package is an adapter of [universal logger interface](https://github.com/clevergo/log) for [logrus](https://github.com/sirupsen/logrus).

## Usage

```go
import (
    "github.com/clevergo/log"
    logrusadapter "github.com/clevergo/log-logrus"
    "github.com/sirupsen/logrus"
)

var logger log.Logger

func main() {
    logger = logrusadapter.New(logrus.New())
    logger.Debug("debug msg")
    // ...
}
```
