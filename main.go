package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

var userNpw = make(map[string]string, 0)
var logedin = make(map[string]bool, 0)

func main() {
	router := chi.NewRouter()
	router.Post("/log-in", login)
	router.Post("/log-out", logout)
	router.Post("/newUser", newUser)
	http.ListenAndServe(":8080", router)
}

func login(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var req Request
	json.NewDecoder(r.Body).Decode(&req)
	value, ok := userNpw[req.Username]
	if !ok {
		w.Write([]byte("Username not found"))
		return
	}
	if req.Password == value {
		w.Write([]byte("You are logged-in"))
		logedin[req.Username] = true
		return
	}
	w.Write(([]byte("Password is wrong")))
}
func newUser(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var req Request
	json.NewDecoder(r.Body).Decode(&req)
	userNpw[req.Username] = req.Password
	logedin[req.Username] = false
}

func logout(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var req Request
	json.NewDecoder(r.Body).Decode(&req)
	value, ok := userNpw[req.Username]
	if !ok {
		w.Write([]byte("Username not found"))
		return
	}
	if value == req.Password {
		w.Write([]byte("You are logged-out"))
		logedin[req.Username] = false
		return
	}
	w.Write(([]byte("Password is wrong")))
}
