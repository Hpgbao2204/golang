package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

type Sql struct {
	Db       *sqlx.DB
	Host     string
	Port     int
	UserName string
	Password string
	DBName   string
}

func (s *Sql) Connect() {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		s.Host, s.Port, s.UserName, s.Password, s.DBName)
	s.Db = sqlx.MustConnect("postgres", datasource)

	if err := s.Db.Ping(); err != nil {
		log.Error(err)
	}

	fmt.Println("Connected to database")
}

func (s *Sql) Close() {
	_ = s.Db.Close()
}
