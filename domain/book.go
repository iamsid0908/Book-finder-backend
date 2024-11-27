package domain

import (
	"core/config"
	"core/models"
)

type BookDomain interface {
	Insert(param models.Books) error
	GetAll(userID int64) ([]models.BookWithCart, error)
}
type BookDomainCtx struct{}

func (c *BookDomainCtx) Insert(param models.Books) error {
	db := config.DbManager()
	err := db.Create(&param).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *BookDomainCtx) GetAll(userID int64) ([]models.BookWithCart, error) {
	db := config.DbManager()
	result := []models.BookWithCart{}
	err := db.Table("books").
		Select(`
		books.id, books.title, books.thumbnail, books.writter_name, books.created_at, books.updated_at,
		CASE WHEN cart.book_id IS NOT NULL THEN TRUE ELSE FALSE END AS cart
	`).
		Joins("LEFT JOIN cart ON books.id = cart.book_id AND cart.user_id = ?", userID).
		Scan(&result).Error

	if err != nil {
		return []models.BookWithCart{}, err
	}
	return result, nil
}
