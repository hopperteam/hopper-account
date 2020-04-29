package web

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	http *http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (server *Server) Start() {
	r := mux.NewRouter()

	r.Path("/api").Methods("GET").HandlerFunc(InfoHandler)

	r.Path("/api/login").Methods("POST").HandlerFunc(loginHandler)
	r.Path("/api/signUp").Methods("POST").HandlerFunc(signUpHandler)

	userRoute := r.Path("/api/user")

	userRoute.Methods("GET")
	userRoute.Methods("PUT")

	log.Fatal(http.ListenAndServe(":80",  handlers.CORS()(r)))
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hopper Auth Server v1.0")
	if err != nil {
		log.Fatal(err.Error())
	}
}
