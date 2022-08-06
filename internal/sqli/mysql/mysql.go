package sqli_mysql

import (
	"core/sql"
	"internal/sqli/modules/test"
)

type Mysql struct{
	Name string
	Data core_sql.Sql_values
}

func (m *Mysql) GetName() string{
	return "MySQL"
}

func (m *Mysql) GetCommentSymbols() []string {
	return m.Data.Comment
}

func (m *Mysql) GetQuoteSymbols() []string {
	return m.Data.Quotes
}

func (m *Mysql) GetStringConcat() []string {
	return m.Data.Concat
}

func (m *Mysql) GetSelectVersion() []string {
	return m.Data.SelectVersion
}

func (m *Mysql) GetDefaultDb() []string {
	return m.Data.Default_db
}

func (m *Mysql) GetNumericTesting() []string {
	return m.Data.Numeric_testing
}

func NewMysql() Mysql{
	mysql := Mysql{}
	mysql.Data.Comment = []string{
		"#", 
		"-- -",
		"`",
		"/*", 
		"*/",
	}
	mysql.Data.SelectVersion = []string{
		"@@version",
	}
	mysql.Data.Quotes = []string{
		"'",
		"\"",
	}
	mysql.Data.Numeric_testing = []string{
		"AND 1",
		"AND 0",
		"AND true",
		"AND false",
		"1-false",
		"1-true",
		"1*5",
	}

	mysql.Data.Default_funcs = map[string][]string{
		"current_user": []string{"user()",},
		"version": []string{"VERSION()","@@version",},
		"current_database": []string{"database()",},
		"system_user": []string{"system_user()",},
		"base_dir": []string{"@@basedir",},
	}

	return mysql
}

func NewMysqlInterface() internals_sqli_modules_test.Test_interface{
	mysql := NewMysql()
	var my_interface internals_sqli_modules_test.Test_interface
	my_interface = &mysql
	return my_interface
}