package http_funcs

import(
	"time"
	"os"
	//"log"
	"strconv"
)

var date_ddMMyyyy_layout = "21.01.2001"

func SaveObjectInJsonFile(obj Save_to_json){
	jsonString := obj.ToJson()
	os.WriteFile(GetYearMonthDayNow(), jsonString, os.ModePerm)
}

func GetYearMonthDayNow() string{
	t := time.Now()
	return strconv.Itoa(t.Year()) + "-" + t.Month().String() + "-" + strconv.Itoa(t.Day())
}

// func GetObjectFromJsonFile(filename string){
	
// }