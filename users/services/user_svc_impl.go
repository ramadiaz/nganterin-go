package services

import (
	"fmt"
	"net/http"
	emailServices "nganterin-go/emails/services"
	"nganterin-go/exceptions"
	"nganterin-go/helpers"
	"nganterin-go/models/database"
	"nganterin-go/models/dto"
	"nganterin-go/users/repositories"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo repositories.CompRepositories
	DB   *gorm.DB
}

func NewComponentServices(compRepositories repositories.CompRepositories, db *gorm.DB) CompServices {
	return &CompServicesImpl{
		repo: compRepositories,
		DB:   db,
	}
}

func (s *CompServicesImpl) CreateCredentials(ctx *gin.Context, data dto.User) *exceptions.Exception {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	hashed_password, err := helpers.HashPassword(data.Password)
	if err != nil {
		return exceptions.NewException(http.StatusInternalServerError, exceptions.ErrCredentialsHash)
	}

	user_data := database.Users{
		Name:            data.Name,
		Email:           data.Email,
		HashedPassword:  hashed_password,
		PhoneNumber:     data.PhoneNumber,
		Country:         data.Country,
		Province:        data.Province,
		City:            data.City,
		ZipCode:         data.ZipCode,
		CompleteAddress: data.CompleteAddress,
	}

	userID, err := s.repo.CreateCredentials(ctx, tx, user_data)
	if err != nil {
		return err
	}

	token, err := s.repo.CreateVerificationToken(ctx, tx, *userID)
	if err != nil {
		return err
	}

	go func() {
		verificationEmail := dto.EmailVerification{
			Email:   data.Email,
			Subject: "Nganterin - Verification Email",
			VerificationURL: os.Getenv("WEBCLIENT_BASE_URL") + "/auth/verify?token=" + *token,
		}

		err = emailServices.NewComponentServices().VerificationEmail(verificationEmail)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	return nil
}

func (s *CompServicesImpl) LoginCredentials(ctx *gin.Context, email string, password string) (*string, *exceptions.Exception) {
	data, err := s.repo.FindByEmail(ctx, s.DB, email)
	if err != nil {
		return nil, err
	}

	err = helpers.CheckPasswordHash(password, data.HashedPassword)
	if err != nil {
		return nil, err
	}

	if data.EmailVerifiedAt == nil {
		return nil, exceptions.NewException(http.StatusForbidden, exceptions.ErrEmailNotVerified)
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = data.ID
	claims["email"] = data.Email
	claims["name"] = data.Name
	claims["email_verified_at"] = data.EmailVerifiedAt
	claims["phone_number"] = data.PhoneNumber
	claims["country"] = data.Country
	claims["province"] = data.Province
	claims["city"] = data.City
	claims["zip_code"] = data.ZipCode
	claims["complete_address"] = data.CompleteAddress

	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	secretKey := []byte(secret)
	tokenString, signError := token.SignedString(secretKey)
	if signError != nil {
		return nil, exceptions.NewException(http.StatusInternalServerError, exceptions.ErrTokenGenerate)
	}

	return &tokenString, nil
}

func (s *CompServicesImpl) VerifyEmail(ctx *gin.Context, token string) *exceptions.Exception {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := s.repo.VerifyEmail(ctx, tx, token)
	if err != nil {
		return err
	}

	return nil
}

func (s *CompServicesImpl) LoginGoogleOAuth(ctx *gin.Context, email string, googleSUB string) (*string, *exceptions.Exception) {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	data, err := s.repo.FindByEmail(ctx, tx, email)
	if err != nil {
		return nil, err
	}

	err = helpers.CheckPasswordHash(googleSUB, data.HashedGoogleSUB)
	if err != nil {
		return nil, err
	}
	
	secret := os.Getenv("JWT_SECRET")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = data.ID
	claims["email"] = data.Email
	claims["name"] = data.Name
	claims["email_verified_at"] = data.EmailVerifiedAt
	claims["phone_number"] = data.PhoneNumber
	claims["country"] = data.Country
	claims["province"] = data.Province
	claims["city"] = data.City
	claims["zip_code"] = data.ZipCode
	claims["complete_address"] = data.CompleteAddress

	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	secretKey := []byte(secret)
	tokenString, signErr := token.SignedString(secretKey)
	if signErr != nil {
		return nil, exceptions.NewException(http.StatusInternalServerError, exceptions.ErrTokenGenerate)
	}

	return &tokenString, nil
}
