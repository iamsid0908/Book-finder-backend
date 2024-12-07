package domain

import (
	"core/config"
	"core/models"
)

type BookSummaryDomain interface {
	Insert(param models.BookSummary) error
	GetBookDetails(param models.BookSummary) (models.GetBookSummaryDetailsResp, error)
}
type BookSummaryDomainCtx struct{}

func (c *BookSummaryDomainCtx) Insert(param models.BookSummary) error {
	db := config.DbManager()
	err := db.Create(&param).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *BookSummaryDomainCtx) GetBookDetails(param models.BookSummary) (models.GetBookSummaryDetailsResp, error) {
	db := config.DbManager()
	result := models.GetBookSummaryDetailsResp{}
	err := db.Table("books").
		Select("books.id, books.title, books.thumbnail, books.writter_name,books.category, book_summary.summary, book_summary.author_details, book_summary.published_date").
		Joins("inner join book_summary on books.id = book_summary.book_id").
		Where("books.id = ?", param.BookID).
		First(&result).Error

	if err != nil {
		return models.GetBookSummaryDetailsResp{}, err
	}
	return result, nil
}
