package repositories

import "github.com/juscilan/go-clean-arch/internal/entities"

type CategoryRepositoryInterface interface {
	Save(category *entities.Category) error
	List() ([]*entities.Category, error)
}
