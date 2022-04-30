# log

![Go](https://github.com/go-ecosystem/log/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-ecosystem/log)](https://goreportcard.com/report/github.com/go-ecosystem/log)
[![Release](https://img.shields.io/github/release/go-ecosystem/log.svg)](https://github.com/go-ecosystem/log/releases)

> Golang log based on [zap](https://github.com/uber-go/zap).

## Installation

```shell
go get -u github.com/go-ecosystem/log
```

## Quick Start

```go
package main

import (
	"errors"

	"github.com/go-ecosystem/log"
)

const (
	simpleLog        = "log"
	errLog           = "error log"
	logWithFields    = "log with fields"
	errLogWithFields = "error log with fields"
	errLogWithMsg    = "error log with message"
)

func main() {
	// Development
	log.SetUp(false)
	defer log.Sync()

	log.Info(simpleLog)
	log.Info(logWithFields, log.Any("field", 1))

	log.Error(errLog)
	log.Error(errLogWithFields, log.Any("field", "field value"))
	log.ErrorE(errLogWithMsg, errors.New("test error"))

	log.Print("\n------\n")

	// Production
	log.SetUp(true, log.Any("serverName", "go-ecosystem"))

	log.Info(simpleLog)
	log.Info(logWithFields, log.Any("field", 1))

	log.Error(errLog)
	log.Error(errLogWithFields, log.Any("field", "field value"))
	log.ErrorE(errLogWithMsg, errors.New("test error"))

	log.Panic("panic")
	log.PanicE("panic with error", errors.New("test error"))
}
```

```shell
2020-11-10T22:28:47.483+0800    INFO    example/main.go:22      log
2020-11-10T22:28:47.483+0800    INFO    example/main.go:23      log with fields {"field": 1}
2020-11-10T22:28:47.483+0800    ERROR   example/main.go:25      error log
main.main
        /Users/catchzeng/Documents/Code/go-ecosystem/log/example/main.go:25
runtime.main
        /usr/local/go/src/runtime/proc.go:203
2020-11-10T22:28:47.483+0800    ERROR   example/main.go:26      error log with fields   {"field": "field value"}
main.main
        /Users/catchzeng/Documents/Code/go-ecosystem/log/example/main.go:26
runtime.main
        /usr/local/go/src/runtime/proc.go:203
2020-11-10T22:28:47.483+0800    ERROR   example/main.go:27      error log with message  {"error": "test error"}
main.main
        /Users/catchzeng/Documents/Code/go-ecosystem/log/example/main.go:27
runtime.main
        /usr/local/go/src/runtime/proc.go:203
2020-11-10T22:28:47.483+0800    DEBUG   example/main.go:29
------

{"level":"info","ts":1605018527.483859,"caller":"example/main.go:34","message":"log","serverName":"go-ecosystem"}
{"level":"info","ts":1605018527.483885,"caller":"example/main.go:35","message":"log with fields","serverName":"go-ecosystem","field":1}
{"level":"error","ts":1605018527.483894,"caller":"example/main.go:37","message":"error log","serverName":"go-ecosystem","stacktrace":"main.main\n\t/Users/catchzeng/Documents/Code/go-ecosystem/log/example/main.go:37\nruntime.main\n\t/usr/local/go/src/runtime/proc.go:203"}
{"level":"error","ts":1605018527.4841652,"caller":"example/main.go:38","message":"error log with fields","serverName":"go-ecosystem","field":"field value","stacktrace":"main.main\n\t/Users/catchzeng/Documents/Code/go-ecosystem/log/example/main.go:38\nruntime.main\n\t/usr/local/go/src/runtime/proc.go:203"}
{"level":"error","ts":1605018527.484202,"caller":"example/main.go:39","message":"error log with message","serverName":"go-ecosystem","error":"test error","stacktrace":"main.main\n\t/Users/catchzeng/Documents/Code/go-ecosystem/log/example/main.go:39\nruntime.main\n\t/usr/local/go/src/runtime/proc.go:203"}
{"level":"panic","ts":1605018527.4842188,"caller":"example/main.go:41","message":"panic","serverName":"go-ecosystem","stacktrace":"main.main\n\t/Users/catchzeng/Documents/Code/go-ecosystem/log/example/main.go:41\nruntime.main\n\t/usr/local/go/src/runtime/proc.go:203"}
2020/11/10 22:28:47 logger sync failed: sync /dev/stderr: inappropriate ioctl for device
panic: panic
        panic: logger sync failed: sync /dev/stderr: inappropriate ioctl for device

goroutine 1 [running]:
log.Panicf(0x1318438, 0x16, 0xc000111960, 0x1, 0x1)
        /usr/local/go/src/log/log.go:345 +0xc0
github.com/go-ecosystem/log.Sync()
        /Users/catchzeng/Documents/Code/go-ecosystem/log/log.go:52 +0x95
panic(0x12a1920, 0xc0000122d0)
        /usr/local/go/src/runtime/panic.go:679 +0x1b2
go.uber.org/zap/zapcore.(*CheckedEntry).Write(0xc0001060c0, 0x0, 0x0, 0x0)
        /Users/catchzeng/go/pkg/mod/go.uber.org/zap@v1.16.0/zapcore/entry.go:234 +0x567
go.uber.org/zap.(*Logger).Panic(0xc0000a8420, 0x1313c02, 0x5, 0x0, 0x0, 0x0)
        /Users/catchzeng/go/pkg/mod/go.uber.org/zap@v1.16.0/logger.go:226 +0x7f
github.com/go-ecosystem/log.Panic(...)
        /Users/catchzeng/Documents/Code/go-ecosystem/log/log.go:125
main.main()
        /Users/catchzeng/Documents/Code/go-ecosystem/log/example/main.go:41 +0x8c4
exit status 2
make: *** [example] Error 1
```

## Configuration

- Production

    ```go
    encoderConf := zap.NewProductionEncoderConfig()
    encoderConf.MessageKey = "message"
    
    conf := zap.NewProductionConfig()
    conf.EncoderConfig = encoderConf
    
    logger, err = conf.Build(zap.AddCaller(),
        zap.AddCallerSkip(1),
        zap.AddStacktrace(zapcore.WarnLevel),
        zap.Fields(fs...))
    ```

- Development

    ```go
    conf := zap.NewDevelopmentConfig()
    logger, err = conf.Build(zap.AddCaller(),
        zap.AddCallerSkip(1),
        zap.AddStacktrace(zapcore.WarnLevel))
    ```