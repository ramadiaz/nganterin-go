package repositories

import (
	"nganterin-go/exceptions"
	"nganterin-go/models/database"
	"nganterin-go/models/dto"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositoriesImpl struct {
}

func NewComponentRepository() CompRepositories {
	return &CompRepositoriesImpl{}
}

func (r *CompRepositoriesImpl) Create(ctx *gin.Context, tx *gorm.DB, data database.HotelReservations) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) FindByID(ctx *gin.Context, tx *gorm.DB, id string) (*database.HotelReservations, *exceptions.Exception) {
	var data database.HotelReservations

	result := tx.First(&data, id)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(result.Error)
	}

	return &data, nil
}

func (r *CompRepositoriesImpl) FindByUserID(ctx *gin.Context, tx *gorm.DB, id string) ([]database.HotelOrders, *exceptions.Exception) {
	var data []database.HotelOrders

	result := tx.
		Joins("JOIN hotel_reservations ON hotel_reservations.hotel_orders_id = hotel_orders.id").
		Preload("HotelReservations").
		Preload("Hotel").
		Preload("Hotel.HotelsLocation").
		Preload("HotelRoom").
		Preload("HotelRoom.HotelRoomPhotos").
		Where("hotel_orders.user_id = ?", id).
		Order("hotel_orders.created_at DESC").
		Find(&data)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(result.Error)
	}

	return data, nil
}

func (r *CompRepositoriesImpl) FindByHotelID(ctx *gin.Context, tx *gorm.DB, hotelID string) ([]database.HotelOrders, *exceptions.Exception) {
	var data []database.HotelOrders

	result := tx.
		Joins("JOIN hotel_reservations ON hotel_reservations.hotel_orders_id = hotel_orders.id").
		Preload("HotelReservations").
		Preload("Hotel").
		Preload("HotelRoom").
		Preload("User").
		Where("hotel_orders.hotel_id = ?", hotelID).
		Order("hotel_orders.created_at DESC").
		Find(&data)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(result.Error)
	}

	return data, nil
}

func (r *CompRepositoriesImpl) FindByReservationKey(ctx *gin.Context, tx *gorm.DB, reservationKey string) (*database.HotelOrders, *exceptions.Exception) {
	var data database.HotelOrders
	result := tx.
		Preload("HotelReservations").
		Preload("User").
		Preload("Hotel").
		Preload("Hotel.HotelsLocation").
		Preload("HotelRoom").
		Preload("HotelRoom.HotelRoomPhotos").
		Joins("JOIN hotel_reservations ON hotel_reservations.hotel_orders_id = hotel_orders.id").
		Where("hotel_reservations.reservation_key = ?", reservationKey).
		Order("hotel_orders.created_at DESC").
		First(&data)

	if result.Error != nil {
		return nil, exceptions.ParseGormError(result.Error)
	}

	return &data, nil
}

func (r *CompRepositoriesImpl) CheckIn(ctx *gin.Context, tx *gorm.DB, reservationKey string) *exceptions.Exception {
	var data database.HotelReservations

	result := tx.Model(&data).Where("reservation_key = ?", reservationKey).Update("reservation_status",
		"CheckedIn")
	if result.Error != nil {
		return exceptions.ParseGormError(result.Error)
	}
	return nil
}

func (r *CompRepositoriesImpl) CheckOut(ctx *gin.Context, tx *gorm.DB, reservationKey string) *exceptions.Exception {
	var data database.HotelReservations

	result := tx.Model(&data).Where("reservation_key = ?", reservationKey).Update("reservation_status",
		"CheckedOut")
	if result.Error != nil {
		return exceptions.ParseGormError(result.Error)
	}
	return nil
}

func (r *CompRepositoriesImpl) Reviewed(ctx *gin.Context, tx *gorm.DB, orderID string) *exceptions.Exception {
	var data database.HotelReservations

	result := tx.Model(&data).Where("hotel_orders_id = ?", orderID).Update("reservation_status",
		"Reviewed")
	if result.Error != nil {
		return exceptions.ParseGormError(result.Error)
	}
	return nil
}

func (r *CompRepositoriesImpl) FindLast12MonthReservationCount(ctx *gin.Context, tx *gorm.DB, partnerID string) ([]dto.HotelMonthlyReservation, *exceptions.Exception) {
	now := time.Now()
	lastYear := now.AddDate(0, -11, 0)

	var data []dto.HotelMonthlyReservation

	result := tx.Raw(`
			WITH months AS (
				SELECT generate_series(
					date_trunc('month', ?::timestamp),
					date_trunc('month', ?::timestamp),
					'1 month'
				) AS month
			)
			SELECT 
				TO_CHAR(months.month, 'FMMonth YYYY') AS month_year,
				COALESCE(COUNT(hr.id), 0) AS reservation_count
			FROM months
			LEFT JOIN hotel_reservations hr 
				ON date_trunc('month', hr.created_at) = months.month
			LEFT JOIN hotel_orders ho 
				ON hr.hotel_orders_id = ho.id
			LEFT JOIN hotels h 
				ON ho.hotel_id = h.id AND h.partner_id = ?
			GROUP BY months.month
			ORDER BY months.month
		`, lastYear, now, partnerID).Scan(&data)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(result.Error)
	}

	return data, nil
}
