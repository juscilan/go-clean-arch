package entities

import (
	"testing"
)

func TestItShouldReturnAnErrorNewCategory(t *testing.T) {
	category, err := NewCategory("Test")

	if category == nil {
		t.Log(err.Error())
	}
}

func TestItShouldCreateNewCategory(t *testing.T) {
	category, err := NewCategory("Testing Category")

	if category == nil {
		t.Errorf(err.Error())
	}
}
