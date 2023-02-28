package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetUser(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateUser(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteUser(w, r)
	}
	return fmt.Errorf("undefined method %s", r.Method)
}

func (s *APIServer) handleBoard(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetBoarding(w, r)
	}
	if r.Method == "POST" {
		return s.handleAddBoardinPass(w, r)
	}
	return fmt.Errorf("undefined method %s", r.Method)
}

func (s *APIServer) handleGetUser(w http.ResponseWriter, r *http.Request) error {

	users, err := s.store.getUsers()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, &users)
}
func (s *APIServer) handleGetUserByPN(w http.ResponseWriter, r *http.Request) error {
	pnstr := mux.Vars(r)["pn"]
	pn, err := strconv.Atoi(pnstr)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
	}
	user, err := s.store.getUserByPN(pn)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, &user)
}
func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	usrRequest := new(CreateUserRequest)
	// usrRequest := CreateUserRequest()

	if err := json.NewDecoder(r.Body).Decode(usrRequest); err != nil {
		return err
	}
	user := NewUser(usrRequest.Name, usrRequest.UserSchema, usrRequest.PersonalNumber)
	if err := s.store.CreateUser(user); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, user)
}
func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleGetBoarding(w http.ResponseWriter, r *http.Request) error {
	boardings, err := s.store.getBoardings()
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, &boardings)
}

func (s *APIServer) handleAddBoardinPass(w http.ResponseWriter, r *http.Request) error {
	bpRequest := new(AddBoardRequest)

	if err := json.NewDecoder(r.Body).Decode(bpRequest); err != nil {
		return err
	}
	bp := NewBoardingPass(bpRequest.Name, bpRequest.Booking, bpRequest.JDate, bpRequest.TypePasss)
	if err := s.store.CreateBoardinPass(bp); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, bp)
}
func (s *APIServer) handleGetBoardingPassByBooking(w http.ResponseWriter, r *http.Request) error {
	booking := mux.Vars(r)["booking"]
	bookingPass, err := s.store.getBoardingPassByBooking(booking)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
	}
	return WriteJSON(w, http.StatusOK, &bookingPass)
}
func (s *APIServer) handleUpdateBoardingPass(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type APIServer struct {
	listenAddr string
	store      Stogage
}
type ApiError struct {
	Error string
}
type apiFunc func(http.ResponseWriter, *http.Request) error

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
func NewAPIServer(listenAddr string, store Stogage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}
func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/user", makeHTTPHandlerFunc(s.handleUser))

	router.HandleFunc("/user/{pn}", makeHTTPHandlerFunc(s.handleGetUserByPN))

	router.HandleFunc("/boardingpass", makeHTTPHandlerFunc(s.handleBoard))

	router.HandleFunc("/boardingpass/{booking}", makeHTTPHandlerFunc(s.handleGetBoardingPassByBooking))

	log.Println("json api server run on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}
