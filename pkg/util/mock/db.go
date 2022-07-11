package mock

import (
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnection() (*gorm.DB, sqlmock.Sqlmock) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	mysql := mysql.New(mysql.Config{
		Conn:       dbMock,
		DriverName: "mysql",
		SkipInitializeWithVersion:true,
	})

	db, err := gorm.Open(mysql, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db, mock

}
