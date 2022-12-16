package products

import (
	"gin-sandbox/utils/errors"
	"strings"
	"time"
)

type Product struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Detail    string `json:"detail"`
	Price     uint64 `json:"price"`
	Img       []byte `json:"img"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (p *Product) Validate() *errors.ApiErr {
	p.Name = strings.TrimSpace(strings.ToLower(p.Name))
	if p.Name == "" {
		return errors.NewBadRequestError("invalid product name")
	}
	return nil
}
