package orm

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBMock() (*gorm.DB, sqlmock.Sqlmock) {
	conn, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: conn}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db, mock
}

func AnyMockArg() sqlmock.Argument {
	return sqlmock.AnyArg()
}
