package core_nosql

import(
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

func (c *Couch_db) GetUUIDsURL(num int) string{
	return c.Url + c.GetUUIDUrl + strconv.Itoa(num)
}

func (c *Couch_db) GetByUUIDURL(id string) string{
	return c.Url + "/" + c.Db + "/" + id + "?confilcts=true"
}

func (c *Couch_db) GetModulesByTypeURL(t string) string{
	return c.Url + "/module_history/_design/modules/_view/byType?key=\"" + t + "\""
}

func (c *Couch_db) GetResultsByModuleIdURL(id string) string{
	return c.Url + "/module_result/_design/result/_view/byIdModule?key=\"" + id + "\""
}

func (c *Couch_db) GetRequestResultsURL() string{
	return c.Url + "/http_history/_design/requests/_view/get_all"
}

func (c *Couch_db) GetResponseByRequestIdURL(id string) string{
	return c.Url + "/http_response/_design/responses/_view/get_by_request_id?key=\"" + id + "\""
}

type Couch_db_default_fields struct{
	Id string `json:"_id"`
	Rev string `json:"_rev"`
}

type Couch_db_put_result struct{
	Id string
	Rev string
	Ok bool
}

func NewCouchDbPutResult(result []byte) Couch_db_put_result{
	new := Couch_db_put_result{}

	err := json.Unmarshal(result, &new)
	if err != nil{
		log.Fatalln(err)
	}

	return new
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