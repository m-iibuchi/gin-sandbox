package products

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
    c.String(http.StatusNotImplemented, "Not yet implement!")
}

func CreateProduct(c *gin.Context) {
    c.String(http.StatusNotImplemented, "Not yet implement!")
}