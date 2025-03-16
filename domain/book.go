package domain

import (
	"core/config"
	"core/models"
)

type BookDomain interface {
	Get() ([]models.Books, error)
	Insert(param models.Books) (int64, error)
	GetAll(param models.SearchByInputParam) ([]models.BookWithCart, int64, error)
	GetCategory(param models.RecommendReqs) (models.Books, error)
}
type BookDomainCtx struct{}

func (c *BookDomainCtx) Get() ([]models.Books, error) {
	db := config.DbManager()
	limit := 10
	offset := 0
	books := []models.Books{}
	err := db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil

}
func (c *BookDomainCtx) Insert(param models.Books) (int64, error) {
	db := config.DbManager()
	err := db.Create(&param).Error
	if err != nil {
		return 0, err
	}

	return param.ID, nil
}

func (c *BookDomainCtx) GetAll(param models.SearchByInputParam) ([]models.BookWithCart, int64, error) {
	db := config.DbManager()
	result := []models.BookWithCart{}
	var totalCount int64

	query := db.Table("books").
		Select(`
		books.id, books.title, books.thumbnail, books.description, books.writter_name, books.created_at, books.updated_at,
		CASE WHEN cart.book_id IS NOT NULL THEN TRUE ELSE FALSE END AS cart
	`).
		Joins("LEFT JOIN cart ON books.id = cart.book_id AND cart.user_id = ?", param.UserID)

	// Apply filters
	if param.WritterName != "" && param.Title != "" {
		query = query.Where("writter_name ILIKE ? OR title ILIKE ?", "%"+param.WritterName+"%", "%"+param.Title+"%")
	} else if param.WritterName != "" {
		query = query.Where("writter_name ILIKE ?", "%"+param.WritterName+"%")
	} else if param.Title != "" {
		query = query.Where("title ILIKE ?", "%"+param.Title+"%")
	}

	// Get total count (without LIMIT & OFFSET)
	if err := query.Count(&totalCount).Error; err != nil {
		return []models.BookWithCart{}, 0, err
	}

	// Apply pagination
	limit := int64(8)
	offset := int64(0)
	if param.Limit > 0 {
		limit = param.Limit
	}
	if param.Page > 0 {
		offset = (param.Page - 1) * limit
	}

	err := query.Limit(int(limit)).Offset(int(offset)).Scan(&result).Error
	if err != nil {
		return []models.BookWithCart{}, 0, err
	}

	return result, totalCount, nil
}

func (c *BookDomainCtx) GetCategory(param models.RecommendReqs) (models.Books, error) {
	db := config.DbManager()
	result := models.Books{}
	err := db.Where("id = ?", param.BookID).First(&result).Error
	if err != nil {
		return models.Books{}, err
	}
	return result, nil
}
