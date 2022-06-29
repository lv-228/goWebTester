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
	"strconv"
)

type Page struct {
	Title string
	Body []byte
	Headers map[string]string
}

var conf target.Config

var html_folder = "./server_ui/html/"

var validPath = regexp.MustCompile("^/(brute_module|settings|req_resp|resp)/([a-zA-Z0-9]+)$")

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

func req_respHandler(w http.ResponseWriter, r *http.Request, title string){
	p, err := loadPage("req_resp/" + title)
	if err != nil{
		http.Redirect(w, r, "/main/", http.StatusFound)
		return
	}
	files := []string{
		html_folder + "req_resp/index.tmpl",
		html_folder + "req_resp/request.tmpl",
	}

	all_files := append(files, tmpl_files...)
	templates := template.Must(template.ParseFiles(all_files...))

	data := target.GetConfig()
	err1 := json.Unmarshal(data, &conf)
    if err1 != nil {
        log.Println("error:", err)
    }

	p.Headers = conf.Http_user_headers
	renderTemplate(w, "index", p, templates)
}

func sendRequestHandler(w http.ResponseWriter, r *http.Request, title string){

	data := target.GetConfig()

	err := json.Unmarshal(data, &conf)
    if err != nil {
        log.Println("error:", err)
    }

    headers_request, replace_header := http_funcs.ParseTextareaHeaders(r.FormValue("list_headers"))

	request1 := &http_funcs.ReqData{
		Req_type: r.FormValue("method"),
		Url: conf.Target["host"] + r.FormValue("url"),
		Headers: headers_request,
	}

	if len(replace_header) == 1{
		values, err := strconv.Atoi(r.FormValue("values"))
		if err != nil{
			log.Fatalln("Ошибка диапазона значений")
		}

		step, err := strconv.Atoi(r.FormValue("step"))
		if err != nil{
			log.Fatalln("Ошибка шага")
		}

		for i := 1; i < values; i += step {
			for key, elem := range replace_header{
				headers_request[key] = http_funcs.ValueHeaderReplace(elem, strconv.Itoa(i), http_funcs.Var_simbol_http)
			}
			request1.Headers = headers_request
			headers_response, body := http_funcs.SendRequest(request1, r.FormValue("data"))
			log.Println(headers_response)
			log.Println(http_funcs.GetHtmlTagByNameAndClass(body, "p", "is-warning"))
		}
	}

	bf_file_path := r.FormValue("bf_journal")

	if(bf_file_path != ""){
		http_funcs.RequestRepeater(http_funcs.SendRequest, request1, r.FormValue("data"), bf_file_path)
		return
	}

	_, body_response := http_funcs.SendRequest(request1, r.FormValue("data"))

	p := &Page{Title: title, Body: []byte(body_response)}

	files := []string{
		html_folder + "req_resp/response.tmpl",
	}

	all_files := append(files, tmpl_files...)
	templates := template.Must(template.ParseFiles(all_files...))

	renderTemplate(w, "response", p, templates)

}

func StartUiServer(){
	http.HandleFunc("/brute_module/", makeHandler(bruteHandler))
	http.HandleFunc("/settings/", makeHandler(settingsHandler))
	http.HandleFunc("/req_resp/", makeHandler(req_respHandler))
	http.HandleFunc("/resp/", makeHandler(sendRequestHandler))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(html_folder + "static"))))

	log.Fatal(http.ListenAndServe(":8081", nil))
}
