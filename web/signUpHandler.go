package web

import (
	"encoding/json"
	"github.com/hopperteam/hopper-account/config"
	"github.com/hopperteam/hopper-account/model"
	"github.com/hopperteam/hopper-account/security"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type signUpRequestBody struct {
	EMail string `json:"email"`
	Password string `json:"password"`
	FName string `json:"firstName"`
	LName string `json:"lastName"`
}


func signUpHandler(w http.ResponseWriter, r *http.Request) {
	body := &signUpRequestBody{}

	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	usr, err := model.LoadUserByEmail(body.EMail)

	if err == nil {
		http.Error(w, "Email already exists", http.StatusBadRequest)
		return
	}

	pwHash, err := security.GetPasswordHash(body.Password)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	usr = &model.User{
		FirstName:    body.FName,
		LastName:     body.LName,
		EMail:        body.EMail,
		PwHash:       pwHash,
		Roles:        config.Config.DefaultRoles,
	}

	id, err := model.CreateUser(usr)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	usr.Id = bson.ObjectIdHex(id)
	createAndReplySession(w, usr.ToSessionUser())
}
