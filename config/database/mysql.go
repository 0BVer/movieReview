package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"movieReview/config/env"
)

func GetMysql() *gorm.DB {
	env.LoadEnv()

	USER := env.GetEnv("DB_USER", "root")
	PASS := env.GetEnv("MYSQL_ROOT_PASSWORD", "password")
	HOST := env.GetEnv("DB_HOST", "localhost")
	PORT := env.GetEnv("DB_PORT", "3306")
	DBNAME := env.GetEnv("DB_NAME", "movie_review")
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	fmt.Println(URL)
	db, err := gorm.Open(mysql.Open(URL))

	if err != nil {
		panic("Failed to connect to database!")
	} else {
		fmt.Println("Database connection established")
	}

	return db
}
