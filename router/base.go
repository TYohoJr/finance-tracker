package router

import (
	"finance-tracker/model"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type Server struct {
	DB     *model.DB
	Router chi.Router
}

func (s *Server) Initialize() {
	s.Router = chi.NewRouter()
	s.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Access-Control-Allow-Headers", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300,
		Debug:            false,
	}))

	fmt.Println("Backend Initialized")
	err := http.ListenAndServe("0.0.0.0:8080", s.Router)
	if err != nil {
		log.Fatal(err)
	}
}
