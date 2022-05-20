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

func TestPaymentTypeRepository_ReadAll(t *testing.T) {
	dbMock, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	mockPaymentType := []domain.PaymentType{
		domain.PaymentType{
			ID:          1,
			PayTypeName: "Transfer Bank BRI / BNI / BCA",
		},
		domain.PaymentType{
			ID:          2,
			PayTypeName: "Via Dana / GOPAY / OVO",
		},
	}

	rows := sqlMock.NewRows([]string{"id", "pay_type_name"}).
		AddRow(mockPaymentType[0].ID, mockPaymentType[0].PayTypeName).
		AddRow(mockPaymentType[1].ID, mockPaymentType[1].PayTypeName)

	query := "SELECT"

	mock.ExpectQuery(query).WillReturnRows(rows)
	mock.ExpectCommit()

	paymentTypeRepository := repository.NewPaymentTypeRepository(db)
	response, err := paymentTypeRepository.ReadAll()

	assert.NotEmpty(t, response)
	assert.NoError(t, err)
	assert.Len(t, *response, 2)
}

func TestPaymentTypeRepository_ReadByID(t *testing.T) {
	dbMock, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	paymentType := domain.PaymentType{
		ID:          1,
		PayTypeName: "Transfer Bank BRI / BNI / BCA",
	}

	mock.ExpectQuery(
		"SELECT(.*)").
		WithArgs(1).
		WillReturnRows(
			sqlMock.NewRows([]string{"pay_type_name", "id"}).
				AddRow(paymentType.PayTypeName, paymentType.ID))

	paymentTypeRepository := repository.NewPaymentTypeRepository(db)
	response, err := paymentTypeRepository.ReadByID(paymentType.ID)
	assert.NoError(t, err)
	assert.NotNil(t, *response)

	_, err = paymentTypeRepository.ReadByID(2)
	assert.Error(t, err)

}

func TestPaymentTypeRepository_Create(t *testing.T) {
	dbMock, mock, err := sqlMock.New(sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	paymentType := domain.PaymentType{
		ID:          1,
		PayTypeName: "Transfer Bank BRI / BNI / BCA",
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `paymentTypes` (`pay_type_name`,`id`) VALUES (?,?)").
		WithArgs(paymentType.PayTypeName, paymentType.ID).
		WillReturnResult(sqlMock.NewResult(1, 1))
	mock.ExpectCommit()

	paymentTypeRepository := repository.NewPaymentTypeRepository(db)
	response, err := paymentTypeRepository.Create(&paymentType)

	assert.NoError(t, err)
	assert.NotNil(t, *response)

	paymentType2 := domain.PaymentType{
		ID:          1,
		PayTypeName: "Transfer Bank BRI / BNI / BCA",
	}

	_, err = paymentTypeRepository.Create(&paymentType2)

	assert.Error(t, err)
}
