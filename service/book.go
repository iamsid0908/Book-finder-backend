package service

import (
	"core/domain"
	"core/models"
	"fmt"
)

type BookService struct {
	BookDomain domain.BookDomain
}

func (b *BookService) Insert(param models.BookReqs) error {
	useParam := models.Books{
		Title:       param.Title,
		Thumbnail:   param.Thumbnail,
		WritterName: param.WritterName,
	}
	err := b.BookDomain.Insert(useParam)
	if err != nil {
		return err
	}
	return nil
}

func (b *BookService) GellAllBook(param models.SearchByInputParam) ([]models.BooksResp, error) {
	fmt.Println(param.WritterName)
	data, err := b.BookDomain.GetAll(param)
	if err != nil {
		return []models.BooksResp{}, err
	}

	resposnse := make([]models.BooksResp, len(data))
	for i, resp := range data {
		resposnse[i] = models.BooksResp{
			ID:          resp.ID,
			Title:       resp.Title,
			Thumbnail:   resp.Thumbnail,
			WritterName: resp.WritterName,
			Cart:        resp.Cart,
			CreatedAt:   resp.CreatedAt,
			UpdatedAt:   resp.UpdatedAt,
		}
	}
	return resposnse, nil

}
