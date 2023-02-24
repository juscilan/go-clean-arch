package entities

import "fmt"

type Product struct {
	Name  string
	Price int64
}

func (p *Product) isValid() error {
	if p.Price <= 0 {
		return fmt.Errorf("price must not be less or equal than zero")
	}
	return nil
}

func NewProduct(name string, price int64) (*Product, error) {
	product := &Product{
		Name:  name,
		Price: price,
	}

	err := product.isValid()
	if err != nil {
		return nil, err
	}
	return product, nil
}
