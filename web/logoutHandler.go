package web

import (
	"github.com/hopperteam/hopper-account/config"
	"net/http"
	"time"
)

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("HOPPER_SESSION")

	if err != nil {
		apiError(w, "No session", http.StatusForbidden)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "HOPPER_SESSION", Value: "", Path: "/", Expires: time.Unix(0,0), Domain: config.Config.CookieDomainName })

	apiResult(w, nil)
}

