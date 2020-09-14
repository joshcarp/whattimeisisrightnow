package whattimeisitrightnow

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	log.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte(time.Now().String()))

}
