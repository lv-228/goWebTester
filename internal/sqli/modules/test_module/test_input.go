package internals_sqli_modules_test

import(
	"core/http"
	"core/nosql"
	//"core/html/domobjs"
	"core/data/http"
	"core/data/replace"
	"time"
	"encoding/json"
	"core/os"
	"os"
	"log"
)

type Test_input struct{
	Url *core_http.Url
	Data_type core_data_http_types.Obj
}

func (t *Test_input) do(value string, JsonObject *SqliPostTestJsonObject, req core_http.Req) *SqliPostTestJsonObject{
	req.Url = t.Url.GetUrlWithoutParams()
	answer := req.SendAndGetResult(value)
	JsonObject.AppendData(req.Url, answer.StatusCode, value, answer.Body.ToString(), answer.Ttf)
	return JsonObject
}

func (t *Test_input) RunPostTest(url string, db_obj Test_interface){

	test_data_string := core_data_replace.NewReplaceS([]string{"str", "test2", "test3"}, db_obj.GetQuoteSymbols(), []string{"4", "5", "6"})
	test_data_numeric := core_data_replace.NewReplaceS([]string{"str", "test2", "test3"}, db_obj.GetNumericTesting(), []string{"4", "5", "6"})

	test_string_strings := test_data_string.GenerateStrings()
	test_numeric_strings := test_data_numeric.GenerateStrings()

	request := core_http.NewReq("GET", url, "url")

    couch_db := core_nosql.NewCouchDB("http://admin:123456@localhost:5984", "module_history")
	cuuid := core_nosql.NewCouchDBUuidResult([]byte(couch_db.GetUUIDsURL(1)))
    test_module_json := NewTestModuleJsonPut(cuuid.Uuids[0], db_obj.GetName(), "test_input", time.Now().UTC())
	test_module_json.Put(request, couch_db)

    var JsonObject SqliPostTestJsonObject
    var JsonObjects SqliPostTestJsonObject_array
	//log.Println(couch_db.GetByUUID(request, "a6b6047e16a7c3bfe9b2bc4c9e007749"))

    for _, string := range test_string_strings{
    	JsonObjects.Elem = append(JsonObjects.Elem, *t.do(string, &JsonObject, *request))
    }

    for _, string := range test_numeric_strings{
    	JsonObjects.Elem = append(JsonObjects.Elem, *t.do(string, &JsonObject, *request))
    }

    JsonObjects.Put(cuuid.Uuids[0], request, couch_db)
}

type SqliPostTestJsonObject struct {
	Url string
	Status int
	Body string
	Payload string
	Ttf time.Duration
}

func (s *SqliPostTestJsonObject) ToByte() []byte{
	answer, err := json.Marshal(s)
	core_os.CheckErrValue(err, "Json marshal error! Object: response")
	return answer
}

func (s *SqliPostTestJsonObject) AppendData(url string, status int, payload string, body string, ttf time.Duration){
	s.Url = url
	s.Status = status
	s.Payload = payload
	s.Body = body
	s.Ttf = ttf
}

func (s *SqliPostTestJsonObject) GetFolderFromSave() string{
	return "sqli/test_input"
}

type SqliPostTestJsonObject_array struct{
	Elem []SqliPostTestJsonObject
}

func (sq *SqliPostTestJsonObject_array) Put(id_module string, req *core_http.Req, c core_nosql.Couch_db){
	req.Data_type = "json"
	req.Req_type = "GET"
	couch_uuid := c.GetUUIDsURL(len(sq.Elem))
	c.Db = "module_result"
	uuid := core_nosql.NewCouchDBUuidResult([]byte(couch_uuid))
	for key, value := range sq.Elem{
		result_test_put := NewResultTestModuleJsonPut(&uuid.Uuids[key], &id_module, &value)
		log.Println(result_test_put.Put(req, c))
	}
}

func (sa *SqliPostTestJsonObject_array) GetDataFromFile(path string){
	jsonInFile, err1 := os.ReadFile(path)
	if err1 != nil{
		log.Fatalln(err1)
	}
	err2 := json.Unmarshal(jsonInFile, &sa)
	core_os.CheckErrValue(err2, "Ошибка дессериализации!")
}