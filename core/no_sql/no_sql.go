package core_nosql

import(
	"core/http"
	"log"
)

type couch_db struct{
	Url string
	GetUUIDUrl string
}

func NewCouchDB(url string) couch_db{
	new := couch_db{}
	new.Url = url
	new.GetUUIDUrl = "/_uuids?count=1"
	return new
}

func (c *couch_db) GetUUID(req *core_http.Req){
	req.Url = c.Url + c.GetUUIDUrl
	answer := req.SendAndGetResult("")
	log.Fatalln(answer.Body.ToString())
}