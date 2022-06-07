package service

import "github.com/thalesmacedo1/go-repository-pattern/entity"

type CategoryService struct {
	Repository entity.CategoryRepository
}

func NewCategoryService(repository entity.CategoryRepository) *CategoryService {
	return &CategoryService{Repository: repository}
}

func (c *CategoryService) findById(id string) (entity.Category, error) {
	category, err := c.Repository.Get(id)
	if err != nil {
		return entity.Category{}, err
	}
	return category, nil
}

func (c *CategoryService) Create(name string) (entity.Category, error) {
	category := entity.NewCategory()
	category.Name = name
	cat, err := c.Repository.Create(*category)
	if err != nil {
		return entity.Category{}, err
	}
	return cat, nil
}
