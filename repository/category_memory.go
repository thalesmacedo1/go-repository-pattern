package repository

import (
	"errors"

	"github.com/thalesmacedo1/go-repository-pattern/entity"
)

type CategoriesMemoryDb struct {
	Categories []entity.Category
}

type CategoryRepositoryMemory struct {
	db CategoriesMemoryDb
}

func (c *CategoryRepositoryMemory) Get(id string) (entity.Category, error) {
	for _, category := range c.db.Categories {
		if category.ID == id {
			return category, nil
		}
	}
	return entity.Category{}, errors.New("no category found with this ID")
}

func (c *CategoryRepositoryMemory) Create(category entity.Category) (entity.Category, error) {
	c.db.Categories = append(c.db.Categories, category)
	return category, nil
}
