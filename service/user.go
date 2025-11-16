package service

import (
	"core/domain"
	"core/models"
)

type UserService struct {
	UserDomain domain.UserDomain
}

func (c *UserService) List() ([]models.ListOfUser, error) {
	data, err := c.UserDomain.GetAll()
	if err != nil {
		return []models.ListOfUser{}, err
	}
	response := make([]models.ListOfUser, len(data))
	for i, resp := range data {
		response[i] = models.ListOfUser{
			ID:        resp.ID,
			Email:     resp.Email,
			Name:      resp.Name,
			Role:      resp.Role,
			Language:  resp.Language,
			CreatedAt: resp.CreatedAt,
			UpdatedAt: resp.UpdatedAt,
		}
	}
	return response, nil
}

func FindRole(paraid int16) string {
	if paraid == 1 {
		return "Sales"
	} else if paraid == 2 {
		return "Accountant"
	} else if paraid == 3 {
		return "HR"
	} else if paraid == 4 {
		return "Administrator"
	} else if paraid == 5 {
		return "customer"
	}
	return ""
}
func (c *UserService) Update(param models.UpdateUserParam) error {
	useParam := models.User{
		ID:       param.UserID,
		Email:    param.Email,
		Name:     param.Name,
		Role:     param.Role,
		Language: param.Language,
	}

	err := c.UserDomain.Update(useParam)
	if err != nil {
		return err
	}
	return nil
}

func (c *UserService) GetUserName(userID int64) (models.UserDataResponse, error) {
	useParam := models.User{
		ID: userID,
	}
	data, err := c.UserDomain.GetUserName(useParam)
	if err != nil {
		return models.UserDataResponse{}, err
	}
	resp := models.UserDataResponse{
		ID:       data.ID,
		Email:    data.Email,
		Name:     data.Name,
		Role:     data.Role,
		Language: data.Language,
	}
	return resp, nil
}
