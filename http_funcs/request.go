package http_funcs

import (
	"net/http"
	"net/url"
	"io"
	"log"
	"compress/gzip"
	//"os"
	//"regexp"
	//"strings"
	"bytes"
	"time"
	"net/http/httptrace"
)

type ReqData struct{
	Req_type, Url string
	Headers map[string]string
}

var Var_simbol_data = "¡"

var Var_simbol_http = "¶"

var client = &http.Client{}

func GetRequest(url string) *http.Response{
	resp, err := client.Get(url)
	if err != nil{
		log.Fatalln(err)
	}
	return resp
}

//Поменять потом на шаблон
func GetRespHeaders(resp *http.Response) map[string][]string{
	answer := map[string][]string{}
	for i, elem := range resp.Header{
		answer[i] = elem
	}
	return answer
}

func SendRequest(request *ReqData, data string) (map[string][]string, string){
	data_reader := bytes.NewBuffer([]byte(data))
	switch request.Headers["Content-Type"] {
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

	req, err := http.NewRequest(request.Req_type, request.Url, data_reader)

	if err != nil{
		log.Fatalln(err)
	}

	for i, elem := range request.Headers{
		req.Header.Set(i, elem)
	}

	trace := GetMetricsObject()

    req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
    start = time.Now()
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    log.Printf("Total time: %v\n", time.Since(start))

   	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
		case "gzip":
    		reader, err = gzip.NewReader(resp.Body)
    		if err != nil{
    			log.Fatalln(err)
    		}
    		defer reader.Close()
		default:
    		reader = resp.Body
	}

	bytes, err := io.ReadAll(reader)
	if err != nil{
		log.Fatalln(err)
	}

	log.Println(resp.StatusCode)

	return GetRespHeaders(resp), string(bytes)
}

// func RequestRepeater(SendRequest func(*ReqData, string) (map[string][]string, string), request *ReqData, data string, bf_journal_path string){

// 	words := strings.Fields(string(bf_file))

// 	if request.Headers["Content-Type"] == "application/json"{
// 		for _, elem := range words{
// 			headers, _ := SendRequest(request, valueJsonReplace(data, elem, Var_simbol_data))
// 			log.Println(elem, headers)
// 			//log.Println(GetHtmlTagByNameAndClass(body, "p", ))
// 		}
// 	}	else if request.Headers["Content-Type"] == "application/x-www-form-urlencoded"{
// 		for _, elem := range words{
// 			headers, body := SendRequest(request, valuePurlReplace(data, elem, Var_simbol_data))
// 			log.Println(elem, headers)
// 			log.Println(GetHtmlTagByNameAndClass(body, "p", "is-warning"))
// 		}
// 	}

// }