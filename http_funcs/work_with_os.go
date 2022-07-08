package http_funcs

import(
	"time"
	"os"
	"log"
	"strconv"
)

func SaveObjectInJsonFile(obj Save_to_json){
	jsonString := obj.ToByte()
	os.WriteFile(GetYearMonthDayNow(), jsonString, os.ModePerm)
}

func GetYearMonthDayNow() string{
	t := time.Now()
	return strconv.Itoa(t.Year()) + "-" + t.Month().String() + "-" + strconv.Itoa(t.Day())
}

func CheckErrValue(err error, message string){
	if err != nil{
		log.Printf(message + " Err: %s", err)
		os.Exit(1)
	}
}