package main

import "github.com/alibabacloud-go/debug/debug"

var d = debug.Init("sdk")

func main() {
	// try `go run demo.go`
	// and `DEBUG=sdk go run demo.go`
	d("this debug information just print when DEBUG environment variable was set")
}
