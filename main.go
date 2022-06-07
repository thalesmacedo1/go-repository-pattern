package main

import (
	"database/sql"
	"fmt"

	"github.com/thalesmacedo1/go-repository-pattern/repository"
	"github.com/thalesmacedo1/go-repository-pattern/service"
)

func main() {
	// db := repository.CategoriesMemoryDb{Categories: []entity.Category{}}
	// repositoryMemory := repository.NewCategoryRepositoryMemory(db)

	db, _ := sql.Open("sqlite3", "./sqlite.repository")
	repositorySqlite := repository.NewCategorySqlite(db)

	service := service.NewCategoryService(repositorySqlite)
	cat, _ := service.Create("MY CATEGORY")

	fmt.Println(cat)
}
