package controller

import (
	"encoding/json"
	"errors"
	"finance-tracker/auth"
	"finance-tracker/model"
	"net/http"
)

func (s *Server) AppUserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		aui := model.AppUser{}
		err := json.NewDecoder(r.Body).Decode(&aui)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = s.validateNewAppUser(&aui)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		existingAppUser, err := s.DB.GetAppUserByEmail(*aui.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if existingAppUser != nil {
			http.Error(w, "user already exists with that email", http.StatusBadRequest)
			return
		}
		_, err = s.DB.CreateAppUser(&aui)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(201)
		return
	}
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		appUser := &model.AppUser{}
		err := json.NewDecoder(r.Body).Decode(&appUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		u, err := s.DB.GetAppUserByEmail(*appUser.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if u == nil {
			http.Error(w, "invalid email/password", http.StatusInternalServerError)
			return
		}
		err = u.VerifyPassword(*appUser.Password)
		if err != nil {
			http.Error(w, "invalid email/password", http.StatusInternalServerError)
			return
		}
		token, err := auth.CreateToken(*u.ID, *u.Email, *u.FirstName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		err = json.NewEncoder(w).Encode(map[string]string{
			"token": token,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		return
	}
}

func (s *Server) validateNewAppUser(u *model.AppUser) error {
	if u.Email == nil {
		return errors.New("email cannot be empty")
	}
	if u.FirstName == nil {
		return errors.New("first name cannot be empty")
	}
	if u.Password == nil {
		return errors.New("password cannot be empty")
	}
	if *u.Password == "" {
		return errors.New("password cannot be empty")
	}
	return nil
}
