package core_http

import (
	"strings"
	//"log"
)

type Url struct{
	Value string
}

func (u *Url) GetRequestParams() map[string]string{
	param_index := strings.Index(u.Value, "?")
	url_params_string := u.Value[param_index+1:]
	url_params := strings.Split(url_params_string, "&")
	
	answer := make(map[string]string, len(url_params))

	for _, elem := range url_params{
		equal_index := strings.Index(elem, "=")
		answer[elem[:equal_index]] = elem[equal_index+1:]
	}

	return answer
}

func (u *Url) GetUrlWithoutParams() string{
	return u.Value[:strings.Index(u.Value, "?")]
}