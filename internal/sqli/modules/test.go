package internals_sqli_modules

import (
	"core/http"
)

type String_test interface {
	GetCommentSymbols()
	GetQuoteSymbols()
	GetStringConcat()
}

type Numeric_test interface {
	GetCommentSymbols()
	GetQuoteSymbols()
	GetStringConcat()
}

type Test_url struct{
	Url core_http.Url
	String_test map[string]string
	Numeric_test map[string]string
}
