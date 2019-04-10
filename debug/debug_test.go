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
	assertEqual(t, "testing", output)
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
	assertEqual(t, "", output)
}
