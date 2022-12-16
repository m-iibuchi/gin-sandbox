package products

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-sandbox/domains/products"
	"gin-sandbox/utils/errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func requestHandler(p interface{}) (*gin.Context, *httptest.ResponseRecorder) {
	// リクエスト準備
	response := httptest.NewRecorder()
	byteProduct, _ := json.Marshal(p)
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodPost,
		"/products",
		bytes.NewBuffer(byteProduct),
	)
	return c, response
}

func getRequestHandler(id string) (*gin.Context, *httptest.ResponseRecorder) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)

	param := gin.Param{Key: "product_id", Value: id}
	c.Params = gin.Params{param}

	c.Request, _ = http.NewRequest(
		http.MethodGet,
		"/products/:product_id",
		nil,
	)

	return c, response
}

func TestCreateProductNoError(t *testing.T) {
	// テストデータ準備
	p := products.Product{
		ID:   123,
		Name: "coca cola",
	}

	// リクエスト準備
	c, response := requestHandler(p)

	// リクエスト処理
	CreateProduct(c)

	// レスポンス整形
	var product products.Product
	err := json.Unmarshal(response.Body.Bytes(), &product)

	// 検証
	assert.EqualValues(t, http.StatusCreated, response.Code)
	assert.Nil(t, err)
	fmt.Println(product)
	assert.EqualValues(t, uint64(123), product.ID)
}

func TestCreateProductWith404Error(t *testing.T) {
	// テストデータ準備
	type demiProduct struct {
		ID   string `json: "id"`
		Name string `json: "name"`
	}
	p := demiProduct{
		ID:   "123",
		Name: "coca cola",
	}

	// リクエスト準備
	c, response := requestHandler(p)

	// リクエスト処理
	CreateProduct(c)

	var apiErr errors.ApiErr
	err := json.Unmarshal(response.Body.Bytes(), &apiErr)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	assert.Nil(t, err)
	assert.EqualValues(t, "invalid json body", apiErr.Message)
	assert.EqualValues(t, 400, apiErr.Status)
	assert.EqualValues(t, "bad_request", apiErr.Error)
}

func TestProductValidateNoError(t *testing.T) {
	p := products.Product{ID: 123, Name: "coca cola"}

	err := p.Validate()

	assert.Nil(t, err)
}

func TestProductValidateBadRequestError(t *testing.T) {
	p := products.Product{ID: 123}

	err := p.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid product name", err.Message)
	assert.EqualValues(t, 400, err.Status)
	assert.EqualValues(t, "bad_request", err.Error)
}

// 正常系
func TestGetProductNoError(t *testing.T) {
	// Arrange
	p := products.Product{ID: 1, Name: "coca cola"}
	c, _ := requestHandler(p)
	CreateProduct(c)

	c2, response := getRequestHandler("1")

	// Act ---
	GetProduct(c2)

	// Assert ---
	var product products.Product
	err := json.Unmarshal(response.Body.Bytes(), &product)
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.Nil(t, err)
	assert.EqualValues(t, uint64(1), product.ID)
}

// 不正なIDのテスト
func TestGetProductWithInvalidID(t *testing.T) {
	// Arrange
	c, response := getRequestHandler("a")

	// Act ---
	GetProduct(c)

	// Assert ---
	var apiErr errors.ApiErr
	json.Unmarshal(response.Body.Bytes(), &apiErr)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, apiErr.Message, "product id should be a number")
	assert.EqualValues(t, apiErr.Status, 400)
	assert.EqualValues(t, apiErr.Error, "bad_request")
}

// 指定したIDのプロダクトが存在しないテスト
func TestGetProductWithNoProduct(t *testing.T) {
	// Arrange ---
	c, response := getRequestHandler("10000")

	// Act ---
	GetProduct(c)

	// Assert ---
	var apiErr errors.ApiErr
	json.Unmarshal(response.Body.Bytes(), &apiErr)
	assert.EqualValues(t, http.StatusNotFound, response.Code)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, apiErr.Message, "product 10000 not found")
	assert.EqualValues(t, apiErr.Status, 404)
	assert.EqualValues(t, apiErr.Error, "not_found")
}
