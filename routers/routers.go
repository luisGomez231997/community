package routers

import (
	"encoding/json"
	"enconding/json"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user) //object string of uniquely reading
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {

	}

}
