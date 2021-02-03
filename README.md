# Go Environment

Simple environment variable parser.

## Installation

```shell
go get github.com/ensarkovankaya/goenv
```

## Usage

```go
// Import
import (
    env "github.com/ensarkovankaya/goenv"
)

// Environment variable, Required, Default value
var Secret := env.String("SECRET", false, "") // type: string
var Debug := env.Bool("DEBUG", false, false)  // type: bool
var Timeout := env.Int64("TIMEOUT", false, 0) // type: int64

// If "API_KEY" environment not defined panics
var APIKey := env.String("API_KEY", true, "")
```
