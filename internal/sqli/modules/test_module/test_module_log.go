package internals_sqli_modules_test

import (
	"time"
)

type LogObj struct{
	Name string
	Type string
	TimeStart time.Time
}

func NewLogObj(name string, t string, tm time.Time) LogObj{
	log_obj := LogObj{}
	log_obj.Name = name
	log_obj.Type = t
	log_obj.TimeStart = tm
	return log_obj
}