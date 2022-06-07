package main

import (
	"fmt"

	"github.com/thalesmacedo1/go-repository-pattern/entity"
	"github.com/thalesmacedo1/go-repository-pattern/repository"
	"github.com/thalesmacedo1/go-repository-pattern/service"
)

func main() {
	db := repository.CategoriesMemoryDb(Categories: []entity.Category{})
	repositoryMemory := repository.NewCategoryRepositoryMemory(db)

	service := service.NewCategoryService(repositoryMemory)
	cat, _ := service.Create("MY CATEGORY")

	fmt.Println(cat)
}
