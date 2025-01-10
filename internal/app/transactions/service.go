package transactions

import (
	"context"
	"errors"
	"fmt"
	"kreditplus/internal/repositories"
	"kreditplus/internal/utils"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type ITransactionService interface {
	CreateTransaction(ctx context.Context, req CreateTransactionRequest) (err error)
}

type TransactionService struct {
	repoDb repositories.Querier
}

func NewTransactionService(db repositories.Querier) *TransactionService {
	return &TransactionService{repoDb: db}
}

func (s *TransactionService) CreateTransaction(ctx context.Context, req CreateTransactionRequest) (err error) {

	data, err := s.repoDb.GetCustomerLimitById(ctx, req.CustomerId)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	uuid7, err := uuid.NewV7()
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	custData, err := s.repoDb.GetCustomerTransactionByLimitIdAndCustomerId(ctx)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	dateContract := time.Now().Format("02012006")

	contractNumber := fmt.Sprintf("%04d", len(custData)+1) + dateContract

	index := -1
	for i, v := range data {
		if v.Tenor == int32(req.Tenor) {
			index = i
			break
		}
	}
	if index == -1 {
		err = errors.New("tenor not found")
		log.Error().Msg("Tenor not found")
		return
	}

	limitAmount, err := strconv.ParseFloat(data[index].LimitAmount, 64)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	if req.Otr > limitAmount {
		err = errors.New("limit amount is not enough")
		log.Error().Msg("Limit amount is not enough")
		return
	}

	historyLimit, err := s.repoDb.GetCustomerTransactionOtr(ctx, repositories.GetCustomerTransactionOtrParams{
		CustomerID:      req.CustomerId,
		CustomerLimitID: data[index].CustomerLimitID,
	},
	)

	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	historyLimitFloat, err := strconv.ParseFloat(string(historyLimit.([]uint8)), 64)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	if condition := historyLimitFloat + req.Otr; condition > limitAmount {
		err = errors.New("limit amount is not enough")
		log.Error().Msg("Limit amount is not enough")
		return

	}

	err = s.repoDb.CreateTransaction(ctx, repositories.CreateTransactionParams{
		TransactionID:     uuid7.String(),
		CustomerID:        req.CustomerId,
		CustomerLimitID:   data[index].CustomerLimitID,
		ContractNumber:    contractNumber,
		OtrAmount:         utils.ParseFloatToString(req.Otr),
		AdminFee:          utils.ParseFloatToString(req.AdminFee),
		InstallmentAmount: utils.ParseFloatToString(req.Installment),
		InterestAmount:    utils.ParseFloatToString(req.Interest),
		AssetName:         req.AssetName,
	})

	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	return
}
