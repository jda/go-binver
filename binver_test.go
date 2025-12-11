package binver_test

import (
	"fmt"
	"scm.0xbad.coffee/jda/go-binver"
)

func ExampleCanonVersion() {
	fmt.Println(binver.CanonVersion())
	// Output: NO_VERSION
}