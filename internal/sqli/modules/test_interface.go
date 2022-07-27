package internals_sqli_modules

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