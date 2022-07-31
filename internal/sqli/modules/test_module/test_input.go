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
	t.Url = &core_http.Url{
			url,
		}

	test_data_string := core_data_replace.NewReplaceS([]string{"str", "test2", "test3"}, db_obj.GetQuoteSymbols(), []string{"4", "5", "6"})
	test_data_numeric := core_data_replace.NewReplaceS([]string{"str", "test2", "test3"}, db_obj.GetNumericTesting(), []string{"4", "5", "6"})

	test_string_strings := test_data_string.GenerateStrings()
	test_numeric_strings := test_data_numeric.GenerateStrings()


	headers := &core_http.HeaderData{}

	headers.SetHeadersFromConfig()

	request := &core_http.Req{
    	Req_type: "GET",
    	Headers_obj: headers,
    }

    couch_db := core_nosql.NewCouchDB("http://localhost:5984")
    couch_db.GetUUID(request)

    var JsonObject SqliPostTestJsonObject
    var JsonObjects SqliPostTestJsonObject_array

	module_json := NewLogObj(db_obj.GetName(), "test_module", time.Now().UTC())

	log.Println(module_json)

    for _, string := range test_string_strings{
    	JsonObjects.Elem = append(JsonObjects.Elem, *t.do(string, &JsonObject, *request))
    }

    for _, string := range test_numeric_strings{
    	JsonObjects.Elem = append(JsonObjects.Elem, *t.do(string, &JsonObject, *request))
    }

    rawDataOut, err := json.MarshalIndent(&JsonObjects, "", "  ")
	if err != nil{
		log.Fatalln(err)
	}

	log.Println(string(rawDataOut))

	//core_data_json.SaveToJsonFile(rawDataOut, "./modules_data/" + JsonObject.GetFolderFromSave() + "/")
}

type SqliPostTestJsonObject struct {
	Url string
	Params string
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

func (sa *SqliPostTestJsonObject_array) GetDataFromFile(path string){
	jsonInFile, err1 := os.ReadFile(path)
	if err1 != nil{
		log.Fatalln(err1)
	}
	err2 := json.Unmarshal(jsonInFile, &sa)
	core_os.CheckErrValue(err2, "Ошибка дессериализации!")
}