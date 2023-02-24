package entities_test

import (
	"testing"

	"github.com/juscilan/go-clean-arch/internal/entities"
)

func TestItShouldReturnAnErrorNewCategory(t *testing.T) {
	category, err := entities.NewCategory("Test")

	if category != nil {
		t.Errorf("it should not create a category, something is wrong.")
	}

	if err != nil {
		t.Log(err.Error())
	}
}

func TestItShouldCreateNewCategory(t *testing.T) {
	category, _ := entities.NewCategory("Testing Category")

	coming := category.Name
	expected := "Testing Category"

	if coming != expected {
		t.Errorf("category name is diferent, expected %s, Got %s.", expected, coming)
	}
}
