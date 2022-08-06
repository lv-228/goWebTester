package core_http

import(
	"time"
	"encoding/json"
	"core/html"
	"core/nosql"
	"log"
)

type Resp struct {
	StatusCode int
	Body core_html.Body
	Headers_obj *HeaderData
	Ttf time.Duration
}

func (r *Resp) Create(StatusCode int, body []byte, Headers map[string][]string){
	r.StatusCode = StatusCode
	r.Body.Value = body
	r.Headers_obj = &HeaderData{
		Headers: Headers,
	}
}

func (r *Resp) ToByte() []byte {
	answer, err := json.Marshal(r)
	if err != nil{
		log.Fatalln("Json marshal error! Object: response")
	}
	return answer
}

type Resp_to_json_put struct{
	Id string `json: "_id"`
	RequestId string
	StatusCode int
	Body string
	Headers map[string]string
	Ttf time.Duration
}

func NewRespToJsonPut(response *Resp) Resp_to_json_put{
	new := Resp_to_json_put{}

	new.StatusCode = response.StatusCode
	new.Body = response.Body.ToString()

	headers := make(map[string]string, len(response.Headers_obj.Headers))

	for key, value := range response.Headers_obj.Headers{
		for _, elem := range value{
			headers[key] = elem
		}
	}

	new.Headers = headers

	new.Ttf = response.Ttf

	return new
}

func (r *Resp_to_json_put) Put(request_id string){
	r.RequestId = request_id

	request := NewReq("GET", "", "url")

	couch_db := core_nosql.NewCouchDB("http://admin:123456@localhost:5984", "http_response")
	request.Url = couch_db.GetUUIDsURL(1)

	r.Id = core_nosql.NewCouchDBUuidResult([]byte(request.SendAndGetResult("").Body.ToString())).Uuids[0]

	request.Req_type = "PUT"
	request.Data_type = "json"
	request.Url = couch_db.Url + "/" + couch_db.Db + "/" + r.Id
	json, err := json.Marshal(r)
	if err != nil{
		log.Fatalln(err)
	}

	answer := request.SendAndGetResult(string(json))

	log.Fatalln(answer.Body.ToString())
}