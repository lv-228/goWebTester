package sqli

import (
	"log"
)

type Sqli_toolkit struct{
	Symbols map[string][]string
}

func (s *Sqli_toolkit) CreateSymbols(){
	s.Symbols = map[string][]string{
		"quote": {"'", "\""},
		"comment": {"#", "--"},
	}
}

// func (s *Sqli_toolkit) ErrorBasedInjection(){

// }

func Test(){
	test := &Sqli_toolkit{}
	test.CreateSymbols()
	log.Fatalln(test)
}

func FirstTry(){
	//Шаг 1: Кавычка
}

//Simple SQL
//SELECT * FROM user WHERE id = 10;