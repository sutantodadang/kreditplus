// filepath: /C:/Users/sutan/OneDrive/Documents/Programming/interview/kreditplus/internal/app/transactions/service_test.go
package transactions

import (
	"context"
	"testing"

	"kreditplus/internal/repositories"
	"kreditplus/internal/repositories/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransaction(t *testing.T) {
	tests := []struct {
		name    string
		req     CreateTransactionRequest
		setup   func(*mocks.Querier)
		wantErr bool
	}{
		{
			name: "success create transaction",
			req: CreateTransactionRequest{
				CustomerId:  "customer-123",
				Otr:         1000.0,
				Tenor:       12,
				AdminFee:    100.0,
				Installment: 100.0,
				Interest:    50.0,
				AssetName:   "Test Asset",
			},
			setup: func(q *mocks.Querier) {
				q.On("GetCustomerLimitById", mock.Anything, "customer-123").
					Return([]repositories.GetCustomerLimitByIdRow{
						{CustomerLimitID: "limit-123", CustomerID: "customer-123", LimitAmount: "2000.0", Tenor: 12},
					}, nil).
					Once()
				q.On("GetCustomerTransactionByLimitIdAndCustomerId", mock.Anything).
					Return([]repositories.GetCustomerTransactionByLimitIdAndCustomerIdRow{}, nil).
					Once()
				q.On("GetCustomerTransactionOtr", mock.Anything, mock.Anything).
					Return([]uint8("0"), nil).
					Once()
				q.On("CreateTransaction", mock.Anything, mock.Anything).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
		{
			name: "error - tenor not found",
			req: CreateTransactionRequest{
				CustomerId: "customer-123",
				Tenor:      24,
				Otr:        1000.0,
			},
			setup: func(q *mocks.Querier) {
				q.On("GetCustomerLimitById", mock.Anything, "customer-123").
					Return([]repositories.GetCustomerLimitByIdRow{
						{CustomerLimitID: "limit-123", CustomerID: "customer-123", LimitAmount: "2000.0", Tenor: 12},
					}, nil).
					Once()
				q.On("GetCustomerTransactionByLimitIdAndCustomerId", mock.Anything).
					Return([]repositories.GetCustomerTransactionByLimitIdAndCustomerIdRow{}, nil).
					Once()
				// No expectation for GetCustomerTransactionOtr as tenor is not found
			},
			wantErr: true,
		},
		{
			name: "error - limit not enough",
			req: CreateTransactionRequest{
				CustomerId:  "customer-456",
				Otr:         2000.0,
				Tenor:       12,
				AdminFee:    100.0,
				Installment: 100.0,
				Interest:    50.0,
				AssetName:   "Test Asset",
			},
			setup: func(q *mocks.Querier) {
				q.On("GetCustomerLimitById", mock.Anything, "customer-456").
					Return([]repositories.GetCustomerLimitByIdRow{
						{CustomerLimitID: "limit-456", CustomerID: "customer-456", LimitAmount: "1500.0", Tenor: 12},
					}, nil).
					Once()
				q.On("GetCustomerTransactionByLimitIdAndCustomerId", mock.Anything).
					Return([]repositories.GetCustomerTransactionByLimitIdAndCustomerIdRow{}, nil).
					Once()
				q.On("GetCustomerTransactionOtr", mock.Anything, mock.Anything).
					Return([]uint8("0"), nil).
					Once()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockQ := mocks.NewQuerier(t)
			if tt.setup != nil {
				tt.setup(mockQ)
			}
			service := NewTransactionService(mockQ)
			err := service.CreateTransaction(context.Background(), tt.req)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockQ.AssertExpectations(t)
		})
	}
}
