package core_sql

import (
	"github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	//"fmt"
    //"os"
)

type Sql_db_connect struct{
	User string
	Passwd string
	Addr string
	DBName string
	Connection *sql.DB
}

func (s *Sql_db_connect) ConnectToDb(){
	cfg := mysql.Config{
		User: s.User,
		Passwd: s.Passwd,
		Net: "tcp",
		Addr: s.Addr,
		DBName: s.DBName,
	}

	var err error
	s.Connection, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil{
		log.Fatal(err)
	}

	pingErr := s.Connection.Ping()
	if pingErr != nil{
		log.Fatal(pingErr)
	}
	log.Println("Connected!")
}

func (s *Sql_db_connect) Query(query string){
	rows, err := s.Connection.Query(query)
}

func Test(){
	log.Fatalln("sql")
}