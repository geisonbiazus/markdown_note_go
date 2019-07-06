package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertValid(t *testing.T, r ValidationResult) {
	t.Helper()
	assert.Equal(t, ValidationResult{
		Valid:  true,
		Errors: []Error{},
	}, r)
}

func assertValidationError(t *testing.T, r ValidationResult, field, errType string) {
	t.Helper()
	assert.Equal(t, ValidationResult{
		Valid: false,
		Errors: []Error{
			Error{
				Field: "Title",
				Type:  "REQUIRED",
			},
		},
	}, r)
}
