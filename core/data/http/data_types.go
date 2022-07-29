package core_data_http_types

import(
	"core/html/domobjs"
	"log"
	"strings"
)

type Obj struct{
	Type string
}

func (o *Obj) EncodeData(input core_html_domobjs.Input) string{
	if o.Type == "application/x-www-form-urlencoded"{
		return o.UrlEncodeData(input)
	}
	return ""
}

func (o *Obj) UrlEncodeData(input core_html_domobjs.Input) string{
	return input.Main_data.Name + "=" + input.Main_data.Value + "&"
}

func (o *Obj) TrimAmpersand(data string) string{
	return strings.TrimSuffix(data, "&")
}

func (o *Obj) GetKeyValueParam(data string) map[string]string{
	data = o.TrimAmpersand(data)
	params := strings.Split(data, "&")
	answer := make(map[string]string, len(params))
	log.Println(len(params))
	for _, elem := range params{
		log.Println(elem)
		equal_index := strings.Index(elem, "=")
		answer[elem[:equal_index]] = elem[equal_index+1:]
	}
	return answer
}