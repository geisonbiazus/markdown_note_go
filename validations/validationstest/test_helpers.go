package validationstest

import (
	"testing"

	"github.com/geisonbiazus/markdown_notes/validations"
	"github.com/stretchr/testify/assert"
)

func AssertValid(t *testing.T, r validations.Result) {
	t.Helper()
	assert.Equal(t, validations.Result{
		Valid:  true,
		Errors: []validations.Error{},
	}, r)
}

func AssertValidationError(t *testing.T, r validations.Result, field, errType string) {
	t.Helper()
	assert.Contains(t, r.Errors,
		validations.Error{
			Field: field,
			Type:  errType,
		},
	)
}
