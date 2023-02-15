package usecases

import (
	"github.com/juscilan/go-clean-arch/internal/entities"
	"github.com/juscilan/go-clean-arch/internal/interfaces"
)

type createCategoryUseCase struct {
	repository interfaces.CategoryRepositoryInterface
}

func NewCreateCategoryUseCase(repository interfaces.CategoryRepositoryInterface) *createCategoryUseCase {
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
