package main

import (
	"net/http"
	//"log"
	//"io/ioutil"
	//"io"
	//"os"
	//"web_tester/http_funcs"
	"web_tester/server_ui"
	"web_tester/target"
	//"encoding/json"
)

var conf target.Config

//var request http_funcs.ReqData

var client = &http.Client{}

func main(){

	// data := target.GetConfig()

	// err := json.Unmarshal(data, &conf)
 //    if err != nil {
 //        log.Println("error:", err)
 //    }

	// if conf.Target["host"] == "host"{
	// 	log.Fatalln("Need correct host in config.json!")
	// }

	// request1 := &http_funcs.ReqData{
	// 	Req_type: "POST",
	// 	Url: conf.Target["host"] + conf.Target["auth_path"],
	// 	Data: []byte(`{
	// 		"email":"qweqe@qweqe",
	// 		"password":"13"
	// 	}`),
	// 	Headers: conf.Http_user_headers,
	// }

	// _, body := http_funcs.SendRequest(request1)

	// log.Println(body)

	server_ui.StartUiServer()



	//resp := http_funcs.GetRequest(Conf.Target["host"] + Conf.Target["login_page"])

	// answer,_ := io.ReadAll(resp.Body)

	// log.Println(string(answer))

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	//log.Println(string(body))
	//os.Exit(1)

	//forms := http_funcs.GetForms(string(body))
	//inputs := http_funcs.GetInputs(forms)

	// attributes := map[int][][]string{}

	// for i, elem := range inputs{
	// 	for _, input := range elem{
	// 		attributes[i] = append(attributes[i], http_funcs.GetHtmlAttributesByName(input))
	// 	}
	// }

 //  for i, elem := range http_funcs.GetRespHeaders(resp){
 //  	log.Println("Header:" + i, elem)
 //  }
}