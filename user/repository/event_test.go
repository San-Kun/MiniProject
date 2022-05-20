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

func TestEventRepository_ReadAll(t *testing.T) {
	dbMock, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	mockEvent := []domain.Event{
		domain.Event{
			ID:           1,
			EventName:    "Forum Diskusi",
			EventTanggal: "12 Maret 2022",
			Capacity:     "100",
		},
		domain.Event{
			ID:           2,
			EventName:    "Workshop Kebhinekaan",
			EventTanggal: "20 Maret 2022",
			Capacity:     "150",
		},
	}

	rows := sqlMock.NewRows([]string{"id", "event_name", "event_tanggal", "capacity"}).
		AddRow(mockEvent[0].ID, mockEvent[0].EventName, mockEvent[0].EventTanggal, mockEvent[0].Capacity).
		AddRow(mockEvent[1].ID, mockEvent[1].EventName, mockEvent[1].EventTanggal, mockEvent[1].Capacity)

	query := "SELECT"

	mock.ExpectQuery(query).WillReturnRows(rows)
	mock.ExpectCommit()

	eventRepository := repository.NewEventRepository(db)
	response, err := eventRepository.ReadAll()

	assert.NotEmpty(t, response)
	assert.NoError(t, err)
	assert.Len(t, *response, 2)
}

func TestEventRepository_ReadByID(t *testing.T) {
	dbMock, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	event := domain.Event{
		ID:           1,
		EventName:    "Forum Diskusi",
		EventTanggal: "12 Maret 2022",
		Capacity:     "100",
	}

	mock.ExpectQuery(
		"SELECT(.*)").
		WithArgs(1).
		WillReturnRows(
			sqlMock.NewRows([]string{"event_name", "event_tanggal", "capacity", "id"}).
				AddRow(event.EventName, event.EventTanggal, event.Capacity, event.ID))

	eventRepository := repository.NewEventRepository(db)
	response, err := eventRepository.ReadByID(event.ID)
	assert.NoError(t, err)
	assert.NotNil(t, *response)

	_, err = eventRepository.ReadByID(2)
	assert.Error(t, err)

}

func TestEventRepository_Create(t *testing.T) {
	dbMock, mock, err := sqlMock.New(sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	event := domain.Event{
		ID:           1,
		EventName:    "Forum Diskusi",
		EventTanggal: "12 Maret 2022",
		Capacity:     "100",
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `events` (`event_name`,`event_tanggal`,`capacity`,`id`) VALUES (?,?,?,?)").
		WithArgs(event.EventName, event.EventTanggal, event.Capacity, event.ID).
		WillReturnResult(sqlMock.NewResult(1, 1))
	mock.ExpectCommit()

	eventRepository := repository.NewEventRepository(db)
	response, err := eventRepository.Create(&event)

	assert.NoError(t, err)
	assert.NotNil(t, *response)

	event2 := domain.Event{
		ID:           1,
		EventName:    "Forum Diskusi",
		EventTanggal: "12 Maret 2022",
		Capacity:     "100",
	}

	_, err = eventRepository.Create(&event2)

	assert.Error(t, err)
}
