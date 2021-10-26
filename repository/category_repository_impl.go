package repository

import (
	"api-test/helper"
	"api-test/model/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (c *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	var id int
	SQL := "INSERT INTO category(name) VALUES($1) RETURNING id"
	err := tx.QueryRow(SQL, category.Name).Scan(&id)
	helper.PanicIfError(err)

	// if using mysql
	// SQL := "INSERT INTO category(name) VALUES(?)"
	// result, err := tx.ExecContext(ctx, SQL, category.Name)
	// helper.PanicIfError(err)

	// id, err := result.LastInsertId()
	// helper.PanicIfError(err)

	category.Id = int(id)

	return category
}

func (c *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	// if using mysql
	// SQL := "UPDATE category set name = ? where id = ?"
	SQL := "UPDATE category set name = $1 where id = $2"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (c *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	// if using mysql
	// SQL := "DELETE from category where id = ?"
	SQL := "DELETE from category where id = $1"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}

func (c *CategoryRepositoryImpl) FindId(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	// if using sql
	// SQL := "SELECT id, name FROM category where id = ?"
	SQL := "SELECT id, name FROM category where id = $1"
	row, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)

	category := domain.Category{}
	if row.Next() {
		err := row.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("Category not found")
	}
}

func (c *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT * FROM category"
	row, err := tx.Query(SQL)
	helper.PanicIfError(err)

	var categories []domain.Category
	var order int = 0
	for row.Next() {
		order++
		category := domain.Category{}
		row.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}

	fmt.Println(order)
	return categories
}
