package service

import (
	"core/domain"
	"core/models"
)

type CartService struct {
	CartDomain domain.CartDomain
}

func (c *CartService) Insert(param models.InsertCartParam) error {
	useParam := models.Cart{
		UserID: param.UserID,
		BookID: param.BookID,
	}
	err := c.CartDomain.Insert(useParam)
	if err != nil {
		return err
	}
	return nil
}
func (c *CartService) GetCartByUserId(userID int64) ([]models.CartResp, error) {
	data, err := c.CartDomain.Get(userID)
	if err != nil {
		return []models.CartResp{}, err
	}
	response := make([]models.CartResp, len(data))
	for i, resp := range data {
		response[i] = models.CartResp{
			BookID:      resp.BookID,
			Title:       resp.Title,
			Thumbnail:   resp.Thumbnail,
			WritterName: resp.WritterName,
			Cart:        true,
		}
	}
	return response, nil
}

func (c *CartService) GetSizeofCart(userID int64) (int64, error) {
	data, err := c.CartDomain.Get(userID)
	if err != nil {
		return 0, err
	}
	response := int64(len(data))
	return response, nil
}

func (c *CartService) RemoveFromCart(param models.RemoveFromCartReqs) error {
	useParam := models.Cart{
		BookID: param.BookID,
		UserID: param.UserID,
	}
	err := c.CartDomain.RemoveFromCart(useParam)
	if err != nil {
		return err
	}
	return nil
}
