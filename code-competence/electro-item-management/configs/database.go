package configs

import (
	"electro-item-management/domains"
	"electro-item-management/helpers"
	"errors"
	"fmt"
	"time"

	mysqlConfig "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() (*gorm.DB, error) {
	// connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	AppConfig.DBUsername,
	// 	AppConfig.DBPassword,
	// 	AppConfig.DBHost,
	// 	AppConfig.DBPort,
	// 	AppConfig.DBName,
	// )
	locationString := AppConfig.DBLocation
	location, err := time.LoadLocation(locationString)
	if err != nil {
		return nil, err
	}
	myConfig := mysqlConfig.Config{
		User:      AppConfig.DBUsername,
		Passwd:    AppConfig.DBPassword,
		Net:       "tcp",
		Addr:      fmt.Sprintf("%s:%s", AppConfig.DBHost, AppConfig.DBPort),
		DBName:    AppConfig.DBName,
		ParseTime: true,
		Loc:       location,
	}
	return gorm.Open(mysql.Open(myConfig.FormatDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		domains.User{},
		domains.Category{},
		domains.Item{},
	)
}

func MigrateAndSeedDB(db *gorm.DB) error {
	// drop table if exists
	err := db.Migrator().DropTable(domains.User{}, domains.Category{}, domains.Item{})
	if err != nil {
		return errors.New("fail to drop table")
	}

	// migrate table
	err = db.AutoMigrate(domains.User{}, domains.Category{}, domains.Item{})
	if err != nil {
		return errors.New("fail to migrate")
	}

	// seed the db
	var (
		hashedPassword, _ = helpers.HashPassword("123")
		adminUser         = domains.User{
			Name:     "admin",
			Email:    "admin@gmail.com",
			Password: hashedPassword,
		}
		categories = []domains.Category{{Name: "Analog"}, {Name: "Digital"}}
		items      = []domains.Item{
			{
				Name:        "Radio",
				CategoryID:  1,
				Description: "Radio analog klasik",
				Stock:       10,
				Price:       200000,
			},
			{
				Name:        "Speaker",
				CategoryID:  1,
				Description: "pengeras suara",
				Stock:       5,
				Price:       500000,
			},
			{
				Name:        "Television",
				CategoryID:  2,
				Description: "televisi digital",
				Stock:       3,
				Price:       1000000,
			},
			{
				Name:        "Laptop",
				CategoryID:  2,
				Description: "televisi digital",
				Stock:       6,
				Price:       2000000,
			},
		}
	)

	err = db.Create(&adminUser).Error
	if err != nil {
		return err
	}

	err = db.Create(&categories).Error
	if err != nil {
		return err
	}

	err = db.Create(&items).Error
	if err != nil {
		return err
	}

	return nil
}
