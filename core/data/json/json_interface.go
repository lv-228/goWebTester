package core_data_json

import(
	"encoding/json"
	//"log"
	"os"
	"core/http"
	"core/os"
)

type Save_to_json interface{
	ToByte() []byte
	//ToMap()
}

type HttpJsonObject struct{
	id int
	Request_obj *core_http.Req
	Response_obj *core_http.Resp
}

func (h *HttpJsonObject) CreateHttpJsonObject(request *core_http.Req, response *core_http.Resp){
	h.Request_obj = request
	h.Response_obj = response
}

func (h *HttpJsonObject) ToByte() []byte{
	answer, err := json.Marshal(h)
	core_os.CheckErrValue(err, "Json marshal error! Object: response")
	return answer
}

type JsonFile struct{
	HttpJsonObject_objs []HttpJsonObject
}

func (j *JsonFile) ToByte() []byte{
	answer, err := json.Marshal(j)
	core_os.CheckErrValue(err, "Json marshal error! Object: response")
	return answer
}

func (j *JsonFile) SaveJsonFile(httpJsonObject *HttpJsonObject){
	filename := core_os.GetYearMonthDayNow()
	j.GetJsonObject(filename)
	j.HttpJsonObject_objs = append(j.HttpJsonObject_objs, *httpJsonObject)
	rawDataOut, err1 := json.MarshalIndent(&j, "", "  ")
	core_os.CheckErrValue(err1, "Marshal error!")
	err2 := os.WriteFile(filename, rawDataOut, 0)
	core_os.CheckErrValue(err2, "Write file error!")
}

func (j *JsonFile) GetJsonObject(filename string){
	jsonInFile, err1 := os.ReadFile(filename)
	if err1 != nil{
		SaveObjectInJsonFile(j)
	}
	jsonInFile, _ = os.ReadFile(filename)
	//CheckErrValue(err1, "Ошибка открытия файла!")
	err2 := json.Unmarshal(jsonInFile, &j)
	core_os.CheckErrValue(err2, "Ошибка дессериализации!")
}

func (j *JsonFile) JsonFileToByte(filename string) []byte{
	j.GetJsonObject(filename)
	return j.ToByte()
}

func (j *JsonFile) ToMap() map[string][]HttpJsonObject {
	answer := map[string][]HttpJsonObject{}
	err := json.Unmarshal(j.ToByte(), &answer)
	core_os.CheckErrValue(err, "Ошибка! Не получилось перевести файл в таблицу")
	return answer
}

func SaveObjectInJsonFile(obj Save_to_json){
	jsonFolder := "http_json"
	jsonString := obj.ToByte()
	os.WriteFile(jsonFolder + core_os.GetYearMonthDayNow(), jsonString, os.ModePerm)
}