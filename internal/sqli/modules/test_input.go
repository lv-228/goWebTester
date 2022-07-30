package internals_sqli_modules

import(
	"core/http"
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
	JsonObject.AppendData(req.Url, answer.StatusCode, answer.Body.ToString(), answer.Ttf)
	return JsonObject
}

func (t *Test_input) RunPostTest(url string, params string){
	t.Url = &core_http.Url{
			url,
		}

	test_data := core_data_replace.NewReplaceS([]string{"test1", "test2", "test3"}, []string{"abc", "2", "3"}, []string{"4", "5", "6"})

	test_data.GenerateStrings()

	//log.Println(test_data.ResultStr)

	// headers := &core_http.HeaderData{}

	// headers.SetHeadersFromConfig()

	// request := &core_http.Req{
 //    	Req_type: "POST",
 //    	Headers_obj: headers,
 //    }

 //    var JsonObject SqliPostTestJsonObject
 //    var JsonObjects SqliPostTestJsonObject_array

 //    for key_url, _ := range req_params{
 //    	for _, elem_db_quote := range db_obj.GetQuoteSymbols(){
 //    		JsonObjects.Elem = append(JsonObjects.Elem, *t.do(elem_db_quote, &JsonObject, *request))
 //    	}
 //    }

 //    rawDataOut, err := json.MarshalIndent(&JsonObjects, "", "  ")
	// if err != nil{
	// 	log.Fatalln(err)
	// }

	// log.Println(string(rawDataOut))

	//core_data_json.SaveToJsonFile(rawDataOut, "./modules_data/" + JsonObject.GetFolderFromSave() + "/")
}

type SqliPostTestJsonObject struct {
	Url string
	Params string
	Status int
	Body string
	Ttf time.Duration
}

func (s *SqliPostTestJsonObject) ToByte() []byte{
	answer, err := json.Marshal(s)
	core_os.CheckErrValue(err, "Json marshal error! Object: response")
	return answer
}

func (s *SqliPostTestJsonObject) AppendData(url string, status int, body string, ttf time.Duration){
	s.Url = url
	//s.GetParam = param
	s.Status = status
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