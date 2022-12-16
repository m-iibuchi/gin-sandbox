package products

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-sandbox/domains/products"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateProductNoError(t *testing.T) {
	// テストデータ準備
	p := products.Product{
		ID:   123,
		Name: "coca cola",
	}
	byteProduct, _ := json.Marshal(p)

	// リクエスト準備
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodPost,
		"/products",
		bytes.NewBuffer(byteProduct),
	)

	// リクエスト処理
	CreateProduct(c)

	// レスポンス整形
	var product products.Product
	err := json.Unmarshal(response.Body.Bytes(), &product)

	// 検証
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.Nil(t, err)
	fmt.Println(product)
	assert.EqualValues(t, uint64(123), product.ID)
}
