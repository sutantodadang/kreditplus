package users

import (
	"context"
	"kreditplus/internal/repositories"
	"kreditplus/internal/repositories/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// Dummy sql.Result for successful "CreateUser" calls
type mockSQLResult struct{}

func (r mockSQLResult) LastInsertId() (int64, error) { return 1, nil }
func (r mockSQLResult) RowsAffected() (int64, error) { return 1, nil }

func TestRegisterUser(t *testing.T) {
	tests := []struct {
		name    string
		req     RegisterUserRequest
		setup   func(*mocks.Querier)
		wantErr bool
	}{
		{
			name: "success register user",
			req: RegisterUserRequest{
				Email:    "test@example.com",
				Password: "password123",
			},
			setup: func(q *mocks.Querier) {
				q.On("CreateUser", mock.Anything, mock.Anything).
					Return(mockSQLResult{}, nil).
					Once()
			},
			wantErr: false,
		},
		{
			name: "error register user - db error",
			req: RegisterUserRequest{
				Email:    "test@example.com",
				Password: "password123",
			},
			setup: func(q *mocks.Querier) {
				q.On("CreateUser", mock.Anything, mock.Anything).
					Return(nil, assert.AnError).
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

			service := NewUserService(mockQ)

			err := service.RegisterUser(context.Background(), tt.req)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestLoginUser(t *testing.T) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	tests := []struct {
		name    string
		req     LoginUserRequest
		setup   func(*mocks.Querier)
		wantErr bool
	}{
		{
			name: "success login",
			req: LoginUserRequest{
				Email:    "test@example.com",
				Password: "password123",
			},
			setup: func(q *mocks.Querier) {
				q.On("GetUserByEmail", mock.Anything, "test@example.com").
					Return(repositories.GetUserByEmailRow{
						UserID:   "user-123",
						Password: string(hashedPassword),
					}, nil).
					Once()
			},
			wantErr: false,
		},
		{
			name: "error login - wrong password",
			req: LoginUserRequest{
				Email:    "test@example.com",
				Password: "wrongpassword",
			},
			setup: func(q *mocks.Querier) {
				q.On("GetUserByEmail", mock.Anything, "test@example.com").
					Return(repositories.GetUserByEmailRow{
						UserID:   "user-123",
						Password: string(hashedPassword),
					}, nil).
					Once()
			},
			wantErr: true,
		},
		{
			name: "error login - user not found",
			req: LoginUserRequest{
				Email:    "notfound@example.com",
				Password: "password123",
			},
			setup: func(q *mocks.Querier) {
				q.On("GetUserByEmail", mock.Anything, "notfound@example.com").
					Return(repositories.GetUserByEmailRow{}, assert.AnError).
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

			service := NewUserService(mockQ)

			token, err := service.LoginUser(context.Background(), tt.req)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
			}
		})
	}
}
