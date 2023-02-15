package entities

import (
	"fmt"
	"time"
)

type category struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"uodated_at"`
}

func (c *category) isValid() error {
	if len(c.Name) < 5 {
		return fmt.Errorf("Name must be greater than 5. Got %d", len(c.Name))
	}
	return nil
}

func NewCategory(name string) (*category, error) {
	category := &category{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	//createonal caregory business rules
	err := category.isValid()
	if err != nil {
		return nil, err
	}

	return category, nil
}
