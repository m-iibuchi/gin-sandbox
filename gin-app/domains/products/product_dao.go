package products

import (
	"fmt"

	"gin-sandbox/utils/errors"
)

var (
	productsDB = make(map[uint64]*Product)
)

func (p *Product) Get() *errors.ApiErr {
	result := productsDB[p.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("product %d not found", p.ID))
	}

	p.ID = result.ID
	p.Name = result.Name
	p.Detail = result.Detail
	p.Price = result.Price
	p.Img = result.Img
	p.CreatedAt = result.CreatedAt
	p.UpdatedAt = result.UpdatedAt
	p.DeletedAt = result.DeletedAt

	return nil
}

func (p *Product) Save() *errors.ApiErr {
	current := productsDB[p.ID]
	if current != nil {
		if current.Name == p.Name {
			return errors.NewBadRequestError(fmt.Sprintf("name %s already registered", p.Name))
		}
		return errors.NewBadRequestError(fmt.Sprintf("product %d already exists", p.ID))
	}
	productsDB[p.ID] = p
	return nil
}
