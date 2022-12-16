package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestProductValidateNoError(t *testing.T) {
	p := Product{Model: gorm.Model{ID: 123}, Name: "coca cola"}

	err := p.Validate()

	assert.Nil(t, err)
}

func TestProductValidateBadRequestError(t *testing.T) {
	p := Product{Model: gorm.Model{ID: 123}}

	err := p.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid product name", err.Message)
	assert.EqualValues(t, 400, err.Status)
	assert.EqualValues(t, "bad_request", err.Error)
}
