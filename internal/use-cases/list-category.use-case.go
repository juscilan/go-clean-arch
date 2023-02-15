package usecases

import (
	"github.com/juscilan/go-clean-arch/internal/entities"
	"github.com/juscilan/go-clean-arch/internal/repositories"
)

type listCategoryUseCase struct {
	repository repositories.CategoryRepositoryInterface
}

func NewListCategoryUseCase(repository repositories.CategoryRepositoryInterface) *listCategoryUseCase {
	return &listCategoryUseCase{repository}
}

func (u *listCategoryUseCase) Execute() ([]*entities.Category, error) {

	categories, err := u.repository.List()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
