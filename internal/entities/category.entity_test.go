package entities

import (
	"testing"
)

func TestItShouldReturnAnErrorNewCategory(t *testing.T) {
	category, err := NewCategory("Test")

	if category != nil {
		t.Errorf("it should not create a category, something is wrong.")
	}

	if err != nil {
		t.Log(err.Error())
	}
}

func TestItShouldCreateNewCategory(t *testing.T) {
	category, _ := NewCategory("Testing Category")

	coming := category.Name
	expected := "Testing Category"

	if coming != expected {
		t.Errorf("category name is diferent, expected %s, Got %s.", expected, coming)
	}
}
