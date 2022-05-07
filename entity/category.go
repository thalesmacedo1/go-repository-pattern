package entity

import uuid "github.com/satori/go.uuid"

type CategoryRepository interface {
	Get(id string) (Category, error)
	Create(category Category) (Category, error)
}

type Category struct {
	ID   string
	Name string
}

func NewCategory() *Category {
	category := Category{
		ID: uuid.NewV4().String(),
	}
	return &category
}
