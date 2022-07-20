package core_os

import(
	"time"
	"os"
	"log"
	"strconv"
)

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