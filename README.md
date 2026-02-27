# go-binver

common binary versioning for go

Usage:
```go
import (
	"fmt"
	"github.com/jda/go-binver"
)

func ExampleCanonVersion() {
	fmt.Println(binver.CanonVersion())
}
```

`CanonVersion` returns the build’s VCS revision (full commit hash) or `NO_VERSION` when build metadata is missing. If the working tree was dirty when the binary was built, the suffix `-tainted` is appended.

Example output:
```
d34db33f12c0ffee9badbeef00bab10cdef12345
d34db33f12c0ffee9badbeef00bab10cdef12345-tainted
NO_VERSION
```
