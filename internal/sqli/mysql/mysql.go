package sqli_mysql

import (
	"core/sql"
)

type Mysql struct{
	Data core_sql.Sql_values
}

func (m *Mysql) GetCommentSymbols() []string {
	return m.Data.Comment
}

func (m *Mysql) GetQuoteSymbols() []string {
	return m.Data.Quotes
}

func (m *Mysql) GetConcatSymbols() []string {
	return m.Data.Concat
}

func (m *Mysql) GetSelectVersion() []string {
	return m.Data.SelectVersion
}

func (m *Mysql) GetDefaultDb() []string {
	return m.Data.Default_db
}


func NewMysql() Mysql{
	mysql := Mysql{}
	mysql.Data.Comment = {
		"#", 
		"--", 
		"/*", 
		"*/",
	}
	mysql.Data.SelectVersion = {"@@version"}
	mysql.Data.Quotes = {
		"\'",
		"\"",
	}
}