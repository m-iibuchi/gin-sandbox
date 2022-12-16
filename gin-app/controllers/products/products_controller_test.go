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
