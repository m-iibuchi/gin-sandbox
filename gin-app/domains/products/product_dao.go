package products

import (
	"gin-sandbox/datasources/mysql/products_db"
	"gin-sandbox/utils/errors"
	"gin-sandbox/utils/mysqlutils"
)

const (
	noSearchResult = "record not found"
)

var (
	productsDB = make(map[uint]*Product)
)

func (p *Product) Get() *errors.ApiErr {
	if result := products_db.Client.Where("id = ?", p.ID).Take(&p); result.Error != nil {
		return mysqlutils.ParseError(result.Error)
	}
	return nil
}

func (p *Product) Save() *errors.ApiErr {
	if result := products_db.Client.Create(&p); result.Error != nil {
		return mysqlutils.ParseError(result.Error)
	}
	return nil
}

func (p *Product) Update() *errors.ApiErr {
	if result := products_db.Client.Save(&p); result.Error != nil {
		return mysqlutils.ParseError(result.Error)
	}
	return nil
}

func (p *Product) PartialUpdate() *errors.ApiErr {
	if result := products_db.Client.
		Table("products").
		Where("id IN (?)", p.ID).
		Updates(&p); result.Error != nil {
		return mysqlutils.ParseError(result.Error)
	}
	return nil
}
