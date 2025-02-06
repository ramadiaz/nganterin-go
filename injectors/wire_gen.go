// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injectors

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
	controllers2 "nganterin-go/api/hotels/controllers"
	repositories2 "nganterin-go/api/hotels/repositories"
	services3 "nganterin-go/api/hotels/services"
	controllers3 "nganterin-go/api/orders/controllers"
	repositories3 "nganterin-go/api/orders/repositories"
	services4 "nganterin-go/api/orders/services"
	controllers5 "nganterin-go/api/reservations/controllers"
	repositories5 "nganterin-go/api/reservations/repositories"
	services6 "nganterin-go/api/reservations/services"
	controllers6 "nganterin-go/api/reviews/controllers"
	repositories6 "nganterin-go/api/reviews/repositories"
	services7 "nganterin-go/api/reviews/services"
	controllers4 "nganterin-go/api/storages/controllers"
	repositories4 "nganterin-go/api/storages/repositories"
	services5 "nganterin-go/api/storages/services"
	"nganterin-go/api/users/controllers"
	"nganterin-go/api/users/repositories"
	services2 "nganterin-go/api/users/services"
	"nganterin-go/emails/services"
)

// Injectors from injector.go:

func InitializeUserController(db *gorm.DB, validate *validator.Validate) controllers.CompControllers {
	compRepositories := repositories.NewComponentRepository()
	compServices := services.NewComponentServices()
	servicesCompServices := services2.NewComponentServices(compRepositories, compServices, db, validate)
	compControllers := controllers.NewCompController(servicesCompServices)
	return compControllers
}

func InitializeHotelController(db *gorm.DB, validate *validator.Validate) controllers2.CompControllers {
	compRepositories := repositories2.NewComponentRepository()
	compService := services3.NewComponentServices(compRepositories, db, validate)
	compControllers := controllers2.NewCompController(compService)
	return compControllers
}

func InitializeOrderController(db *gorm.DB, validate *validator.Validate) controllers3.CompControllers {
	compRepositories := repositories3.NewComponentRepository()
	repositoriesCompRepositories := repositories2.NewComponentRepository()
	compRepositories2 := repositories.NewComponentRepository()
	compServices := services4.NewComponentServices(compRepositories, repositoriesCompRepositories, compRepositories2, db)
	compControllers := controllers3.NewCompController(compServices)
	return compControllers
}

func InitializeStorageController(db *gorm.DB, validate *validator.Validate) controllers4.CompControllers {
	compRepositories := repositories4.NewComponentRepository()
	compServices := services5.NewComponentServices(compRepositories, db)
	compControllers := controllers4.NewCompController(compServices)
	return compControllers
}

func InitializeReservationController(db *gorm.DB, validate *validator.Validate) controllers5.CompControllers {
	compRepositories := repositories5.NewComponentRepository()
	repositoriesCompRepositories := repositories2.NewComponentRepository()
	compServices := services6.NewComponentServices(compRepositories, repositoriesCompRepositories, db)
	compControllers := controllers5.NewCompController(compServices)
	return compControllers
}

func InitializeReviewController(db *gorm.DB, validate *validator.Validate) controllers6.CompControllers {
	compRepositories := repositories6.NewComponentRepository()
	repositoriesCompRepositories := repositories3.NewComponentRepository()
	compRepositories2 := repositories5.NewComponentRepository()
	compServices := services7.NewComponentServices(compRepositories, repositoriesCompRepositories, compRepositories2, db, validate)
	compControllers := controllers6.NewCompController(compServices)
	return compControllers
}

// injector.go:

var userFeatureSet = wire.NewSet(repositories.NewComponentRepository, services.NewComponentServices, services2.NewComponentServices, controllers.NewCompController)

var hotelFeatureSet = wire.NewSet(repositories2.NewComponentRepository, services3.NewComponentServices, controllers2.NewCompController)

var orderFeatureSet = wire.NewSet(repositories.NewComponentRepository, repositories2.NewComponentRepository, repositories3.NewComponentRepository, services4.NewComponentServices, controllers3.NewCompController)

var storageFeatureSet = wire.NewSet(repositories4.NewComponentRepository, services5.NewComponentServices, controllers4.NewCompController)

var reservationFeatureSet = wire.NewSet(repositories5.NewComponentRepository, repositories2.NewComponentRepository, services6.NewComponentServices, controllers5.NewCompController)

var reviewFeatureSet = wire.NewSet(repositories3.NewComponentRepository, repositories5.NewComponentRepository, repositories6.NewComponentRepository, services7.NewComponentServices, controllers6.NewCompController)
