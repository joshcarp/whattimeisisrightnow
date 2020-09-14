package whattimeisitrightnow

import (
	"log"
	"net/http"
	"time"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	log.Println(r.Body)
	w.Write([]byte(time.Now().String()))

}
