package repositories

import (
	"nganterin-go/exceptions"
	"nganterin-go/helpers"
	"nganterin-go/models/database"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompRepositoriesImpl struct {
}

func NewComponentRepository() CompRepositories {
	return &CompRepositoriesImpl{}
}

func (r *CompRepositoriesImpl) CreateCredentials(ctx *gin.Context, tx *gorm.DB, data database.Users) (*string, *exceptions.Exception) {
	data.ID = uuid.NewString()

	result := tx.Create(&data)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(result.Error)
	}

	return &data.ID, nil
}

func (r *CompRepositoriesImpl) FindByID(ctx *gin.Context, tx *gorm.DB, id string) (*database.Users, *exceptions.Exception) {
	var user_data database.Users
	result := tx.Where("id = ?", id).First(&user_data)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(result.Error)
	}

	return &user_data, nil
}

func (r *CompRepositoriesImpl) FindByEmail(ctx *gin.Context, tx *gorm.DB, email string) (*database.Users, *exceptions.Exception) {
	var user_data database.Users
	result := tx.Where("email = ?", email).First(&user_data)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(result.Error)
	}

	return &user_data, nil
}

func (r *CompRepositoriesImpl) VerifyEmail(ctx *gin.Context, tx *gorm.DB, token string) *exceptions.Exception {
	var token_data database.UserTokens
	result := tx.Where("token = ?", token).First(&token_data)
	if result.Error != nil {
		return exceptions.ParseGormError(result.Error)
	}

	user_model := database.Users{
		ID: token_data.UserID,
	}

	result = tx.Delete(&token_data)
	if result.Error != nil {
		return exceptions.ParseGormError(result.Error)
	}

	result = tx.Model(&user_model).Select("email_verified_at").Updates(map[string]interface{}{"email_verified_at": time.Now()})
	if result.Error != nil {
		return exceptions.ParseGormError(result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) CreateVerificationToken(ctx *gin.Context, tx *gorm.DB, id string) (*string, *exceptions.Exception) {
	delete_result := tx.Where("user_id = ? AND category = ?", id, "email_verification").Delete(&database.UserTokens{})
	if delete_result.Error != nil {
		return nil, exceptions.ParseGormError(delete_result.Error)
	}

	token, err := helpers.GenerateToken(32)
	if err != nil {
		return nil, err
	}

	token_data := database.UserTokens{
		UserID:    id,
		Token:     token,
		Category:  "email_verification",
		ExpiredAt: time.Now().Add(time.Hour * 24 * 3),
	}

	create_result := tx.Create(&token_data)
	if create_result.Error != nil {
		return nil, exceptions.ParseGormError(create_result.Error)
	}

	return &token, nil
}