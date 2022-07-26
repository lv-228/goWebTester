package internals_sqli_modules

import (
	"core/http"
	"log"
	"time"
	"encoding/json"
	//"web_tester/target"
)

type String_test interface {
	GetCommentSymbols() []string
	GetQuoteSymbols() []string
	GetStringConcat() []string
}

type Numeric_test interface {
	GetNumericTesting() []string
}

type Test_interface interface{
	String_test
	Numeric_test
}

type SqliUrlTestJsonObject struct {
	Url string
	//GetParams map[string]string
	Status int
	Body string
	Ttf time.Duration
}

func (s *SqliUrlTestJsonObject) GetFolderFromSave() string{
	return "sqli/test_url"
}

func (s *SqliUrlTestJsonObject) AppendData(url string, status int, body string, ttf time.Duration){
	s.Url = url
	//s.GetParams = req_params
	s.Status = status
	s.Body = body
	s.Ttf = ttf
}

type Test_url struct{
	Url *core_http.Url
	String_test map[string]string
	Numeric_test map[string]string
}

func (t *Test_url) do(value string, JsonObject *SqliUrlTestJsonObject, req core_http.Req) *SqliUrlTestJsonObject{
	req.Url = t.Url.GetUrlWithoutParams() + value
	answer := req.SendAndGetResult("qwe")
	JsonObject.AppendData(req.Url, answer.StatusCode, answer.Body.ToString(), answer.Ttf)
	return JsonObject
}

func (t *Test_url) RunUrlTest(url string, db_obj Test_interface){
	t.Url = &core_http.Url{
			url,
		}

	req_params := t.Url.GetRequestParams()

	headers := &core_http.HeaderData{}

	headers.SetHeadersFromConfig()

	request := &core_http.Req{
    	Req_type: "GET",
    	Headers_obj: headers,
    }

    var JsonObject SqliUrlTestJsonObject
    var JsonObjects []SqliUrlTestJsonObject

    for key_url, _ := range req_params{
    	for _, elem_db_quote := range db_obj.GetQuoteSymbols(){
    		JsonObjects = append(JsonObjects, *t.do("?" + key_url + "=" + elem_db_quote, &JsonObject, *request))
			JsonObjects = append(JsonObjects, *t.do("?" + key_url + "=" + elem_db_quote + elem_db_quote, &JsonObject, *request))
    	}
    }

    rawDataOut, err := json.MarshalIndent(&JsonObjects, "", "  ")
	if err != nil{
		log.Fatalln(err)
	}

	log.Println(string(rawDataOut))
}