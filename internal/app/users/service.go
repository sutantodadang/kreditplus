package users

import (
	"context"
	"kreditplus/internal/repositories"
	"os"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	RegisterUser(ctx context.Context, req RegisterUserRequest) (err error)
	LoginUser(ctx context.Context, req LoginUserRequest) (token string, err error)
}

type UserService struct {
	repoDb repositories.Querier
}

func NewUserService(db repositories.Querier) *UserService {
	return &UserService{repoDb: db}
}

func (s *UserService) RegisterUser(ctx context.Context, req RegisterUserRequest) (err error) {

	hashPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	newId, err := uuid.NewV7()
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	_, err = s.repoDb.CreateUser(ctx, repositories.CreateUserParams{
		UserID:   newId.String(),
		Email:    req.Email,
		Password: string(hashPass),
	})

	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	return
}

func (s *UserService) LoginUser(ctx context.Context, req LoginUserRequest) (token string, err error) {
	userData, err := s.repoDb.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(req.Password))
	if err != nil {
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userData.UserID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	},
	)

	token, err = claims.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return
	}

	return
}
