package server_ui

import(
	"strings"
	"web_tester/http_funcs"
	"strconv"
	"log"
	"errors"
)

type Replace struct{
	ReplacePosition int
	ReplaceString string
	Values []string
	Len int
	CurrentValuePos int
}

var Var_simbol = "¡"

func (r *Replace) FindReplacePosition(str string) int{
	find_header := strings.Index(str, Var_simbol)
	return find_header
}

func (r *Replace) Create(str string, Values []string){
	r.ReplacePosition = r.FindReplacePosition(str)
	r.ReplaceString = str
	r.Values = Values
	r.Len = len(r.Values)
}

func (r *Replace) Itteration(json bool) (string, error){
	if r.CurrentValuePos < r.Len{
		answer := http_funcs.ValueHeaderReplace(r.ReplaceString, r.Values[r.CurrentValuePos], Var_simbol)
		r.CurrentValuePos++
		return answer, nil
	}
	return "", errors.New("Out of range") 
}

func (r *Replace) CreateRange(){
	values, err := strconv.Atoi(r.Values[0])
	if err != nil{
		log.Fatalln("Ошибка диапазона значений")
	}

	for i := 1; i < values; i++{
		r.Values = append(r.Values, strconv.Itoa(i))
	}
	r.Values = r.Values[1:]
	r.Len = len(r.Values)
}