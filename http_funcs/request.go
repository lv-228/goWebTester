package http_funcs

import (
	"net/http"
	"io"
	"log"
	"compress/gzip"
	"time"
	"net/http/httptrace"
)

type Req struct{
	Req_type, Url string
	Headers_obj *HeaderData
}

func (r *Req) Create(req_type string, url string, headers *HeaderData){
	r.Req_type = req_type
	r.Url = url
	r.Headers_obj = headers
}

func (r *Req) SendAndGetResult(data string) (map[string][]string, string){
	return r.sendRequest(data)
}

func (r *Req) sendRequest(data string) (map[string][]string, string){
	data_reader := GetDataReader(r, data)

	req, err := http.NewRequest(r.Req_type, r.Url, data_reader)

	if err != nil{
		log.Fatalln(err)
	}

	for i, elem := range r.Headers_obj.Headers{
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

var Var_simbol_data = "¡"

var client = &http.Client{}

func GetRequest(url string) *http.Response{
	resp, err := client.Get(url)
	if err != nil{
		log.Fatalln(err)
	}
	return resp
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
