package core_http

import (
	"net/http"
	"io"
	"log"
	"compress/gzip"
	"time"
	"net/http/httptrace"
	"encoding/json"
	"bytes"
	"net/url"
	"core/data/json"
	"core/nosql"
	//"crypto/tls"
)

type Req struct{
	Req_type, Url, Data_type string
	Headers_obj *HeaderData
}

func NewReq(req_type string, url string, data_type string) *Req{
	headers := &HeaderData{}
	headers.SetHeadersFromConfig()

	r := &Req{}

	r.Data_type = data_type
	r.Req_type = req_type
	r.Url = url
	r.Headers_obj = headers
	return r
}

func (r *Req) SendAndGetResult(data string) *Resp{
	return r.sendRequest(data)
}

func (r *Req) sendRequest(data string) *Resp{

	data_reader := GetDataReader(r, data)

	req, err := http.NewRequest(r.Req_type, r.Url, data_reader)

	if err != nil{
		log.Fatalln(err)
	}

	for i, elems := range r.Headers_obj.Headers{
		for _, elem := range elems{
			req.Header.Set(i, elem)
		}
	}

	if r.Data_type == "url"{
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else if r.Data_type == "json"{
		req.Header.Set("Content-Type", "application/json")
	}

	trace := GetMetricsObject()

    req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
    start = time.Now()
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    ttf := time.Since(start)

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

	my_response := &Resp{
		Ttf: ttf,
	}

	my_response.Create(resp.StatusCode, bytes, GetRespHeaders(resp))

	return my_response
}

func (r *Req) ToByte() []byte {
	answer, err := json.Marshal(r)
	if err != nil{
		log.Fatalln("Json marshal error! Object: request")
	}
	return answer
}

var start, connect, dns, tlsHandshake time.Time

func GetMetricsObject() *httptrace.ClientTrace{
	return &httptrace.ClientTrace{
        // DNSStart: func(dsi httptrace.DNSStartInfo) { dns = time.Now() },
        // DNSDone: func(ddi httptrace.DNSDoneInfo) {
        //     log.Printf("DNS Done: %v\n", time.Since(dns))
        // },

        // TLSHandshakeStart: func() { tlsHandshake = time.Now() },
        // TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
        //     log.Printf("TLS Handshake: %v\n", time.Since(tlsHandshake))
        // },

        // ConnectStart: func(network, addr string) { connect = time.Now() },
        // ConnectDone: func(network, addr string, err error) {
        //     log.Printf("Connect time: %v\n", time.Since(connect))
        // },

        GotFirstResponseByte: func() {
            //log.Printf("Time from start to first byte: %v\n", time.Since(start))
            time.Since(start)
        },
    }
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

func GetDataReader(request *Req, data string) *bytes.Buffer{
	data_reader := bytes.NewBuffer([]byte(data))

	switch request.Headers_obj.Headers["Content-Type"][0] {
		case "application/json":
			if json.Valid([]byte(data)) == false{
				data = core_data_json.UrlToJSON(data)
			}
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

type Req_to_json_put struct{
	Id string `json: "_id"`
	Req_type, Url, Data_type string
	Headers map[string]string
}

func NewReqToJson(req *Req) Req_to_json_put{

	new := Req_to_json_put{}

	new.Req_type = req.Req_type
	new.Url = req.Url
	new.Data_type = req.Headers_obj.Headers["Content-Type"][0]

	headers := make(map[string]string, len(req.Headers_obj.Headers))

	for key, value := range req.Headers_obj.Headers{
		for _, elem := range value{
			headers[key] = elem
		}
	}

	new.Headers = headers
	
	return new
}

func (r *Req_to_json_put) Put() string{
	request := NewReq("GET", "", "url")

	couch_db := core_nosql.NewCouchDB("http://admin:123456@localhost:5984", "http_history")
	request.Url = couch_db.GetUUIDsURL(1)

	r.Id = core_nosql.NewCouchDBUuidResult([]byte(request.SendAndGetResult("").Body.ToString())).Uuids[0]

	request.Req_type = "PUT"
	request.Data_type = "json"
	request.Url = couch_db.Url + "/" + couch_db.Db + "/" + r.Id
	json, err := json.Marshal(r)
	if err != nil{
		log.Fatalln(err)
	}

	answer := request.SendAndGetResult(string(json))

	res := core_nosql.NewCouchDbPutResult(answer.Body.Value)

	return res.Id
}