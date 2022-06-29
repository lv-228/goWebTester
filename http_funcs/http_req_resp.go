package http_funcs

import (
	"net/http"
	"net/url"
	"io"
	"log"
	"compress/gzip"
	"os"
	//"regexp"
	"strings"
	"bytes"
	"net/http/httptrace"
	"time"
	"crypto/tls"
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
	//heds := resp.Header.Get("content-type")
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

	//resp, err := client.Do(req)
	// if err != nil{
	// 	log.Fatalln(err)
	// }
	// defer resp.Body.Close()

	var start, connect, dns, tlsHandshake time.Time

    trace := &httptrace.ClientTrace{
        DNSStart: func(dsi httptrace.DNSStartInfo) { dns = time.Now() },
        DNSDone: func(ddi httptrace.DNSDoneInfo) {
            log.Printf("DNS Done: %v\n", time.Since(dns))
        },

        TLSHandshakeStart: func() { tlsHandshake = time.Now() },
        TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
            log.Printf("TLS Handshake: %v\n", time.Since(tlsHandshake))
        },

        ConnectStart: func(network, addr string) { connect = time.Now() },
        ConnectDone: func(network, addr string, err error) {
            log.Printf("Connect time: %v\n", time.Since(connect))
        },

        GotFirstResponseByte: func() {
            log.Printf("Time from start to first byte: %v\n", time.Since(start))
        },
    }

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

func RequestRepeater(SendRequest func(*ReqData, string) (map[string][]string, string), request *ReqData, data string, bf_journal_path string){
	bf_file, err := os.ReadFile(bf_journal_path)
	if err != nil{
		log.Fatalln(err)
	}

	words := strings.Fields(string(bf_file))

	if request.Headers["Content-Type"] == "application/json"{
		for _, elem := range words{
			headers, _ := SendRequest(request, valueJsonReplace(data, elem, Var_simbol_data))
			log.Println(elem, headers)
			//log.Println(GetHtmlTagByNameAndClass(body, "p", ))
		}
	}	else if request.Headers["Content-Type"] == "application/x-www-form-urlencoded"{
		for _, elem := range words{
			headers, body := SendRequest(request, valuePurlReplace(data, elem, Var_simbol_data))
			log.Println(elem, headers)
			log.Println(GetHtmlTagByNameAndClass(body, "p", "is-warning"))
		}
	}

}

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

func findHttpHeaderForReplace(headers map[string]string) map[string]string {
	answer := map[string]string{}
	for key, elem := range headers{
		find_header := strings.Index(elem, Var_simbol_http)
		if find_header == -1{
			continue
		}
		answer[key] = elem
		break
	}
	return answer
}
