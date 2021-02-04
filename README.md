# Go Environment

Simple environment variable parser.

## Installation

```shell
go get github.com/ensarkovankaya/goenv
```

## Usage

```go
// Import
import "github.com/ensarkovankaya/goenv"

// Environment variable, Required, Default value
var Secret = goenv.String("SECRET", false, "") // type: string
var Debug = goenv.Bool("DEBUG", false, false)  // type: bool
var Timeout = goenv.Int64("TIMEOUT", false, 0) // type: int64

// If "API_KEY" environment not defined panics
var APIKey = goenv.String("API_KEY", true, "")
```
