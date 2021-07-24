package middlewears

import (
	"net/http"

	"github.com/luisGomez231997/community/database"
)

func checkBD(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if database.CheckConnection() == 0 {
			http.Error(rw, "Conecction lossed", 500)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
