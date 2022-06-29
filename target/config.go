package target

import(
	"log"
	"io/ioutil"
  "os"
  "fmt"
  "encoding/json"
)

type Config struct{
  Http_user_headers map[string]string
  Creds map[string]string
  Target map[string]string
}

func GetConfig() []byte{
	data, err := ioutil.ReadFile("./target/config.json")
    if err != nil {
      log.Fatalln(err)
    }

    return data
}

func AddHeader(){
  f, err := os.OpenFile("./target/config.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
  if err != nil {
    log.Fatalln(err)
    return
  }
  defer f.Close()

  new_header := map[string]string{
    "new_test": "test",
  }

  byte_header, err := json.Marshal(new_header)
  if err != nil {
    log.Fatalln(err)
  }

  if _, err := fmt.Fprintf(f, "%s\n", byte_header); err != nil {
    log.Fatalln(err)
  }

}