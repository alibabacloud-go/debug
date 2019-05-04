# Debug for Golang

- [![Build Status](https://travis-ci.org/alibabacloud-go/debug.svg?branch=master)](https://travis-ci.org/alibabacloud-go/debug)
- [![codecov](https://codecov.io/gh/alibabacloud-go/debug/branch/master/graph/badge.svg)](https://codecov.io/gh/alibabacloud-go/debug)

## Usage

If you want to use the `Debug` function, you need to add the
following code to your project:

```go
// Step 0: Import it
import "github.com/alibabacloud-go/debug/debug"

// Step 1: Initialize variables, you can replace flag with
var d = debug.Init("sdk")

// Step 2: Use it
func main() {
  d("this debug information just print when DEBUG environment variable was set")
}
```

Just run it with `go run demo.go`:

```sh
go run demo.go
```

No any results.

Set DEBUG environment variable with `export DEBUG=sdk` and run again:

```sh
$ DEBUG=sdk go run demo.go
this debug information just print when DEBUG environment variable was set
```

## Testing

Run `go test -v ./debug`

## License

The Apache License 2.0.
