package usecases

import "github.com/juscilan/go-clean-arch/internal/entities"

type createCategoryUseCase struct {
}

func NewCreateCategoryUseCase() *createCategoryUseCase {
	return &createCategoryUseCase{}
}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)
	if err != nil {
		return err
	}

	println(category)

	return nil
}
