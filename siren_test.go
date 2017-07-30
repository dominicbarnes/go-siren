package siren_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Validator interface {
	Validate() error
}

type ValidatorSpec struct {
	Input    Validator
	Expected error
}

func RunValidatorTests(t *testing.T, specs map[string]ValidatorSpec) {
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			err := spec.Input.Validate()
			if spec.Expected != nil {
				assert.EqualError(t, err, spec.Expected.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
