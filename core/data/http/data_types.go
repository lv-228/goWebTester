package core_data_http_types

import(
	"core/html/domobjs"
	"log"
)

type Obj struct{
	Type string
}

func (o *Obj) EncodeData(input core_html_domobjs.Input) string{
	if o.Type == "application/x-www-form-urlencoded"{
		return o.UrlEncodeData(input)
	}
}

func (o *Obj) UrlEncodeData(input core_html_domobjs.Input) string{
	return input.Main_data.Name + "=" + input.Main_data.Value + "&"
}