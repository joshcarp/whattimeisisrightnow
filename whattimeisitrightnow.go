package whattimeisitrightnow

import (
	"net/http"
	"time"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(time.Now().String()))
}