package repository_test

import (
	"testing"

	sqlMock "github.com/DATA-DOG/go-sqlmock"
	"github.com/San-Kun/MiniProject/domain"
	"github.com/San-Kun/MiniProject/user/repository"
	"github.com/stretchr/testify/assert"
)

func TestPaymentRepository_ReadAll(t *testing.T) {
	dbMock, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	mockPayment := []domain.Payment{
		domain.Payment{
			ID:            1,
			PaymentDate:   "20 Maret 2022",
			PaymentAmount: 45000,
		},
		domain.Payment{
			ID:            2,
			PaymentDate:   "20 Maret 2022",
			PaymentAmount: 50000,
		},
	}

	rows := sqlMock.NewRows([]string{"id", "payment_date", "payment_amount"}).
		AddRow(mockPayment[0].ID, mockPayment[0].PaymentDate, mockPayment[0].PaymentAmount).
		AddRow(mockPayment[1].ID, mockPayment[1].PaymentDate, mockPayment[1].PaymentAmount)

	query := "SELECT"

	mock.ExpectQuery(query).WillReturnRows(rows)
	mock.ExpectCommit()

	paymentRepository := repository.NewPaymentRepository(db)
	response, err := paymentRepository.ReadAll()

	assert.NotEmpty(t, response)
	assert.NoError(t, err)
	assert.Len(t, *response, 2)
}

func TestPaymentRepository_ReadByID(t *testing.T) {
	dbMock, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	payment := domain.Payment{
		ID:            1,
		PaymentDate:   "20 Maret 2022",
		PaymentAmount: 45000,
	}

	mock.ExpectQuery(
		"SELECT(.*)").
		WithArgs(1).
		WillReturnRows(
			sqlMock.NewRows([]string{"payment_date", "payment_amount", "id"}).
				AddRow(payment.PaymentDate, payment.PaymentAmount, payment.ID))

	paymentRepository := repository.NewPaymentRepository(db)
	response, err := paymentRepository.ReadByID(payment.ID)
	assert.NoError(t, err)
	assert.NotNil(t, *response)

	_, err = paymentRepository.ReadByID(2)
	assert.Error(t, err)

}

func TestPaymentRepository_Create(t *testing.T) {
	dbMock, mock, err := sqlMock.New(sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	payment := domain.Payment{
		ID:            1,
		PaymentDate:   "20 Maret 2022",
		PaymentAmount: 45000,
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `payments` (`payment_date`,`payment_amount`,`id`) VALUES (?,?,?)").
		WithArgs(payment.PaymentDate, payment.PaymentAmount, payment.ID).
		WillReturnResult(sqlMock.NewResult(1, 1))
	mock.ExpectCommit()

	paymentRepository := repository.NewPaymentRepository(db)
	response, err := paymentRepository.Create(&payment)

	assert.NoError(t, err)
	assert.NotNil(t, *response)

	payment2 := domain.Payment{
		ID:            1,
		PaymentDate:   "20 Maret 2022",
		PaymentAmount: 45000,
	}

	_, err = paymentRepository.Create(&payment2)

	assert.Error(t, err)
}
