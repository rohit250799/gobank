package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func (http.ResponseWriter, *http.Request) error 

type apiError struct {
	Error string
}
	
func makeHttpHandleFunc (f apiFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		if err := f(w, r); err != nil {
			//handle error
			WriteJson(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}
}

type ApiServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
	}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHttpHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHttpHandleFunc(s.handleGetAccount))

	log.Println("JSON API Server running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *ApiServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET"{
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST"{
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE"{
		return s.handleDeleteAccount(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	fmt.Println(id)
	//account := NewAccount("Rohit", "Sarkar")
	return WriteJson(w, http.StatusOK, &Account{})
}

func (s *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

