package endpoint

import (
	"net/http"

	"encoding/json"

	"github.com/go-kickstart-orm/business"
	"github.com/go-kickstart-orm/model/entity"
)

//SaveUser endpoint path on /user/save
func SaveUser(w http.ResponseWriter, req *http.Request) {
	var user entity.User
	json.NewDecoder(req.Body).Decode(&user)
	uf := business.UserService{}
	uf.Save(&user)
}

//GetAllUser endpoint path on /user/all
func GetAllUser(w http.ResponseWriter, req *http.Request) {
	uf := business.UserService{}
	users := uf.GetAll()
	js, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
