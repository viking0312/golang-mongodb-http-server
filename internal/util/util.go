package util

import "fmt"

func PanicWithError(s string, err error) {
	panic(fmt.Errorf("%s %w", s, err))
}
