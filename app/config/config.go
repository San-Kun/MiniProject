package config

import (
	"fmt"

	"github.com/San-Kun/MiniProject/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() *gorm.DB {

	config := Config{
		DB_Username: "pma",
		DB_Password: "root123",
		DB_Port:     "3306",
		DB_Host:     "127.0.0.1",
		DB_Name:     "mini_project_golang",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",

		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{
		//DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	InitialMigration()

	return DB
}

func InitialMigration() {
	DB.AutoMigrate(&domain.User{})
	DB.AutoMigrate(&domain.Bank{})
	DB.AutoMigrate(&domain.BankAccount{})
	DB.AutoMigrate(&domain.Event{})
	DB.AutoMigrate(&domain.EventTicket{})
	DB.AutoMigrate(&domain.Payment{})
	DB.AutoMigrate(&domain.PaymentType{})
}
