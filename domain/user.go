package domain

import (
	"core/config"
	"core/models"
	"errors"

	"gorm.io/gorm"
)

type UserDomain interface {
	GetWithRole(param models.GetUserParam) (models.UserData, error)
	Get(param models.GetUserParam) (models.User, error)
	Insert(param models.User) error
	GetLoginUser(params *models.User) (*models.User, error)
	GetAll() ([]models.User, error)
	Update(param models.User) error
	GetUserName(param models.User) (models.User, error)
	Create(param models.User) (models.User, error)
}
type UserDomainCtx struct{}

func (c *UserDomainCtx) GetWithRole(param models.GetUserParam) (models.UserData, error) {
	db := config.DbManager()
	var user models.UserData
	db = db.Table("users").Select("users.id, users.email, users.name, users.password,users.role_id, role.role as role, users.language, users.created_at, users.updated_at").
		Joins("left join role on users.role_id = role.id")
	if param.ID != 0 {
		db = db.Where("users.id = ?", param.ID)
	}
	if param.Email != "" {
		db = db.Where("users.email = ?", param.Email)
	}
	if err := db.First(&user).Error; err != nil {
		return models.UserData{}, err
	}
	return user, nil
}

func (c *UserDomainCtx) Get(param models.GetUserParam) (models.User, error) {
	db := config.DbManager()
	user := models.User{}
	if param.ID != 0 {
		db = db.Where("id = ?", param.ID)
	}

	if param.Email != "" {
		db = db.Where("email = ?", param.Email)
	}
	err := db.First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}

func (c *UserDomainCtx) Insert(param models.User) error {
	db := config.DbManager()
	err := db.Create(&param).Error
	if err != nil {
		return err
	}
	return nil
}
func (c *UserDomainCtx) Create(param models.User) (models.User, error) {
	db := config.DbManager()
	err := db.Create(&param).Error
	if err != nil {
		return models.User{}, err
	}
	return param, nil
}

func (c *UserDomainCtx) GetLoginUser(param *models.User) (*models.User, error) {
	db := config.DbManager()
	user := models.User{}

	if param.Email != "" {
		db = db.Where("email = ?", param.Email)
	}
	err := db.First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *UserDomainCtx) GetAll() ([]models.User, error) {
	db := config.DbManager()
	var users []models.User
	err := db.Find(&users).Error
	if err != nil {
		return []models.User{}, nil
	}
	return users, nil
}

func (c *UserDomainCtx) Update(param models.User) error {
	db := config.DbManager().Model(&models.User{})
	userID := param.ID
	update := map[string]interface{}{}
	if param.Email != "" {
		update["email"] = param.Email
	}
	if param.Name != "" {
		update["name"] = param.Name
	}
	if param.RoleId != 0 {
		update["role_id"] = param.RoleId
	}
	if param.Language != "" {
		update["language"] = param.Language
	}

	err := db.Where("id = ?", userID).Updates(update).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *UserDomainCtx) GetUserName(param models.User) (models.User, error) {
	db := config.DbManager()
	result := models.User{}
	err := db.Where("id = ?", param.ID).First(&result).Error
	if err != nil {
		return models.User{}, err
	}

	return result, nil
}
