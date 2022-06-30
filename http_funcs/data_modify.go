package http_funcs

import(
	"strings"
)

func valueJsonReplace(str string, value string, replace_symbol string) string{
	find_place_in_str := strings.Index(str, replace_symbol)
	new_data_str := str[:find_place_in_str] + "\"" + value + "\"" + str[find_place_in_str+2:]
	return new_data_str
}

func valuePurlReplace(str string, value string, replace_symbol string) string{
	find_place_in_str := strings.Index(str, replace_symbol)
	new_data_str := str[:find_place_in_str] + value + str[find_place_in_str+2:]
	return new_data_str
}

func ValueHeaderReplace(str string, value string, replace_symbol string) string{
	return valuePurlReplace(str, value, replace_symbol)
}