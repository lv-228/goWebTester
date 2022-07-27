package internals_sqli_modules

import(
	"core/http"
	"core/html/domobjs"
)

type Test_input struct{
	Url *core_http.Url
	Data
}

func (t *Test_input) do(value string, JsonObject *SqliUrlTestJsonObject, req core_http.Req){
	req.Url = t.Url.GetUrlWithoutParams()
	answer = req.SendAndGetResult("qwe")
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

func (s *SqliPostTestJsonObject) GetFolderFromSave() string{
	return "sqli/test_url"
}

func (s *SqliPostTestJsonObject) AppendData(url string, param string, status int, body string, ttf time.Duration){
	s.Url = url
	s.GetParam = param
	s.Status = status
	s.Body = body
	s.Ttf = ttf
}