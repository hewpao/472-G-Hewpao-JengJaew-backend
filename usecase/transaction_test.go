package usecase_test

import (
	"errors"
	"testing"

	"github.com/hewpao/hewpao-backend/domain"
	"github.com/hewpao/hewpao-backend/usecase"
	"github.com/hewpao/hewpao-backend/usecase/test/mock_repos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransaction(t *testing.T) {
	mockTransactionRepo := new(mock_repos.MockTransactionRepository)
	transactionService := usecase.NewTransactionService(mockTransactionRepo)

	// Mock transaction data
	expectedTransaction := &domain.Transaction{
		UserID:   "1",
		Amount:   123.0,
		Currency: "THB",
	}

	t.Run("Success", func(t *testing.T) {
		mockTransactionRepo.ExpectedCalls = nil

		// Mock repository response
		mockTransactionRepo.On("Store", mock.Anything, mock.MatchedBy(func(t *domain.Transaction) bool {
			return t.UserID == expectedTransaction.UserID && t.Amount == expectedTransaction.Amount && t.Currency == expectedTransaction.Currency
		})).Return(nil)

		// Call the function
		transaction, err := transactionService.CreateTransaction(expectedTransaction.UserID, expectedTransaction.Amount, expectedTransaction.Currency)

		// Assertions
		assert.NoError(t, err)
		assert.NotNil(t, transaction)
		assert.Equal(t, expectedTransaction.UserID, transaction.UserID)
		assert.Equal(t, expectedTransaction.Amount, transaction.Amount)
		assert.Equal(t, expectedTransaction.Currency, transaction.Currency)

		// Verify expectations
		mockTransactionRepo.AssertExpectations(t)
	})

	t.Run("Error_StoreTransaction", func(t *testing.T) {
		mockTransactionRepo.ExpectedCalls = nil

		mockTransactionRepo.On("Store", mock.Anything, mock.Anything).Return(errors.New("database error"))

		transaction, err := transactionService.CreateTransaction(expectedTransaction.UserID, expectedTransaction.Amount, expectedTransaction.Currency)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, transaction)
		assert.EqualError(t, err, "database error")

		// Verify expectations
		mockTransactionRepo.AssertExpectations(t)
	})
}
