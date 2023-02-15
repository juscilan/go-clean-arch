package usecases

import (
	"github.com/juscilan/go-clean-arch/internal/entities"
	"github.com/juscilan/go-clean-arch/internal/interfaces"
)

type listCategoryUseCase struct {
	repository interfaces.CategoryRepositoryInterface
}

func NewListCategoryUseCase(repository interfaces.CategoryRepositoryInterface) *listCategoryUseCase {
	return &listCategoryUseCase{repository}
}

func (u *listCategoryUseCase) Execute() ([]*entities.Category, error) {

	categories, err := u.repository.List()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
