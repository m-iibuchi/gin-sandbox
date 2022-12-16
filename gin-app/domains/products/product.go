package products

import (
	"gin-sandbox/utils/errors"
	"strings"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name   string `json:"name"`
	Detail string `json:"detail"`
	Price  uint64 `json:"price"`
	Img    []byte `json:"img"`
}

func (p *Product) Validate() *errors.ApiErr {
	p.Name = strings.TrimSpace(strings.ToLower(p.Name))
	if p.Name == "" {
		return errors.NewBadRequestError("invalid product name")
	}
	return nil
}
