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
	Insert(param models.User) (models.User, error)
	GetLoginUser(params *models.User) (*models.User, error)
	GetAll() ([]models.User, error)
	Update(param models.User) error
	GetUserName(param models.User) (models.User, error)
	Create(param models.User) (models.User, error)
	GetUserByEmail(param models.User) (models.User, error)
}
type UserDomainCtx struct{}

func (c *UserDomainCtx) GetWithRole(param models.GetUserParam) (models.UserData, error) {
	db := config.DbManager()
	var user models.UserData
	db = db.Table("users")
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

func (c *UserDomainCtx) Insert(param models.User) (models.User, error) {
	db := config.DbManager()
	err := db.Create(&param).Error
	if err != nil {
		return models.User{}, err
	}
	return param, nil
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
	if param.Otp != "" {
		update["otp"] = param.Otp
	}
	if !param.OtpExpiry.IsZero() {
		update["otp_expiry"] = param.OtpExpiry
	}
	if param.Name != "" {
		update["name"] = param.Name
	}
	if param.Role != "" {
		update["role"] = param.Role
	}
	if param.Language != "" {
		update["language"] = param.Language
	}
	if param.IsActive {
		update["is_active"] = param.IsActive
	}

	return db.Where("id = ?", userID).Updates(update).Error
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

func (c *UserDomainCtx) GetUserByEmail(param models.User) (models.User, error) {
	db := config.DbManager()
	result := models.User{}
	err := db.Where("email = ?", param.Email).First(&result).Error
	if err != nil {
		return models.User{}, err
	}

	return result, nil
}
