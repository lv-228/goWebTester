package core_http

import(
	"time"
	"encoding/json"
	"log"
)

type Resp struct {
	StatusCode int
	Body []byte
	Headers_obj *HeaderData
	Ttf time.Duration
}

func (r *Resp) Create(StatusCode int, body []byte, Headers map[string][]string){
	r.StatusCode = StatusCode
	r.Body = body
	r.Headers_obj = &HeaderData{
		Headers: Headers,
	}
}

func (r *Resp) ToByte() []byte {
	answer, err := json.Marshal(r)
	if err != nil{
		log.Fatalln("Json marshal error! Object: response")
	}
	return answer
}