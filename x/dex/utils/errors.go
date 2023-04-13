package utils

import (
	"fmt"
	"strings"
)

func JoinErrors(parentError error, errors ...error) error {
	// Can be used for bundling multiple errors. For now only the parent error and first error in []errors
	// will be matched by ErrorIs.
	// TODO: switch to errors.Join when we bump to golang 1.20
	errorFmt := strings.Repeat("%w", len(errors))
	return fmt.Errorf("%w Additional errors: %w %v "+errorFmt, parentError, errors[0], errors[1:])
}
