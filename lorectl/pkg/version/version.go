package version

import (
	"fmt"
)

const (
	major = "0"
	minor = "0"
	patch = "0"
)

var Version = fmt.Sprintf("%s.%s.%s", major, minor, patch)
