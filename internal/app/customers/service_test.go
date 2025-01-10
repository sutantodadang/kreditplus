package customers

import (
	"context"
	"errors"
	"testing"

	"kreditplus/internal/repositories"
	"kreditplus/internal/repositories/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCustomer(t *testing.T) {
	tests := []struct {
		name    string
		req     CreateCustomerRequest
		setup   func(*mocks.Querier)
		wantErr bool
	}{
		{
			name: "success create customer",
			req: CreateCustomerRequest{
				Fullname:             "John Doe",
				IdentificationNumber: "123456789",
				LegalName:            "John D.",
				PlaceOfBirth:         "City",
				DateOfBirth:          "1990-01-01",
				Salary:               "5000.00",
				PhotoKTP:             "ktp.jpg",
				PhotoSelfie:          "selfie.jpg",
				UserID:               uuid.New().String(),
				CustomerLimits: []CustomerLimit{
					{
						Tenor:       12,
						LimitAmount: 10000.00,
					},
				},
			},
			setup: func(q *mocks.Querier) {
				// Mock CreateCustomers
				q.On("CreateCustomers",
					mock.Anything,
					mock.MatchedBy(func(arg repositories.CreateCustomersParams) bool {
						return arg.FullName == "John Doe" && arg.IdentificationNumber == "123456789"
					}),
				).Return(nil).Once()

				// Mock CreateCustomersLimits
				q.On("CreateCustomersLimits",
					mock.Anything,
					mock.AnythingOfType("[]repositories.CreateCustomersLimitsParams"),
				).Return(int64(1), nil).Once()
			},
			wantErr: false,
		},
		{
			name: "error parse date of birth",
			req: CreateCustomerRequest{
				Fullname:             "Jane Doe",
				IdentificationNumber: "987654321",
				LegalName:            "Jane D.",
				PlaceOfBirth:         "Town",
				DateOfBirth:          "invalid-date",
				Salary:               "6000.00",
				PhotoKTP:             "ktp2.jpg",
				PhotoSelfie:          "selfie2.jpg",
				UserID:               uuid.New().String(),
				CustomerLimits: []CustomerLimit{
					{
						Tenor:       24,
						LimitAmount: 15000.00,
					},
				},
			},
			setup: func(q *mocks.Querier) {
				// No database calls expected
			},
			wantErr: true,
		},
		{
			name: "error create customer in DB",
			req: CreateCustomerRequest{
				Fullname:             "Mark Smith",
				IdentificationNumber: "555555555",
				LegalName:            "Mark S.",
				PlaceOfBirth:         "Village",
				DateOfBirth:          "1985-05-15",
				Salary:               "7000.00",
				PhotoKTP:             "ktp3.jpg",
				PhotoSelfie:          "selfie3.jpg",
				UserID:               uuid.New().String(),
				CustomerLimits: []CustomerLimit{
					{
						Tenor:       36,
						LimitAmount: 20000.00,
					},
				},
			},
			setup: func(q *mocks.Querier) {
				// Mock CreateCustomers to return error
				q.On("CreateCustomers",
					mock.Anything,
					mock.AnythingOfType("repositories.CreateCustomersParams"),
				).Return(errors.New("database error")).Once()
			},
			wantErr: true,
		},
		{
			name: "error create customer limits in DB",
			req: CreateCustomerRequest{
				Fullname:             "Lucy Heart",
				IdentificationNumber: "444444444",
				LegalName:            "Lucy H.",
				PlaceOfBirth:         "Metropolis",
				DateOfBirth:          "1992-08-20",
				Salary:               "8000.00",
				PhotoKTP:             "ktp4.jpg",
				PhotoSelfie:          "selfie4.jpg",
				UserID:               uuid.New().String(),
				CustomerLimits: []CustomerLimit{
					{
						Tenor:       48,
						LimitAmount: 25000.00,
					},
				},
			},
			setup: func(q *mocks.Querier) {
				// Mock CreateCustomers
				q.On("CreateCustomers",
					mock.Anything,
					mock.AnythingOfType("repositories.CreateCustomersParams"),
				).Return(nil).Once()

				// Mock CreateCustomersLimits to return error
				q.On("CreateCustomersLimits",
					mock.Anything,
					mock.AnythingOfType("[]repositories.CreateCustomersLimitsParams"),
				).Return(int64(0), errors.New("database limit error")).Once()
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

			service := NewCustomerService(mockQ)

			err := service.CreateCustomer(context.Background(), tt.req)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockQ.AssertExpectations(t)
		})
	}
}
