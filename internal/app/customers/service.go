package customers

import (
	"context"
	"fmt"
	"kreditplus/internal/repositories"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type ICustomerService interface {
	CreateCustomer(ctx context.Context, req CreateCustomerRequest) (err error)
}

type CustomerService struct {
	repoDb repositories.Querier
}

func NewCustomerService(db repositories.Querier) *CustomerService {
	return &CustomerService{repoDb: db}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, req CreateCustomerRequest) (err error) {

	dateBirth, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	uuid7, err := uuid.NewV7()
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	err = s.repoDb.CreateCustomers(ctx, repositories.CreateCustomersParams{
		CustomerID:           uuid7.String(),
		FullName:             req.Fullname,
		IdentificationNumber: req.IdentificationNumber,
		LegalName:            req.LegalName,
		PlaceOfBirth:         req.PlaceOfBirth,
		DateOfBirth:          dateBirth,
		Salary:               req.Salary,
		PhotoKtp:             req.PhotoKTP,
		PhotoSelfie:          req.PhotoSelfie,
		UserID:               req.UserID,
	})
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	var customerLimits []repositories.CreateCustomersLimitsParams
	for _, v := range req.CustomerLimits {
		customerLimits = append(customerLimits, repositories.CreateCustomersLimitsParams{
			CustomerLimitID: uuid.New().String(),
			CustomerID:      uuid7.String(),
			Tenor:           int32(v.Tenor),
			LimitAmount:     fmt.Sprintf("%f", v.LimitAmount),
		})
	}
	_, err = s.repoDb.CreateCustomersLimits(ctx, customerLimits)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	return
}
