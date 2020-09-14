package whattimeisitrightnow

import (
	"log"
	"net/http"
	"time"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	w.Write([]byte(time.Now().String()))
}