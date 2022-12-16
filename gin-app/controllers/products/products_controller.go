package products

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-sandbox/domains/products"
	"gin-sandbox/services"
	"gin-sandbox/utils/errors"
)

func GetProduct(c *gin.Context) {
	productID, productErr := strconv.ParseUint(c.Param("product_id"), 10, 64)
	if productErr != nil {
		err := errors.NewBadRequestError("product id should be a number")
		c.JSON(err.Status, err)
		return
	}

	product, getErr := services.GetProduct(uint(productID))
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	// リクエストデータ取得
	var product products.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status, apiErr)
		return
	}

	// 登録処理
	newProduct, saveErr := services.CreateProduct(product)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, newProduct)
}
