package entities_test

import (
	"testing"

	"github.com/juscilan/go-clean-arch/internal/entities"
)

const PRICE_ZERO = 0

func TestItShouldNotCreateANewProductPriceZero(t *testing.T) {
	_, err := entities.NewProduct("Product Number One", PRICE_ZERO)

	if err == nil {
		t.Errorf("it should be an error")
	}
}

func TestProductCreateSucessfully(t *testing.T) {
	product, _ := entities.NewProduct("Product Number One", 10000)

	coming := product.Name
	expected := "Product Number One"

	if coming != expected {
		t.Errorf("it should create a new product")
	}
}
