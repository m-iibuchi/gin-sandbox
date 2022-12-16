package products

import (
	"fmt"

	"gin-sandbox/datasources/mysql/products_db"
	"gin-sandbox/utils/errors"
)

var (
	productsDB = make(map[uint]*Product)
)

func (p *Product) Get() *errors.ApiErr {
	if result := products_db.Client.Where("id = ?", p.ID).Take(&p); result.Error != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to get product: %s", result.Error),
		)
	}
	return nil
}

func (p *Product) Save() *errors.ApiErr {
	if result := products_db.Client.Create(&p); result.Error != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save product: %s", result.Error),
		)
	}
	return nil
}
