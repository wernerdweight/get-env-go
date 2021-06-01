Get ENV helper for Go
====================================

This package helps with using ENV files. It looks for an ENV file in the current working directory and any parent directories until the point it finds an ENV file or reaches the root.

[![Build Status](https://www.travis-ci.com/wernerdweight/get-env-go.svg?branch=master)](https://www.travis-ci.com/wernerdweight/get-env-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/wernerdweight/get-env-go)](https://goreportcard.com/report/github.com/wernerdweight/get-env-go)
[![GoDoc](https://godoc.org/github.com/wernerdweight/get-env-go?status.svg)](https://godoc.org/github.com/wernerdweight/get-env-go)
[![go.dev](https://img.shields.io/badge/go.dev-pkg-007d9c.svg?style=flat)](https://pkg.go.dev/github.com/wernerdweight/get-env-go)


Installation
------------

### 1. Installation

```bash
go get github.com/wernerdweight/get-env-go
```

Configuration
------------

The package needs no configuration.

Usage
------------

By default, this package expects `.env.local` file. You can choose a different file when initializing (see an example below).

**Basic usage**

```go
package main

import (
    "fmt"
    "github.com/wernerdweight/get-env-go/getenv"
)

// let's pretend .env.local file exists and contains the following:
// APP_ENV=prod

func main() {
    err := getenv.Init()
    if nil != err {
        // handle the error (e.g. no env file exists or not readable)
    }
	env, err := getenv.GetEnv("APP_ENV")
    if nil != err {
    	// handle the error (requested env var doesn't exist)
    }
    fmt.Printf("APP_ENV: %s", env) // prints "prod"
}
```

**Custom ENV file**

```go
package main

import (
    "fmt"
    "github.com/wernerdweight/get-env-go/getenv"
)

// let's pretend my-file.txt file exists and contains the following:
// APP_ENV=prod

func main() {
    err := getenv.InitFromFile("my-file.txt")
    if nil != err {
        // handle the error (e.g. no env file exists or not readable)
    }
	env, err := getenv.GetEnv("APP_ENV")
    if nil != err {
    	// handle the error (requested env var doesn't exist)
    }
    fmt.Printf("APP_ENV: %s", env) // prints "prod"
}
```

**Errors**

The following errors can occur (you can check for specific code since different errors have different severity):

```go
package main

import (
	"fmt"
	"github.com/wernerdweight/get-env-go/getenv"
)

func main() {
    // these are the possible err.Code values
    fmt.Print(
        getenv.CwdFailureError,  // "can not determine current working directory"
        getenv.CantAccessEnvFileError, // "error accessing an existing env file '[env-file-name]'"
        getenv.NoEnvFileError, // "no [env-file-name] file exists in any of parent directories"
        getenv.CantLoadEnvFileError, // "can't load [env-file-name] file"
        getenv.NoSuchEnvVarError, // "non-existing ENV variable [env-var-name] requested"
    )
    _, err := getenv.GetEnv("APP_ENV")
    if nil != err {
        // would print "non-existing ENV variable [env-var-name] requested"
    	fmt.Print(
            err.(*getenv.Error).Code,
        )
    }
}
```

License
-------
This package is under the MIT license. See the complete license in the root directory of the bundle.
