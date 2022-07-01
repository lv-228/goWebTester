package http_funcs

import(
	"strings"
	"log"
	"bytes"
	"net/url"
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

func GetDataReader(request *Req, data string) *bytes.Buffer{
	data_reader := bytes.NewBuffer([]byte(data))
	switch request.Headers_obj.Headers["Content-Type"][0] {
		case "application/json":
			data_reader = bytes.NewBuffer([]byte(data))
		case "application/x-www-form-urlencoded":
			url_data := url.Values{}
			for i, elem := range getPurlFieldData(data){
				url_data.Set(i, elem)
			}
			data_reader = bytes.NewBuffer([]byte(url_data.Encode()))
		default:
			log.Fatalln("Error application format!")
	}
	return data_reader
}