package core_http

import(
	"strings"
	"net/http"
	"core/data/replace"
	"web_tester/target"
	"encoding/json"
	"log"
)

type HeaderData struct{
	Headers map[string][]string
}

func (h *HeaderData) CreateFromTextArea(headers string) *core_data_replace.Replace{
	h.Headers = ParseTextareaHeaders(headers)
	elem, key := findHttpHeaderForReplace(h.Headers)
	headerReplace := &core_data_replace.Replace{}
	if elem != "" && key != ""{
		headerReplace.AppendString(elem)
	}
	return headerReplace
}

func (h *HeaderData) SetHeadersFromConfig(){
	var conf target.Config

	data := target.GetConfig()
	err1 := json.Unmarshal(data, &conf)
    if err1 != nil {
        log.Println("error:", err1)
    }

    h.Headers = make(map[string][]string, len(conf.Http_user_headers))

    for key_header, elem_header := range conf.Http_user_headers{
    	header_values := strings.Split(elem_header, ",")
    	h.Headers[key_header] = header_values
    }
}

func findHttpHeaderForReplace(headers map[string][]string) (string, string) {
	for key, elems := range headers{
		for _, elem := range elems{
			find_header := strings.Index(elem, Var_simbol_data)
			if find_header == -1{
				continue
			}
			return elem, key
			break
		}
	}
	return "", ""
}

func ParseTextareaHeaders(textarea_data string) map[string][]string{

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

	answer := map[string][]string{}
	
	for _, elem := range header_value{
		delimeter := strings.Index(elem, ":")
		header := elem[:delimeter]
		value := elem[delimeter+2:]
		answer[header] = append(answer[header], value)
	}

	return answer

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

func GetRespHeaders(resp *http.Response) map[string][]string{
	answer := map[string][]string{}
	for i, elem := range resp.Header{
		answer[i] = elem
	}
	return answer
}