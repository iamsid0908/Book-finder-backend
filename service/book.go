package service

import (
	"core/domain"
	"core/models"
	"fmt"
	"log"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

type BookService struct {
	BookDomain        domain.BookDomain
	BookSummaryDomain domain.BookSummaryDomain
}

func (b *BookService) Insert(param models.BookReqs) error {
	useParam := models.Books{
		Title:       param.Title,
		Thumbnail:   param.Thumbnail,
		WritterName: param.WritterName,
	}
	_, err := b.BookDomain.Insert(useParam)
	if err != nil {
		return err
	}
	return nil
}

func (b *BookService) GellAllBook(param models.SearchByInputParam) (models.BookRespData, error) {
	data, count, err := b.BookDomain.GetAll(param)
	if err != nil {
		return models.BookRespData{}, err
	}

	meta := models.MetaPagination{
		PageNumber:   param.Page,
		PageSize:     int64(len(data)),
		TotalPages:   int64(math.Ceil(float64(count) / float64(param.Limit))),
		TotalRecords: count,
	}

	resposnse := make([]models.BooksResp, len(data))
	for i, resp := range data {
		resposnse[i] = models.BooksResp{
			ID:          resp.ID,
			Title:       resp.Title,
			Thumbnail:   resp.Thumbnail,
			WritterName: resp.WritterName,
			Description: resp.Description,
			Cart:        resp.Cart,
			CreatedAt:   resp.CreatedAt,
			UpdatedAt:   resp.UpdatedAt,
		}
	}
	resp := models.BookRespData{
		Data: resposnse,
		Meta: meta,
	}
	return resp, nil
}

func (b *BookService) BulkInsert(param models.BulkInsertBookReqs) error {

	src, err := param.File.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Open the Excel file
	f, err := excelize.OpenReader(src)
	if err != nil {
		return err
	}

	// Read rows from the first sheet
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		return err
	}
	for i, row := range rows {
		if i == 0 {
			continue
		}
		fmt.Println(len(row))

		if len(row) < 2 {
			continue
		}
		useParam := models.Books{
			Title:       row[0],
			Thumbnail:   row[7],
			WritterName: row[1],
			Description: row[2],
			Category:    row[3],
		}
		bookID, err := b.BookDomain.Insert(useParam)
		if err != nil {
			return err
		}
		layout := "Monday, January 2, 2006"
		parsedDate, err := time.Parse(layout, row[5])
		if err != nil {
			log.Fatalf("Error parsing date: %v", err)
		}

		useParam1 := models.BookSummary{
			BookID:        bookID,
			Thumbnail:     row[7],
			AuthorDetails: row[8],
			Summary:       row[9],
			PublishedDate: parsedDate,
		}
		err = b.BookSummaryDomain.Insert(useParam1)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *BookService) Recommend(param models.RecommendReqs) ([]models.BooksResp, error) {
	books, err := b.BookDomain.GetCategory(param)
	if err != nil {
		return []models.BooksResp{}, err
	}
	tags := strings.Split(books.Category, ",")
	listOfBooks, err := b.BookDomain.Get()
	if err != nil {
		return []models.BooksResp{}, err
	}
	recommendations := RecommendBooks(tags, listOfBooks)

	return recommendations, nil
}

func RecommendBooks(preferences []string, books []models.Books) []models.BooksResp {
	var recommendations []models.BooksResp
	for _, book := range books {
		relevance := 0.0
		for _, genre := range preferences {
			// Split the category string into a slice of genres
			categories := strings.Split(book.Category, ",")
			for _, category := range categories {
				// Trim spaces to handle cases like "Fiction, Romance"
				if strings.TrimSpace(category) == genre {
					relevance += 1.0 // Increase relevance for matching genres
					break            // Avoid double counting if the genre is already matched
				}
			}
		}
		if relevance > 0 {
			recommendations = append(recommendations, models.BooksResp{
				ID:          book.ID,
				Title:       book.Title,
				Thumbnail:   book.Thumbnail,
				WritterName: book.WritterName,
				Description: book.Description,
				Relevance:   relevance,
			})
		}
	}

	// Sort recommendations by relevance
	sort.Slice(recommendations, func(i, j int) bool {
		return recommendations[i].Relevance > recommendations[j].Relevance
	})

	return recommendations
}
