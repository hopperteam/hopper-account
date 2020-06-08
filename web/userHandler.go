package web

import (
	"github.com/hopperteam/hopper-account/security"
	"net/http"
)

func userGetHandler(w http.ResponseWriter, r *http.Request) {
	sess, err := r.Cookie("HOPPER_SESSION")
	if err != nil {
		apiError(w, "No session", http.StatusForbidden)
		return
	}

	usr, err := security.DecodeSession(sess.Value)
	if err != nil {
		apiError(w, "No session", http.StatusForbidden)
		return
	}

	apiResult(w, usr)
}
