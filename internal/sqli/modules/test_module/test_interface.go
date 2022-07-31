package internals_sqli_modules_test

type String_test interface {
	GetCommentSymbols() []string
	GetQuoteSymbols() []string
	GetStringConcat() []string
}

type Numeric_test interface {
	GetNumericTesting() []string
}

type Test_interface interface{
	GetName() string
	String_test
	Numeric_test
}