package repository_test

import (
	"testing"

	sqlMock "github.com/DATA-DOG/go-sqlmock"
	"github.com/San-Kun/MiniProject/domain"
	"github.com/San-Kun/MiniProject/user/repository"
	"github.com/stretchr/testify/assert"
)

func TestEventTicketRepository_ReadAll(t *testing.T) {
	dbMock, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	mockEventTicket := []domain.EventTicket{
		domain.EventTicket{
			ID:          1,
			EventID:     1,
			PaymentID:   1,
			PayTypeID:   1,
			PayTypeName: "Transfer Bank BRI / BNI / BCA",
		},
		domain.EventTicket{
			ID:          2,
			EventID:     2,
			PaymentID:   2,
			PayTypeID:   2,
			PayTypeName: "Via Dana / GOPAY / OVO",
		},
	}

	rows := sqlMock.NewRows([]string{"id", "event_id", "payment_id", "paytype_id", "pay_type_name"}).
		AddRow(mockEventTicket[0].ID, mockEventTicket[0].EventID, mockEventTicket[0].PaymentID, mockEventTicket[0].PayTypeID, mockEventTicket[0].PayTypeName).
		AddRow(mockEventTicket[1].ID, mockEventTicket[1].EventID, mockEventTicket[1].PaymentID, mockEventTicket[1].PayTypeID, mockEventTicket[1].PayTypeName)

	query := "SELECT"

	mock.ExpectQuery(query).WillReturnRows(rows)
	mock.ExpectCommit()

	eventTicketRepository := repository.NewEventTicketRepository(db)
	response, err := eventTicketRepository.ReadAll()

	assert.NotEmpty(t, response)
	assert.NoError(t, err)
	assert.Len(t, *response, 2)
}

func TestEventTicketRepository_ReadByID(t *testing.T) {
	dbMock, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	eventTicket := domain.EventTicket{
		ID:          1,
		EventID:     1,
		PaymentID:   1,
		PayTypeID:   1,
		PayTypeName: "Transfer Bank BRI / BNI / BCA",
	}

	mock.ExpectQuery(
		"SELECT(.*)").
		WithArgs(1).
		WillReturnRows(
			sqlMock.NewRows([]string{"event_id", "payment_id", "paytype_id", "pay_type_name", "id"}).
				AddRow(eventTicket.EventID, eventTicket.PaymentID, eventTicket.PayTypeID, eventTicket.PayTypeName, eventTicket.ID))

	eventTicketRepository := repository.NewEventTicketRepository(db)
	response, err := eventTicketRepository.ReadByID(eventTicket.ID)
	assert.NoError(t, err)
	assert.NotNil(t, *response)

	_, err = eventTicketRepository.ReadByID(2)
	assert.Error(t, err)

}

func TestEventTicketRepository_Create(t *testing.T) {
	dbMock, mock, err := sqlMock.New(sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	eventTicket := domain.EventTicket{
		ID:          1,
		EventID:     1,
		PaymentID:   1,
		PayTypeID:   1,
		PayTypeName: "Transfer Bank BRI / BNI / BCA",
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `eventTickets` ('event_id', 'payment_id', 'paytype_id', 'pay_type_name',`id`) VALUES (?,?,?,?,?)").
		WithArgs(eventTicket.EventID, eventTicket.PaymentID, eventTicket.PayTypeID, eventTicket.PayTypeName, eventTicket.ID).
		WillReturnResult(sqlMock.NewResult(1, 1))
	mock.ExpectCommit()

	eventTicketRepository := repository.NewEventTicketRepository(db)
	response, err := eventTicketRepository.Create(&eventTicket)

	assert.NoError(t, err)
	assert.NotNil(t, *response)

	eventTicket2 := domain.EventTicket{
		ID:          1,
		EventID:     1,
		PaymentID:   1,
		PayTypeID:   1,
		PayTypeName: "Transfer Bank BRI / BNI / BCA",
	}

	_, err = eventTicketRepository.Create(&eventTicket2)

	assert.Error(t, err)
}
