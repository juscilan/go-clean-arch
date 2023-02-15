package repositories

import "github.com/juscilan/go-clean-arch/internal/entities"

type categoryInMemoryRepository struct {
	db []*entities.Category
}

func NewCategoryInMemoryRepository() *categoryInMemoryRepository {
	return &categoryInMemoryRepository{
		db: make([]*entities.Category, 0),
	}
}

func (r *categoryInMemoryRepository) Save(category *entities.Category) error {
	r.db = append(r.db, category)
	return nil
}
