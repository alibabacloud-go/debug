package debug

import (
	"testing"
)

func Test_Main(t *testing.T) {
	debug := Init("sdk")
	debug("%s", "testing")
}

func TestMain_Matched(t *testing.T) {
	originLookGetEnv := hookGetEnv
	originhookPrint := hookPrint
	defer func() {
		hookGetEnv = originLookGetEnv
		hookPrint = originhookPrint
	}()
	hookGetEnv = func() string {
		return "sdk"
	}
	output := ""
	hookPrint = func(input string) {
		output = input
		originhookPrint(input)
	}
	debug := Init("sdk")
	debug("%s", "testing")
	if output != "testing" {
		t.Errorf("%v != %v", output, "testing")
	}
}

func TestMain_UnMatched(t *testing.T) {
	originLookGetEnv := hookGetEnv
	originhookPrint := hookPrint
	defer func() {
		hookGetEnv = originLookGetEnv
		hookPrint = originhookPrint
	}()
	hookGetEnv = func() string {
		return "non-sdk"
	}
	output := ""
	hookPrint = func(input string) {
		output = input
		originhookPrint(input)
	}
	debug := Init("sdk")
	debug("%s", "testing")
	if output != "" {
		t.Errorf("output failed")
	}
}
