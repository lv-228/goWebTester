package main

import (
	"net/http"
	//"log"
	//"io/ioutil"
	//"io"
	//"os"
	"web_tester/web"
	"web_tester/target"
	//"internal/sqli/modules/test"
	//"core/http"
	//"core/data/json"
	//"core/data/http"
	//"core/html/domobjs"
	//"internal/sqli/mysql"
	//"encoding/json"
)

var conf target.Config

//var request http_funcs.ReqData

var client = &http.Client{}

func main(){

	//mysql_interface := sqli_mysql.NewMysqlInterface()

	// test_url_module := &internals_sqli_modules.Test_url{}
	// test_url_module.RunUrlTest("http://localhost/index_action.php?str=1&id=2", mysql_interface)

	// dob := core_html_domobjs.NewInput("test", "test")

	//test_input_module := internals_sqli_modules_test.Test_input{}

	//test_input_module.RunPostTest("http://localhost/index_action.php", mysql_interface)
// 
	// req := core_http.NewReq("POST", "http://localhost:3000/rest/user/login", "json")

	// test := req.SendAndGetResult(core_data_json.UrlToJSON("email=test@mail.ru&password=123"))

	// log.Println(core_data_json.UrlToJSON("email=test@mail.ru&password=123"))

	// log.Println(test.Body.ToString())

	//log.Fatalln("END")

	web_server.StartUiServer()

}