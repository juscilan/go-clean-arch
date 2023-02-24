package entities_test

import (
	"testing"

	"github.com/juscilan/go-clean-arch/internal/entities"
)

func TestItShouldReturnAnErrorNewCategory(t *testing.T) {
	_, err := entities.NewCategory("Test")
	if err == nil {
		t.Errorf("it should not create a category, something is wrong.")
	}

}

func TestItShouldCreateNewCategory(t *testing.T) {
	category, _ := entities.NewCategory("Testing Category")
	output := category.Name
	expected := "Testing Category"

	if output != expected {
		t.Errorf("category name is diferent, expected %s, Got %s.", expected, output)
	}
}
