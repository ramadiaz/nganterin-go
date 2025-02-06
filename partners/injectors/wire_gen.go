// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injectors

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
	"nganterin-go/emails/services"
	"nganterin-go/partners/controllers"
	"nganterin-go/partners/repositories"
	services2 "nganterin-go/partners/services"
)

// Injectors from injector.go:

func InitializePartnerController(db *gorm.DB, validate *validator.Validate) controllers.CompControllers {
	compRepositories := repositories.NewComponentRepository()
	compServices := services.NewComponentServices()
	servicesCompServices := services2.NewComponentServices(compRepositories, compServices, db, validate)
	compControllers := controllers.NewCompController(servicesCompServices)
	return compControllers
}

// injector.go:

var partnerFeatureSet = wire.NewSet(repositories.NewComponentRepository, services.NewComponentServices, services2.NewComponentServices, controllers.NewCompController)
