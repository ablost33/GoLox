package error

import (
	"fmt"
)

func report(line int, where string, msg string) {
	fmt.Printf("[line %d] Error %s: %s", line, where, msg)
}

func ReportError(line int, msg string) {
	report(line, "", msg)
}
