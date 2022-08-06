package web_server

import(
	"os"
	"net/http"
	"log"
	"html/template"
	"regexp"
	"web_tester/target"
	//"web_tester/internal/sqli"
	"encoding/json"
	"core/sql"
	"core/nosql"
	"core/data/json"
	//"core/os"
	"core/http"
	"internal/sqli/modules/test"
	"internal/sqli/mysql"
	//"strings"
	//"strconv"
)

type Page struct {
	Title string
	Body []byte
	JsonList map[string][]core_data_json.HttpJsonObject
	TestUrlResult internals_sqli_modules_test.SqliUrlTestJsonObject_array
	TestResultBody internals_sqli_modules_test.SqliPostTestJsonObject
	TestModulesRows internals_sqli_modules_test.Test_module_json_rows
	Headers map[string]string
}

var conf target.Config

var html_folder = "./web/html/"

var validPath = regexp.MustCompile("^/(brute_module|settings|http_module|resp|sqli_ui|sqli_start_module|sqli_view)/([a-zA-Z0-9]+)$")

var tmpl_files = []string{
	html_folder + "templates/base.layout.tmpl",
	html_folder + "templates/header.tmpl",
	html_folder + "templates/footer.tmpl",
}

func loadPage(title string) (*Page, error){
	filename := html_folder + title + ".tmpl"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page, templates *template.Template){
	err := templates.ExecuteTemplate(w, tmpl + ".tmpl", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func bruteHandler(w http.ResponseWriter, r *http.Request, title string){
	p, err := loadPage("brute_module/" + title)
	if err != nil{
		http.Redirect(w, r, "/main/", http.StatusFound)
		return
	}

	files := []string{
		html_folder + "brute_module/brute.tmpl",
	}

	all_files := append(files, tmpl_files...)

	templates := template.Must(template.ParseFiles(all_files...))
	renderTemplate(w, "brute", p, templates)
}

func settingsHandler(w http.ResponseWriter, r *http.Request, title string){
	p, err := loadPage("settings/" + title)
	if err != nil{
		http.Redirect(w, r, "/main/", http.StatusFound)
		return
	}
	files := []string{
		html_folder + "settings/settings.tmpl",
	}

	all_files := append(files, tmpl_files...)

	templates := template.Must(template.ParseFiles(all_files...))
	renderTemplate(w, "settings", p, templates)
}

func http_moduleHandler(w http.ResponseWriter, r *http.Request, title string){
	p, err := loadPage("http_module/" + title)
	if err != nil{
		http.Redirect(w, r, "/main/", http.StatusFound)
		return
	}
	files := []string{
		html_folder + "http_module/index.tmpl",
		html_folder + "http_module/request.tmpl",
		html_folder + "http_module/list.tmpl",
	}

	all_files := append(files, tmpl_files...)
	templates := template.Must(template.ParseFiles(all_files...))

	data := target.GetConfig()
	err1 := json.Unmarshal(data, &conf)
    if err1 != nil {
        log.Println("error:", err)
    }

	p.Headers = conf.Http_user_headers

	// jsonFile := &core_data_json.JsonFile{}

	// jsonFile.GetJsonObject(core_os.GetYearMonthDayNow())

	// p.JsonList = jsonFile.ToMap()

	//log.Fatalln(test["HttpJsonObject_objs"][0].Request_obj)

	renderTemplate(w, "index", p, templates)
}

func sendRequestHandler(w http.ResponseWriter, r *http.Request, title string){

	data := target.GetConfig()

	err := json.Unmarshal(data, &conf)
    if err != nil {
        log.Println("error:", err)
    }

    headerData := &core_http.HeaderData{}
    headerReplace := headerData.CreateFromTextArea(r.FormValue("list_headers"))

    headerReplace.AppendValues([]string{r.FormValue("values")}, true)

    //log.Println(headerReplace)

    request := &core_http.Req{
    	Req_type: r.FormValue("method"),
    	Url: r.FormValue("url"),
    	Headers_obj: headerData,
    }

    response := request.SendAndGetResult(r.FormValue("data"))

    log.Fatalln(response.Body.ToString())

    // jsonHttpObject := &core_data_json.HttpJsonObject{
    // 	Request_obj: request,
    // 	Response_obj: response,
    // }

    // jsonFile := &core_data_json.JsonFile{}

    // jsonFile.SaveJsonFile(jsonHttpObject)

    // log.Println(jsonFile)

	// bfReplace := &Replace{}

	// //if header_key != ""{
	// 	// step, err := strconv.Atoi(r.FormValue("step"))
	// 	// if err != nil{
	// 	// 	log.Fatalln("Ошибка шага")
	// 	// }

	// 	//log.Println(http_funcs.GetHtmlTagByNameAndClass(body, "p", "is-warning"))
	// //}

	// if(r.FormValue("bf_journal") != ""){
	// 	bf_file, err := os.ReadFile(r.FormValue("bf_journal"))
	// 	if err != nil{
	// 		log.Fatalln(err)
	// 	}

	// 	words := strings.Fields(string(bf_file))

	// 	bfReplace.Create(r.FormValue("data"), words)

	// // 	if request.Headers["Content-Type"] == "application/json"{
	// // 	for _, elem := range words{
	// // 		headers, _ := SendRequest(request, valueJsonReplace(data, elem, Var_simbol_data))
	// // 		log.Println(elem, headers)
	// // 		//log.Println(GetHtmlTagByNameAndClass(body, "p", ))
	// // 	}
	// // }	else if request.Headers["Content-Type"] == "application/x-www-form-urlencoded"{
	// // 	for _, elem := range words{
	// // 		headers, body := SendRequest(request, valuePurlReplace(data, elem, Var_simbol_data))
	// // 		log.Println(elem, headers)
	// // 		log.Println(GetHtmlTagByNameAndClass(body, "p", "is-warning"))
	// // 	}
	// // }
		
	// }

	// for ;; {
	// 	str_bf, err_bf := bfReplace.Itteration(false)
	// 	str_head, err_head := headerReplace.Itteration(false)

	// 	if err_bf != nil && err_head != nil{
	// 		log.Fatalln("DATA END")
	// 	}

	// 	headers_request[header_key] = str_head
	// 	request1.Headers = headers_request
	// 	headers, _ := http_funcs.SendRequest(request1, str_bf)
	// 	log.Println(headers_request, headers, str_bf)
	// }

	// _, body_response := http_funcs.SendRequest(request1, r.FormValue("data"))

	// p := &Page{Title: title, Body: []byte(body_response)}

	// files := []string{
	// 	html_folder + "req_resp/response.tmpl",
	// }

	// all_files := append(files, tmpl_files...)
	// templates := template.Must(template.ParseFiles(all_files...))

	// renderTemplate(w, "response", p, templates)

}

func sqliUiHandler(w http.ResponseWriter, r *http.Request, title string){
	if title == "testUrl"{
		sqliUiTestUrl(title, w, r)
	} else if title == "byerror"{
		sqliUiByError(title, w, r)
	} else if title == "viewResultById"{
		sqliViewResultById(title, w, r)
	} else if title == "viewModulesByType"{
		sqliViewModuleResult(title, w, r)
	}else if title == "viewResultsByIdModule"{
		sqliViewResultsByModuleId(title, w, r)
	}
}

func sqliViewResultsByModuleId(title string, w http.ResponseWriter, r *http.Request){
	req := core_http.NewReq("GET", "", "json")
	couch_db := core_nosql.NewCouchDB("http://admin:123456@localhost:5984", "module_result")
	module_results := couch_db.GetResultsByModuleId(req, "b40caeaecf7f199835cffd88cd098d24")

	p, err1 := loadPage("sqli_view/" + title)
	if err1 != nil{
		http.Redirect(w, r, "/main/", http.StatusFound)
		return
	}

	p.TestModulesRows = internals_sqli_modules_test.NewTestModuleJsonRowsFromByte([]byte(module_results))

	files := []string{
		html_folder + "sqli_view/viewResultsByIdModule.tmpl",
	}

	all_files := append(files, tmpl_files...)

	templates := template.Must(template.ParseFiles(all_files...))
	renderTemplate(w, "viewResultsByIdModule", p, templates)
}

func sqliViewModuleResult(title string, w http.ResponseWriter, r *http.Request){
	req := core_http.NewReq("GET", "", "json")
	couch_db := core_nosql.NewCouchDB("http://admin:123456@localhost:5984", "module_history")
	params := r.URL.Query()
	module_rows := couch_db.GetModulesByType(req, params["type"][0])

	p, err1 := loadPage("sqli_view/" + title)
	if err1 != nil{
		http.Redirect(w, r, "/main/", http.StatusFound)
		return
	}

	p.TestModulesRows = internals_sqli_modules_test.NewTestModuleJsonRowsFromByte([]byte(module_rows))

	files := []string{
		html_folder + "sqli_view/viewModulesByType.tmpl",
	}

	all_files := append(files, tmpl_files...)

	templates := template.Must(template.ParseFiles(all_files...))
	renderTemplate(w, "viewModulesByType", p, templates)
}

func sqliViewResultById(title string, w http.ResponseWriter, r *http.Request){
	req := core_http.NewReq("GET", "", "json")
	couch_db := core_nosql.NewCouchDB("http://admin:123456@localhost:5984", "module_result")
	params := r.URL.Query()

	p, err1 := loadPage("sqli_view/" + title)
	if err1 != nil{
		http.Redirect(w, r, "/main/", http.StatusFound)
		return
	}

	err := json.Unmarshal([]byte(couch_db.GetByUUID(req, params["id"][0])), &p.TestResultBody)
	if err != nil{
		log.Fatalln(err)
	}

	files := []string{
		html_folder + "sqli_view/viewResultById.tmpl",
	}

	all_files := append(files, tmpl_files...)

	templates := template.Must(template.ParseFiles(all_files...))
	renderTemplate(w, "viewResultById", p, templates)
}

func sqliUiTestUrl(title string, w http.ResponseWriter, r *http.Request){
	p, err := loadPage("sqli_test/" + title)
	if err != nil{
		http.Redirect(w, r, "/main/", http.StatusFound)
		return
	}

	files := []string{
		html_folder + "sqli_test/testUrl.tmpl",
		html_folder + "sqli_test/test_start.tmpl",
		html_folder + "sqli_test/test_result.tmpl",
	}

	all_files := append(files, tmpl_files...)

	JsonUrlTestObjects := internals_sqli_modules_test.SqliUrlTestJsonObject_array{}

	p.TestUrlResult = JsonUrlTestObjects

	templates := template.Must(template.ParseFiles(all_files...))
	renderTemplate(w, "testUrl", p, templates)
}

func sqliUiByError(title string, w http.ResponseWriter, r *http.Request){
	sql_connection := &core_sql.Sql_db_connect{
		User: "root",
		Passwd: "",
		Addr: "0.0.0.0:3306",
		DBName: "sqli_test",
	}

	sql_connection.ConnectToDb()

	q_string := "SELECT * FROM user WHERE id = '" + r.FormValue("id") + "'"

	test1, _ := sql_connection.Query(q_string)

	log.Println(test1)

	p, err := loadPage("sqli_test/" + title)
	if err != nil{
		http.Redirect(w, r, "/main/", http.StatusFound)
		return
	}

	files := []string{
		html_folder + "sqli_test/byerror.tmpl",
	}

	all_files := append(files, tmpl_files...)

	templates := template.Must(template.ParseFiles(all_files...))
	renderTemplate(w, "byerror", p, templates)
}

func sqliStartModuleHandler(w http.ResponseWriter, r *http.Request, title string){
	if title == "testUrl"{
		sqliStartTesturl(title, w, r)
	}
	// }else if title == "testPost"{
	// 	sqliStartTesturl(title, w, r)
	// }
}

func sqliStartTesturl(title string, w http.ResponseWriter, r *http.Request){
	test_module := internals_sqli_modules_test.Test_url{}
	mysql_sqli_interface := sqli_mysql.NewMysqlInterface()
	test_module.RunUrlTest(r.FormValue("url"), mysql_sqli_interface)

}

// func sqliStartTestPost(title string, w http.ResponseWriter, r *http.Request){
// 	test_module := internals_sqli_modules.Test_input{}
// 	mysql_sqli_interface := sqli_mysql.NewMysqlInterface()
// 	test_module.RunUrlTest(r.FormValue("url"), mysql_sqli_interface)
// }

func StartUiServer(){
	http.HandleFunc("/brute_module/", makeHandler(bruteHandler))
	http.HandleFunc("/settings/", makeHandler(settingsHandler))
	http.HandleFunc("/http_module/", makeHandler(http_moduleHandler))
	http.HandleFunc("/resp/", makeHandler(sendRequestHandler))
	http.HandleFunc("/sqli_ui/", makeHandler(sqliUiHandler))
	http.HandleFunc("/sqli_start_module/", makeHandler(sqliStartModuleHandler))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(html_folder + "static"))))

	log.Fatal(http.ListenAndServe(":8081", nil))
}
