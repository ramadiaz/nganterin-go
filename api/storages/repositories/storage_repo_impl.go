package repositories

import (
	"nganterin-go/models"
	"nganterin-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompRepositoriesImpl struct {
}

func NewComponentRepository() CompRepositories {
	return &CompRepositoriesImpl{}
}

func (r *CompRepositoriesImpl) Create(ctx *gin.Context, tx *gorm.DB, data models.Files) (*models.Files, *exceptions.Exception) {
	data.ID = uuid.NewString()

	result := tx.Create(&data)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(result.Error)
	}

	return &data, nil
}
