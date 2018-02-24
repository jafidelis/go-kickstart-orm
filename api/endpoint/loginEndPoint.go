package endpoint

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kickstart-orm/business"
	"github.com/go-kickstart-orm/model/entity"
)

func Login(w http.ResponseWriter, req *http.Request) {
	var user entity.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Não foi possível ler o usuário %s", err)
		return
	}

	ls := business.LoginService{}

	fmt.Println(user)

	token, err := ls.Login(&user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Usuário ou senha inválido ", err)
	}

	jsonResult, err := json.Marshal(token)
	if err != nil {
		fmt.Fprintln(w, "Erro ao gerar o json")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Conten-Type", "application/json")
	w.Write(jsonResult)
}
