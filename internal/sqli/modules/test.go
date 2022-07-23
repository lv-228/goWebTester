package internals_sqli_modules

import (
	"core/http"
	"log"
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

type Test_url struct{
	Url *core_http.Url
	GetParams map[string]string
	String_test map[string]string
	Numeric_test map[string]string
}

func (t *Test_url) RunUrlTest(url string, db_obj Test_interface){
	t.Url  = &core_http.Url{
			url,
		}
	t.GetParams = t.Url.GetRequestParams()
	t.String_test = map[string]string{}
	t.Numeric_test = map[string]string{}
	log.Fatalln(t.GetParams)
}