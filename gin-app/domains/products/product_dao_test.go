package products

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 正常系テスト
func TestProductSaveNoError(t *testing.T) {
	// Arrange ---
	p := Product{
		ID:        1,
		Name:      "coca cola",
		Detail:    "Wonderful Drink!!!",
		Price:     120,
		Img:       []byte{1, 2, 3},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
	}

	// Act ---
	err := p.Save()

	// Assert ---
	assert.Nil(t, err)
}

// 同一名の商品が保存されたらエラーが発生
func TestProductSaveBadRequestErrorWithSameName(t *testing.T) {
	// Arrange ---
	p := Product{
		ID:        1,
		Name:      "coca cola",
		Detail:    "Wonderful Drink!!!",
		Price:     120,
		Img:       []byte{1, 2, 3},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
	}

	p.Save()

	p2 := Product{
		ID:        1,
		Name:      "coca cola",
		Detail:    "Wonderful Drink!!!",
		Price:     120,
		Img:       []byte{1, 2, 3},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
	}

	// Act ---
	err := p2.Save()

	// Assert ---
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Message, "name coca cola already registered")
	assert.EqualValues(t, err.Status, 400)
	assert.EqualValues(t, err.Error, "bad_request")
}

// 同じIDを保存した場合エラーになる
func TestProductSaveBadRequestErrorWithSameID(t *testing.T) {
	// Arrange ---
	p := Product{
		ID:        1,
		Name:      "coca cola",
		Detail:    "Wonderful Drink!!!",
		Price:     120,
		Img:       []byte{1, 2, 3},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
	}

	p.Save()

	p2 := Product{
		ID:        1,
		Name:      "orange",
		Detail:    "Wonderful Drink!!!",
		Price:     100,
		Img:       []byte{1, 2, 3},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
	}

	// Act ---
	err := p2.Save()

	// Assert ---
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Message, "product 1 already exists")
	assert.EqualValues(t, err.Status, 400)
	assert.EqualValues(t, err.Error, "bad_request")
}

// 正常系テスト
func TestProductGetNoError(t *testing.T) {
	// Arrange ---
	p := Product{
		ID:        1,
		Name:      "coca cola",
		Detail:    "Wonderful Drink!!!",
		Price:     120,
		Img:       []byte{1, 2, 3},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
	}
	p.Save()
	newP := Product{ID: 1}

	// Act ---
	result := newP.Get()

	// Arrange ---
	assert.Nil(t, result)
	assert.EqualValues(t, p.Name, newP.Name)
	assert.EqualValues(t, p.Detail, newP.Detail)
	assert.EqualValues(t, p.Price, newP.Price)
	assert.EqualValues(t, p.Img, newP.Img)
}

// 商品が存在しない場合のテスト
func TestProductNotFound(t *testing.T) {
	// Arrange ---
	p := Product{ID: 100}

	// Act ---
	err := p.Get()

	// Assert ---
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Message, "product 100 not found")
	assert.EqualValues(t, err.Status, 404)
	assert.EqualValues(t, err.Error, "not_found")
}
