package http_funcs

import(
	"encoding/json"
	//"log"
	"os"
)

type Save_to_json interface{
	ToByte() []byte
	//ToMap()
}

type HttpJsonObject struct{
	id int
	Request_obj *Req
	Response_obj *Resp
}

func (h *HttpJsonObject) CreateHttpJsonObject(request *Req, response *Resp){
	h.Request_obj = request
	h.Response_obj = response
}

func (h *HttpJsonObject) ToByte() []byte{
	answer, err := json.Marshal(h)
	CheckErrValue(err, "Json marshal error! Object: response")
	return answer
}

type JsonFile struct{
	HttpJsonObject_objs []HttpJsonObject
}

func (j *JsonFile) ToByte() []byte{
	answer, err := json.Marshal(j)
	CheckErrValue(err, "Json marshal error! Object: response")
	return answer
}

func (j *JsonFile) SaveJsonFile(httpJsonObject *HttpJsonObject){
	filename := GetYearMonthDayNow()
	j.GetJsonObject(filename)
	j.HttpJsonObject_objs = append(j.HttpJsonObject_objs, *httpJsonObject)
	rawDataOut, err1 := json.MarshalIndent(&j, "", "  ")
	CheckErrValue(err1, "Marshal error!")
	err2 := os.WriteFile(filename, rawDataOut, 0)
	CheckErrValue(err2, "Write file error!")
}

func (j *JsonFile) GetJsonObject(filename string){
	jsonInFile, err1 := os.ReadFile(filename)
	if err1 != nil{
		SaveObjectInJsonFile(j)
	}
	jsonInFile, _ = os.ReadFile(filename)
	//CheckErrValue(err1, "Ошибка открытия файла!")
	err2 := json.Unmarshal(jsonInFile, &j)
	CheckErrValue(err2, "Ошибка дессериализации!")
}

func (j *JsonFile) JsonFileToByte(filename string) []byte{
	j.GetJsonObject(filename)
	return j.ToByte()
}

func (j *JsonFile) ToMap() map[string][]HttpJsonObject {
	answer := map[string][]HttpJsonObject{}
	err := json.Unmarshal(j.ToByte(), &answer)
	CheckErrValue(err, "Ошибка! Не получилось перевести файл в таблицу")
	return answer
}