package usecases

import (
	"github.com/juscilan/go-clean-arch/internal/entities"
	"github.com/juscilan/go-clean-arch/internal/repositories"
)

type createCategoryUseCase struct {
	repository repositories.CategoryRepositoryInterface
}

func NewCreateCategoryUseCase(repository repositories.CategoryRepositoryInterface) *createCategoryUseCase {
	return &createCategoryUseCase{repository}
}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)
	if err != nil {
		return err
	}

	err = u.repository.Save(category)
	if err != nil {
		return err
	}

	return nil
}
