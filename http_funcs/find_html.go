package http_funcs

import (
	"log"
	"strings"
	"regexp"
)

func GetForms(boof_body string) map[int]string {

	forms := map[int]string{}
	for i := 0; ;i++{
		start_form := strings.Index(boof_body, "<form")
		if start_form == -1 {
			break
		}
		end_form := strings.Index(boof_body, "</form>")
		forms[i] = boof_body[start_form:end_form+7]
		boof_body = boof_body[end_form+7:]
	}

	return forms
}

func GetFormAttributes(form string) string{
	start_form := strings.Index(form, "<form")
	if start_form == -1 {
		log.Fatal("ERROR! Need <form>")
	}
	end_form := strings.Index(form, ">")
	return form[start_form:end_form+1]
}

func GetInputs(forms map[int]string) map[int][]string {
	answer := map[int][]string{}
	for j, elem := range forms{
		for i :=0; ;i++{
			input_start := strings.Index(elem, "<input")
			if input_start == -1 {
				break
			}
			input_end := strings.Index(elem[input_start:], ">") + input_start
			answer[j] = append(answer[j],elem[input_start:input_end+1])
			elem = elem[input_end+1:]
		}
	}
	return answer
}

func GetHtmlAttributesByName(html_object string) []string{
	re := regexp.MustCompile(`[a-zA-Z]*?=("|').*?("|')`)
	return re.FindAllString(html_object, -1)
}

func GetHtmlAttributeValue(html_attribute string) (string, string){
	equal := strings.Index(html_attribute, "=")
	attribute := html_attribute[:equal]
	value := html_attribute[equal+1:]

	return attribute, strings.Trim(value, "\"")
}

func GetHtmlTagByNameAndClass(body string, tag string, class string) string{
	re := regexp.MustCompile(`<` + tag + `\s?.*<\/` + tag + `>`)
	find_str := ""
	for _, elem := range re.FindAllString(body, -1){
		class_index := strings.Index(elem, "class=" + class)
		if class_index != -1{
			find_str = elem
		}
	}
	return find_str
}