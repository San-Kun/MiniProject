package repository_test

import (
	"testing"

	sqlMock "github.com/DATA-DOG/go-sqlmock"
	"github.com/San-Kun/MiniProject/domain"
	"github.com/San-Kun/MiniProject/user/repository"
	"github.com/stretchr/testify/assert"
)

func TestBankAccountRepository_ReadAll(t *testing.T) {
	dbMock, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	mockBankAccount := []domain.BankAccount{
		domain.BankAccount{
			ID:            1,
			UserID:        1,
			BankID:        1,
			AccountNumber: "987654321",
			AccountName:   "IKHSAN ENDANG PRASETYA",
		},
		domain.BankAccount{
			ID:            2,
			UserID:        2,
			BankID:        2,
			AccountNumber: "1234567890",
			AccountName:   "ENDANG PRASETYA",
		},
	}

	rows := sqlMock.NewRows([]string{"id", "user_id", "bank_id", "account_number", "account_name"}).
		AddRow(mockBankAccount[0].ID, mockBankAccount[0].UserID, mockBankAccount[0].BankID, mockBankAccount[0].AccountNumber, mockBankAccount[0].AccountName).
		AddRow(mockBankAccount[1].ID, mockBankAccount[1].UserID, mockBankAccount[1].BankID, mockBankAccount[1].AccountNumber, mockBankAccount[1].AccountName)

	query := "SELECT"

	mock.ExpectQuery(query).WillReturnRows(rows)
	mock.ExpectCommit()

	bankAccountRepository := repository.NewBankAccountRepository(db)
	response, err := bankAccountRepository.ReadAll()

	assert.NotEmpty(t, response)
	assert.NoError(t, err)
	assert.Len(t, *response, 2)
}

func TestBankAccountRepository_ReadByID(t *testing.T) {
	dbMock, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	bankAccount := domain.BankAccount{
		ID:            1,
		UserID:        1,
		BankID:        1,
		AccountNumber: "987654321",
		AccountName:   "IKHSAN ENDANG PRASETYA",
	}

	mock.ExpectQuery(
		"SELECT(.*)").
		WithArgs(1).
		WillReturnRows(
			sqlMock.NewRows([]string{"user_id", "bank_id", "account_number", "account_name", "id"}).
				AddRow(bankAccount.UserID, bankAccount.BankID, bankAccount.AccountNumber, bankAccount.AccountName, bankAccount.ID))

	bankAccountRepository := repository.NewBankAccountRepository(db)
	response, err := bankAccountRepository.ReadByID(bankAccount.ID)
	assert.NoError(t, err)
	assert.NotNil(t, *response)

	_, err = bankAccountRepository.ReadByID(2)
	assert.Error(t, err)

}

func TestBankAccountRepository_Create(t *testing.T) {
	dbMock, mock, err := sqlMock.New(sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	bankAccount := domain.BankAccount{
		ID:            1,
		UserID:        1,
		BankID:        1,
		AccountNumber: "987654321",
		AccountName:   "IKHSAN ENDANG PRASETYA",
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `bank_accounts` ('user_id', 'bank_id', 'account_number','account_name',`id`) VALUES (?,?,?,?,?)").
		WithArgs(bankAccount.UserID, bankAccount.BankID, bankAccount.AccountNumber, bankAccount.AccountName, bankAccount.ID).
		WillReturnResult(sqlMock.NewResult(1, 1))
	mock.ExpectCommit()

	bankAccountRepository := repository.NewBankAccountRepository(db)
	response, err := bankAccountRepository.Create(&bankAccount)

	assert.NoError(t, err)
	assert.NotNil(t, *response)

	bankAccount2 := domain.BankAccount{
		ID:            1,
		UserID:        1,
		BankID:        1,
		AccountNumber: "987654321",
		AccountName:   "IKHSAN ENDANG PRASETYA",
	}

	_, err = bankAccountRepository.Create(&bankAccount2)

	assert.Error(t, err)
}
