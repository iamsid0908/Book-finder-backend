package domain

import (
	"core/config"
	"core/models"
)

type CartDomain interface {
	Insert(param models.Cart) error
	Get(usetID int64) ([]models.CartResp, error)
	RemoveFromCart(param models.Cart) error
}

type CartDomainCtx struct{}

func (c *CartDomainCtx) Insert(param models.Cart) error {
	db := config.DbManager()
	err := db.Create(&param).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *CartDomainCtx) Get(userID int64) ([]models.CartResp, error) {
	db := config.DbManager()
	result := []models.CartResp{}
	err := db.Table("cart").
		Select(`
            cart.id as cart_id, cart.user_id, cart.book_id, 
            books.title, books.thumbnail, books.writter_name, cart.created_at
        `).
		Joins("inner join books on cart.book_id = books.id").
		Where("cart.user_id = ?", userID).
		Scan(&result).Error
	if err != nil {
		return []models.CartResp{}, err
	}
	return result, nil
}

func (c *CartDomainCtx) RemoveFromCart(param models.Cart) error {
	db := config.DbManager()
	err := db.Where("user_id = ? AND book_id = ?", param.UserID, param.BookID).Delete(&models.Cart{}).Error
	if err != nil {
		return err
	}
	return nil
}
