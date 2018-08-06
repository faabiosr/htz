# Htz

[![Build Status](https://img.shields.io/travis/fabiorphp/htz/master.svg?style=flat-square)](https://travis-ci.org/fabiorphp/htz)
[![Codecov branch](https://img.shields.io/codecov/c/github/fabiorphp/htz/master.svg?style=flat-square)](https://codecov.io/gh/fabiorphp/htz)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/fabiorphp/htz)
[![Go Report Card](https://goreportcard.com/badge/github.com/fabiorphp/htz?style=flat-square)](https://goreportcard.com/report/github.com/fabiorphp/htz)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](https://github.com/fabiorphp/htz/blob/master/LICENSE)

Go healthcheck for your apps

## Instalation

Htz requires Go 1.10 or later.

```
go get github.com/fabiorphp/htz
```

If you want to get an specific version, please use the example below:

```
go get gopkg.in/fabiorphp/htz.v0
```

## Usage

### Simple usage
```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/fabiorphp/htz"
	"time"
)

func main() {
	checkers := []htz.Checker{
		func() *htz.Check {
			return &htz.Check{
				Name:         "some-api",
				Type:         "internal-service",
				Status:       false,
				ResponseTime: 6 * time.Second,
				Optional:     false,
				Details: map[string]interface{}{
					"url": "internal-service.api",
				},
			}
		},
	}

	h := htz.New("my-app", "0.0.1", checkers...)
	res, _ := json.MarshalIndent(h.Check(), "", "  ")

	fmt.Println(string(res))
}
```

### HTTP Check
```go
package main

import (
	"github.com/fabiorphp/htz"
	"net/http"
	"time"
)

func main() {
	checkers := []htz.Checker{
		func() *htz.Check {
			return &htz.Check{
				Name:         "some-api",
				Type:         "internal-service",
				Status:       false,
				ResponseTime: 6 * time.Second,
				Optional:     false,
				Details: map[string]interface{}{
					"url": "internal-service.api",
				},
			}
		},
	}

	h := htz.New("my-app", "0.0.1", checkers...)

	http.Handle("/htz", h)

	http.ListenAndServe(":8080", nil)
}
```

## Available checkers

### DB
```go
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/fabiorphp/htz"
	"github.com/fabiorphp/htz/checker"
	_"github.com/lib/pq"
)

func main() {
	db, _ := sql.Open("postgres", "user=htz dbname=htz")

	h := htz.New("my-app", "0.0.1", checker.DB(db, true))
	res, _ := json.MarshalIndent(h.Check(), "", "  ")

	fmt.Println(string(res))
}
```

### Redis
```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/fabiorphp/htz"
	"github.com/fabiorphp/htz/checker"
    "github.com/go-redis/redis"
)

func main() {
    client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })

	h := htz.New("my-app", "0.0.1", checker.Redis(client, true))
	res, _ := json.MarshalIndent(h.Check(), "", "  ")

	fmt.Println(string(res))
}
```

### Runtime
```go
package main

import (
	"fmt"
	"github.com/fabiorphp/htz"
	"github.com/fabiorphp/htz/checker"
)

func main() {
	h := htz.New("my-app", "0.0.1", checker.Runtime(false))
	res, _ := json.MarshalIndent(h.Check(), "", "  ")

	fmt.Println(string(res))
}
```

## Development

### Requirements

- Install [Go](https://golang.org)
- Install [go dep](https://github.com/golang/dep)

### Makefile
```sh
// Clean up
$ make clean

// Creates folders and download dependencies
$ make configure

// Run tests and generates html coverage file
make cover

// Download project dependencies
make depend

// Format all go files
make fmt

// Run linters
make lint

// Run tests
make test
```

## License

This project is released under the MIT licence. See [LICENSE](https://github.com/fabiorphp/htz/blob/master/LICENSE) for more details.
