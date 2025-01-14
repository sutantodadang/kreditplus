package transactions

import (
	"context"
	"errors"
	"fmt"
	"kreditplus/internal/repositories"
	"kreditplus/internal/utils"
	"strconv"
	"sync"
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

	var wg sync.WaitGroup
	errChan := make(chan error, 1)

	wg.Add(1)
	go func(ctx context.Context, req CreateTransactionRequest) {
		defer wg.Done()
		data, err := s.repoDb.GetCustomerLimitById(ctx, req.CustomerId)
		if err != nil {
			log.Error().Err(err).Send()
			errChan <- err
			return
		}

		uuid7, err := uuid.NewV7()
		if err != nil {
			log.Error().Err(err).Send()
			errChan <- err
			return
		}

		custData, err := s.repoDb.GetCustomerTransactionByLimitIdAndCustomerId(ctx)
		if err != nil {
			log.Error().Err(err).Send()
			errChan <- err
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
			log.Error().Err(err).Send()
			errChan <- err
			return
		}

		limitAmount, err := strconv.ParseFloat(data[index].LimitAmount, 64)
		if err != nil {
			log.Error().Err(err).Send()
			errChan <- err
			return
		}

		if req.Otr > limitAmount {
			err = errors.New("limit amount is not enough")
			log.Error().Err(err).Send()
			errChan <- err
			return
		}

		historyLimit, err := s.repoDb.GetCustomerTransactionOtr(ctx, repositories.GetCustomerTransactionOtrParams{
			CustomerID:      req.CustomerId,
			CustomerLimitID: data[index].CustomerLimitID,
		},
		)

		if err != nil {
			log.Error().Err(err).Send()
			errChan <- err
			return
		}

		historyLimitFloat, err := strconv.ParseFloat(string(historyLimit.([]uint8)), 64)
		if err != nil {
			log.Error().Err(err).Send()
			errChan <- err
			return
		}

		if condition := historyLimitFloat + req.Otr; condition > limitAmount {
			err = errors.New("limit amount is not enough")
			log.Error().Err(err).Send()
			errChan <- err
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
			errChan <- err
			return
		}
	}(ctx, req)

	wg.Wait()

	select {
	case err = <-errChan:
		if err != nil {
			log.Error().Err(err).Send()
			return
		}
	case <-ctx.Done():
		err = ctx.Err()
		log.Error().Err(err).Send()
		return
	}

	close(errChan)

	return
}
