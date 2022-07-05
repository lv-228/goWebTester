package server_ui

import(
	"os"
	"net/http"
	"log"
	"html/template"
	"regexp"
	"web_tester/http_funcs"
	"web_tester/target"
	"encoding/json"
	//"strings"
	//"strconv"
)

type Page struct {
	Title string
	Body []byte
	JsonList *http_funcs.JsonFile
	Headers map[string]string
}

var conf target.Config

var html_folder = "./server_ui/html/"

var validPath = regexp.MustCompile("^/(brute_module|settings|http_module|resp)/([a-zA-Z0-9]+)$")

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

	jsonFile := &http_funcs.JsonFile{}

	jsonFile.GetJsonObject(http_funcs.GetYearMonthDayNow())

	p.JsonList = jsonFile

	renderTemplate(w, "index", p, templates)
}

func sendRequestHandler(w http.ResponseWriter, r *http.Request, title string){

	data := target.GetConfig()

	err := json.Unmarshal(data, &conf)
    if err != nil {
        log.Println("error:", err)
    }

    headerData := &http_funcs.HeaderData{}
    headerReplace := headerData.CreateFromTextArea(r.FormValue("list_headers"))

    headerReplace.AppendValues([]string{r.FormValue("values")}, true)

    //log.Println(headerReplace)

    request := &http_funcs.Req{
    	Req_type: r.FormValue("method"),
    	Url: r.FormValue("url"),
    	Headers_obj: headerData,
    }

    response := request.SendAndGetResult(r.FormValue("data"))

    jsonHttpObject := &http_funcs.HttpJsonObject{
    	Request_obj: request,
    	Response_obj: response,
    }

    jsonFile := &http_funcs.JsonFile{}

    jsonFile.SaveJsonFile(jsonHttpObject)

    log.Println(jsonFile)

	// request1 := &http_funcs.ReqData{
	// 	Req_type: r.FormValue("method"),
	// 	Url: conf.Target["host"] + r.FormValue("url"),
	// 	Headers: headers_request,
	// }

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

func StartUiServer(){
	http.HandleFunc("/brute_module/", makeHandler(bruteHandler))
	http.HandleFunc("/settings/", makeHandler(settingsHandler))
	http.HandleFunc("/http_module/", makeHandler(http_moduleHandler))
	http.HandleFunc("/resp/", makeHandler(sendRequestHandler))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(html_folder + "static"))))

	log.Fatal(http.ListenAndServe(":8081", nil))
}
