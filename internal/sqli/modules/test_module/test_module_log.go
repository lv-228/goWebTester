package internals_sqli_modules_test

import (
	"time"
	"core/nosql"
	"core/http"
	"log"
	"encoding/json"
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

type test_module_json_put struct{
	Id string `json:"_id"`
	Db string
	Type_test string
	Start time.Time
}

func (t *test_module_json_put) Put(req *core_http.Req, c core_nosql.Couch_db) string{
	req.Req_type = "PUT"
	req.Url = c.Url + "/" + c.Db + "/" + t.Id
	json, err := json.Marshal(t)
	if err != nil{
		log.Fatalln(err)
	}
	answer := req.SendAndGetResult(string(json))
	return answer.Body.ToString()
}

func NewTestModuleJsonPut(id string, db string, tt string, start time.Time) test_module_json_put{
	new := test_module_json_put{}
	new.Id = id
	new.Db = db
	new.Type_test = tt
	new.Start = start
	return new
}