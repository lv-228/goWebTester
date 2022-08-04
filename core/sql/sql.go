package core_sql

import (
	"github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"fmt"
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

func (s *Sql_db_connect) Query(query string) ([][]string, error){
	rows, err1 := s.Connection.Query(query)
	if err1 != nil {
		return nil, fmt.Errorf("error: %s", err1)
	}
	defer rows.Close()

	columns, err2 := rows.Columns()
	if err2 != nil{
		return nil, fmt.Errorf("get columns failure")
	}

	len_columns := len(columns)

	test := make([]interface{}, len_columns)

	boof := make([]string, len_columns)

	answer := [][]string{}

	for i := 0; i < len_columns; i++{
		test[i] = &boof[i]
	}

	for rows.Next(){
		if err := rows.Scan(test[:]...); err != nil{
			return nil, fmt.Errorf("get result failure, err: %s", err)
		}
		answer = append(answer, boof)
	}

	return answer, nil
}

type Sql_values struct{
	Comment []string
	Quotes []string
	Concat []string
	SelectVersion []string
	Default_db []string
	Numeric_testing []string
	Default_funcs map[string][]string
}