package dto

import "time"

type HotelOutputDTO struct {
	ID              string                  `json:"id"`
	PartnerID       string                  `json:"partner_id"`
	Name            string                  `json:"name"`
	Description     string                  `json:"description"`
	PricingStart    int64                   `json:"pricing_start"`
	HotelRooms      []HotelRoomOutput       `json:"hotel_rooms"`
	HotelsLocation  HotelsLocationOutput    `json:"hotels_location"`
	HotelPhotos     []HotelPhotoOutput      `json:"hotel_photos"`
	HotelFacilities []HotelFacilitiesOutput `json:"hotel_facilities"`
}

type HotelRoomOutput struct {
	ID             uint   `json:"id"`
	Type           string `json:"type"`
	MaxVisitor     int    `json:"max_visitor"`
	BedType        string `json:"bed_type"`
	RoomSize       int    `json:"room_size"`
	SmokingAllowed bool   `json:"smoking_allowed"`
	OvernightPrice int64  `json:"overnight_price"`
	TotalRoom      int    `json:"total_room"`
	TotalBooked    int    `json:"total_booked"`

	HotelRoomPhotoInput []HotelRoomPhotoInput `json:"hotel_room_photos" mapstructure:"hotel_room_photos"`
}

type HotelsLocationOutput struct {
	ID              uint   `json:"id"`
	Country         string `json:"country"`
	State           string `json:"state"`
	City            string `json:"city"`
	ZipCode         string `json:"zip_code"`
	CompleteAddress string `json:"complete_address"`
	Gmaps           string `json:"gmaps"`
}

type HotelPhotoOutput struct {
	ID  uint   `json:"id"`
	URL string `json:"url"`
}

type HotelRoomPhotoOutput struct {
	ID  uint   `json:"id"`
	URL string `json:"url"`
}

type HotelFacilitiesOutput struct {
	ID       uint   `json:"id"`
	Facility string `json:"facility"`
}

type HotelOrderOutput struct {
	ID          string `json:"id"`
	Token       string `json:"token"`
	RedirectURL string `json:"redirect_url"`
}

type HotelOrderDetailsOutput struct {
	ID               string    `json:"id"`
	UserID           string    `json:"user_id"`
	HotelID          string    `json:"hotel_id"`
	HotelRoomID      uint      `json:"hotel_room_id"`
	CheckInDate      time.Time `json:"check_in_date"`
	CheckOutDate     time.Time `json:"check_out_date"`
	TotalPrice       int64     `json:"total_price"`
	PaymentStatus    string    `json:"payment_status"`
	SnapToken        string    `json:"snap_token"`
	SpecialRequest   string    `json:"special_request"`
	IsForSomeoneElse bool      `json:"is_for_someone_else"`
	SomeoneName      string    `json:"someone_name"`
	SomeoneRegion    string    `json:"someone_region"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}