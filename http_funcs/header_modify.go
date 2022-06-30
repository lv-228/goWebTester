package http_funcs

import(
	"strings"
)

func findHttpHeaderForReplace(headers map[string]string) map[string]string {
	answer := map[string]string{}
	for key, elem := range headers{
		find_header := strings.Index(elem, Var_simbol_data)
		if find_header == -1{
			continue
		}
		answer[key] = elem
		break
	}
	return answer
}

func ParseTextareaHeaders(textarea_data string) (map[string]string, map[string]string){

	header_value := map[int]string{}

	//get header: value
	for i := 0; ;i++{
		end_string := strings.Index(textarea_data, "~")
		if end_string == -1 {
			break
		}
		header_value[i] = textarea_data[:end_string]
		textarea_data = textarea_data[end_string+3:]
	}

	answer := map[string]string{}
	
	for _, elem := range header_value{
		delimeter := strings.Index(elem, ":")
		header := elem[:delimeter]
		value := elem[delimeter+2:]
		answer[header] = value
	}

	replaceHeader := findHttpHeaderForReplace(answer)

	return answer, replaceHeader

}

func getPurlFieldData(purl string) map[string]string{
	answer := map[string]string{}
	for i := 0; ;i++{
		index_delimeter := strings.Index(purl, "&")
		if index_delimeter == -1{
			return answer
		}
		key_value_left_str := purl[:index_delimeter]
		key_value_right_str := purl[index_delimeter+1:]
		index_eq := strings.Index(key_value_left_str, "=")
		key := key_value_left_str[:index_eq]
		value := key_value_left_str[index_eq+1:]
		answer[key] = value
		index_eq = strings.Index(key_value_right_str, "=")
		key = key_value_right_str[:index_eq]
		value = key_value_right_str[index_eq+1:]
		answer[key] = value
		purl = purl[index_delimeter+1:]
	}
	return answer
}