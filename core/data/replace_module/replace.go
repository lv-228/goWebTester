package core_data_replace

import(
	"strings"
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
	r.AppendString(str)
	r.AppendValues(Values, false)
}

func (r *Replace) AppendString(str string){
	r.ReplacePosition = r.FindReplacePosition(str)
	r.ReplaceString = str
}

func (r *Replace) AppendValues(Values []string, create_range bool){
	r.Values = Values
	r.Len = len(r.Values)
	if create_range == true {
		r.CreateRange()
	}
}

func (r *Replace) Itteration(json bool) (string, error){
	if r.CurrentValuePos < r.Len{
		answer := ValueHeaderReplace(r.ReplaceString, r.Values[r.CurrentValuePos], Var_simbol)
		r.CurrentValuePos++
		return answer, nil
	}
	return ValueHeaderReplace(r.ReplaceString, r.Values[r.CurrentValuePos-1], Var_simbol), errors.New("Out of range") 
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

func valueJsonReplace(str string, value string, replace_symbol string) string{
	find_place_in_str := strings.Index(str, replace_symbol)
	new_data_str := str[:find_place_in_str] + "\"" + value + "\"" + str[find_place_in_str+2:]
	return new_data_str
}

func valuePurlReplace(str string, value string, replace_symbol string) string{
	find_place_in_str := strings.Index(str, replace_symbol)
	new_data_str := str[:find_place_in_str] + value + str[find_place_in_str+2:]
	return new_data_str
}

func ValueHeaderReplace(str string, value string, replace_symbol string) string{
	return valuePurlReplace(str, value, replace_symbol)
}