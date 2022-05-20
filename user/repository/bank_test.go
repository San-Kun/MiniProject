package repository_test

import (
	//"database/sql"
	"testing"

	sqlMock "github.com/DATA-DOG/go-sqlmock"
	"github.com/San-Kun/MiniProject/domain"
	"github.com/San-Kun/MiniProject/user/repository"
	"github.com/stretchr/testify/assert"
	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
)

/*
func SetupDBMock(dbMock *sql.DB) *gorm.DB {
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      dbMock,
		DSN:                       "sqlmock_db_0",
		SkipInitializeWithVersion: true,
	}), &gorm.Config{PrepareStmt: false})
	if err != nil {
		panic(err)
	}

	return gormDB
}
*/

func TestBankRepository_ReadAll(t *testing.T) {
	dbMock, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	mockBank := []domain.Bank{
		domain.Bank{
			ID:       1,
			BankName: "BANK BRI",
		},
		domain.Bank{
			ID:       2,
			BankName: "BANK BCA",
		},
	}

	rows := sqlMock.NewRows([]string{"id", "bank_name"}).
		AddRow(mockBank[0].ID, mockBank[0].BankName).
		AddRow(mockBank[1].ID, mockBank[1].BankName)

	query := "SELECT"

	mock.ExpectQuery(query).WillReturnRows(rows)
	mock.ExpectCommit()

	bankRepository := repository.NewBankRepository(db)
	response, err := bankRepository.ReadAll()

	assert.NotEmpty(t, response)
	assert.NoError(t, err)
	assert.Len(t, *response, 2)
}

func TestBankRepository_ReadByID(t *testing.T) {
	dbMock, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	bank := domain.Bank{
		ID:       1,
		BankName: "BANK BRI",
	}

	mock.ExpectQuery(
		"SELECT(.*)").
		WithArgs(1).
		WillReturnRows(
			sqlMock.NewRows([]string{"bank_name", "id"}).
				AddRow(bank.BankName, bank.ID))

	bankRepository := repository.NewBankRepository(db)
	response, err := bankRepository.ReadByID(bank.ID)
	assert.NoError(t, err)
	assert.NotNil(t, *response)

	_, err = bankRepository.ReadByID(2)
	assert.Error(t, err)

}

func TestBankRepository_Create(t *testing.T) {
	dbMock, mock, err := sqlMock.New(sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	bank := domain.Bank{
		ID:       1,
		BankName: "BANK BRI",
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `banks` (`bank_name`,`id`) VALUES (?,?)").
		WithArgs(bank.BankName, bank.ID).
		WillReturnResult(sqlMock.NewResult(1, 1))
	mock.ExpectCommit()

	bankRepository := repository.NewBankRepository(db)
	response, err := bankRepository.Create(&bank)

	assert.NoError(t, err)
	assert.NotNil(t, *response)

	bank2 := domain.Bank{
		ID:       1,
		BankName: "BANK BRI",
	}

	_, err = bankRepository.Create(&bank2)

	assert.Error(t, err)
}
