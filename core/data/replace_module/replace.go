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
		if json == false{
			answer := ValueHeaderReplace(r.ReplaceString, r.Values[r.CurrentValuePos], Var_simbol)
			r.CurrentValuePos++
			return answer, nil
		} else if json == true{
			answer := valueJsonReplace(r.ReplaceString, r.Values[r.CurrentValuePos], Var_simbol)
			r.CurrentValuePos++
			return answer, nil
		}
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

type ReplaceS struct{
	Keys []string
	Values []string
	DefaultValues []string
	Len int
	CurrentValuePos int
	ResultStr string
}

func NewReplaceS(keys []string, Values []string, DefaultValues []string) ReplaceS{
	if len(Values) != len(DefaultValues){
		errors.New("Количество значений для замены должно совпадать со значениями для замены по умолчанию!")
	}
	new_replaces := ReplaceS{}
	new_replaces.Keys = keys
	new_replaces.Values = Values
	new_replaces.DefaultValues = DefaultValues
	new_replaces.CurrentValuePos = 0
	return new_replaces
}

func (r *ReplaceS) FindReplacePositionS(str string){
	find_simbol := strings.Index(str, Var_simbol)
	var answer []int
	if find_simbol == -1{
		return
	}
	answer = append(answer, find_simbol)
	new_str := str[:find_simbol] + "?" + str[find_simbol+2:] 
	r.FindReplacePositionS(new_str)
	return
}

func (r *ReplaceS) ValuesStrings() map[int][]string{
	answer := make(map[int][]string, len(r.Keys))
	for i := 0; i < len(r.Keys); i++{
		for j := 0; j < len(r.Values); j++{
			answer[i] = append(answer[i], r.Keys[i] + "=" + r.Values[j])
		}
	}
	return answer
}

func (r *ReplaceS) DefaultStrings() []string{
	var answer []string
	for i := 0; i < len(r.Keys); i++{
		answer = append(answer, r.Keys[i] + "=" + r.DefaultValues[i])
	}
	return answer
}

func (r *ReplaceS) GenerateStrings(){
	default_str := r.DefaultStrings()
	values_str := r.ValuesStrings()
	boof := ""
	var answer []string
	for i := 0; i < len(default_str); i++{
		for j := 0; j < len(values_str[i]); j++{
			boof += values_str[i][j]
			for g := 0; g < len(default_str); g++{
				if g != i{
					boof += "&" + default_str[g]
				}
			}
			answer = append(answer, boof)
			boof = ""
		}
	}
	log.Println(answer)
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