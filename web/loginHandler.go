package web

import (
	"encoding/json"
	"github.com/hopperteam/hopper-account/config"
	"github.com/hopperteam/hopper-account/model"
	"github.com/hopperteam/hopper-account/security"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type loginRequestBody struct {
	EMail string `json:"email"`
	Password string `json:"password"`
}

type loginResponseBody struct {
	Session string `json:"session"`
	User *model.SessionUser `json:"user"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	body := &loginRequestBody{}

	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	usr, err := model.LoadUserByEmail(body.EMail)

	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusForbidden)
		return
	}

	if !security.CheckPasswordHash(usr.PwHash, body.Password) {
		http.Error(w, "Invalid credentials", http.StatusForbidden)
		return
	}

	sessUsr := usr.ToSessionUser()
	createAndReplySession(w, sessUsr)
}

func createAndReplySession(w http.ResponseWriter, sessUsr *model.SessionUser) {
	expire := time.Now().Add(config.Config.SessionTime)

	session, err := security.CreateSession(sessUsr, expire.Unix())
	if err != nil {
		log.Error(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	http.SetCookie(w, &http.Cookie{Name: "HOPPER_SESSION", Value: session, Path: "/", Expires: expire, Domain: config.Config.CookieDomainName })

	err = json.NewEncoder(w).Encode(&loginResponseBody{
		Session: session,
		User: sessUsr,
	})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
