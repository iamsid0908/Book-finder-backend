package service

import (
	"core/domain"
	"core/models"
)

type BookSummaryService struct {
	BookSummaryDomain domain.BookSummaryDomain
}

func (c *BookSummaryService) Insert(param models.BookSummaryReqs) error {
	useParam := models.BookSummary{
		BookID:        param.BookID,
		Thumbnail:     param.Thumbnail,
		AuthorDetails: param.AuthorDetails,
		PublishedDate: param.PublishedDate,
		Summary:       param.Summary,
	}
	err := c.BookSummaryDomain.Insert(useParam)
	if err != nil {
		return err
	}
	return nil
}

func (c *BookSummaryService) GetBookDetails(param models.GetBookSummaryDetailsReqs) (models.GetBookSummaryDetailsResponse, error) {
	useParam := models.BookSummary{
		BookID: param.BookID,
	}
	data, err := c.BookSummaryDomain.GetBookDetails(useParam)
	if err != nil {
		return models.GetBookSummaryDetailsResponse{}, err
	}
	response := models.GetBookSummaryDetailsResponse{
		BookID:        param.BookID,
		Thumbnail:     data.Thumbnail,
		AuthorDetails: data.AuthorDetails,
		Summary:       data.Summary,
		PublishedDate: data.PublishedDate,
		Title:         data.Title,
		WritterName:   data.WritterName,
		Category:      data.Category,
	}
	return response, nil
}
