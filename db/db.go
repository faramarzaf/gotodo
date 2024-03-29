package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //to call init method of mysql library
	"time"
)

type Config struct {
	Username string
	Password string
	Port     int
	Host     string
	DBName   string
}

type MysqlDB struct {
	config Config
	db     *sql.DB
}

func New(config Config) *MysqlDB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	))

	if err != nil {
		fmt.Errorf("can't open mysql db: %v", err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MysqlDB{config: config, db: db}

}

func GetDbConfig() Config {
	return Config{
		Username: "fara",
		Password: "123",
		Port:     3307,
		Host:     "localhost",
		DBName:   "todo_db",
	}
}
