package sqli_mysql

import (
	"core/sql"
)

type Mysql struct{
	Data core_sql.Sql_values
}

func NewMysql() Mysql{
	mysql := Mysql{}
	mysql.Data.Comment := {
		"#", 
		"--", 
		"/*", 
		"*/",
	}
	mysql.Data.SelectVersion = "@@version"
}