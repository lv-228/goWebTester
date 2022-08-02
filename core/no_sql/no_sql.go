package core_nosql

import(
	"core/http"
	"log"
	"encoding/json"
	"strconv"
)

type Couch_db struct{
	Url string
	GetUUIDUrl string
	Db string
}

func NewCouchDB(url string, db string) Couch_db{
	new := Couch_db{}
	new.Url = url
	new.Db = db
	new.GetUUIDUrl = "/_uuids?count="
	return new
}

func (c *Couch_db) GetUUIDs(req *core_http.Req, num int) string{
	req.Url = c.Url + c.GetUUIDUrl + strconv.Itoa(num)
	answer := req.SendAndGetResult("")
	return answer.Body.ToString()
}

func (c *Couch_db) GetByUUID(req *core_http.Req, id string) string{
	req.Url = c.Url + "/" + c.Db + "/" + id + "?confilcts=true"
	answer := req.SendAndGetResult("")
	return answer.Body.ToString()
}

func (c *Couch_db) GetModulesByType(req *core_http.Req, t string) string{
	req.Url = c.Url + "/module_history/_design/modules/_view/byType?key=\"" + t + "\""
	answer := req.SendAndGetResult("")
	return answer.Body.ToString()
}

func (c *Couch_db) GetResultsByModuleId(req *core_http.Req, id string) string{
	req.Url = c.Url + "/module_result/_design/result/_view/byIdModule?key=\"" + id + "\""
	answer := req.SendAndGetResult("")
	return answer.Body.ToString()
}

type Couch_db_default_fields struct{
	Id string `json:"_id"`
	Rev string `json:"_rev"`
}

type Couch_db_uuid_result struct{
	Uuids []string
}

func NewCouchDBUuidResult(b []byte) Couch_db_uuid_result{
	new := Couch_db_uuid_result{}
	err := json.Unmarshal(b, &new)
	if err != nil{
		log.Fatalln(err)
	}
	return new
}

func AddIdValueInUrlEncode(str string, id string) string{
	return str + "&_id=" + id
}