package repositories

import (
	"nganterin-go/config"
	"nganterin-go/models"

	"gorm.io/gorm"
)

type CompRepository interface {
	RegisterUserCredential(data models.Users) (*string, error)
	GetUserDetailsByEmail(email string) (*models.Users, error)
	GetUserDetailsByID(id string) (*models.Users, error)

	RegisterEmailVerificationToken(data models.Users) (*string, error)
	VerifyUserEmail(token string) error

	RegisterPartner(data models.Partners) (*string, error)
	GetPartnerDetailsByID(id string) (*models.Partners, error)
	GetPartnerDetailsByEmail(email string) (*models.Partners, error)

	RegisterPartnerEmailVerificationToken(data models.Partners) (*string, error)
	VerifyPartnerEmail(token string) error

	RegisterHotel(data models.Hotels) (*string, error)
	GetAllHotels() ([]models.Hotels, error)
	SearchHotels(keyword string) ([]models.Hotels, error)
	GetHotelByID(id string)(*models.Hotels, error)
	GetHotelRoomByID(id uint) (*models.HotelRooms, error)

	FileUpload(data models.Files) (*models.Files, error)

	RegisterHotelOrder(data models.HotelOrders) error
	GetHotelOrderByID(id string) (*models.HotelOrders, error) 
	UpdateHotelOrderPaymentStatus(id string, status string) error
	GetAllHotelOrderByUserID(id string) ([]models.HotelOrders, error) 

	RegisterHotelReservation(data models.HotelReservations) error
}

type compRepository struct {
	DB *gorm.DB
}

func NewComponentRepository(DB *gorm.DB) *compRepository {
	db := config.InitDB()

	return &compRepository{
		DB: db,
	}
}
