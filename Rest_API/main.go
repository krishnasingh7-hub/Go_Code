package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"sync"
)

var (
	listUsersRe   = regexp.MustCompile(`^\/users[\/]*$`)
	getUsersRe    = regexp.MustCompile(`^\/users\/(\d+)$`) //users/123
	createUsersRe = regexp.MustCompile(`^\/users[\/]*$`)
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type datastore struct {
	m             map[string]user
	*sync.RWMutex // to handle concurrency by defferet goroutines at same time , needed only for in memory datastore
}

type userHandler struct {
	store *datastore
}

// g
func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listUsersRe.MatchString(r.URL.Path):
		h.ListUsers(w, r)
		return
	case r.Method == http.MethodGet && getUsersRe.MatchString(r.URL.Path):
		h.GetUsers(w, r)
		return
	case r.Method == http.MethodPost && listUsersRe.MatchString(r.URL.Path):
		h.CreateUsers(w, r)
		return
	default:
		notFound(w, r)
		return
	}

}
func (h *userHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users := make([]user, 0, len(h.store.m))
	h.store.RLock()
	for _, u := range h.store.m {
		users = append(users, u)
	}
	h.store.RUnlock()
	jsonBytes, err := json.Marshal(users)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
func (h *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	matches := getUsersRe.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		notFound(w, r)
		return
	}
	h.store.RLock()
	user, ok := h.store.m[matches[1]]
	h.store.RUnlock()
	if !ok {
		notFound(w, r)
		return
	}
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *userHandler) CreateUsers(w http.ResponseWriter, r *http.Request) {
	u := &user{}
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		badRequest(w, r)
		return
	}
	h.store.Lock()
	h.store.m[u.ID] = *u
	h.store.Unlock()
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}
func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error":"Not Found"}`))
}
func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"error":"Internal Server error"}`))
}
func badRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`{"error":"Bad Request"}`))
}

func main() {
	fmt.Println("HTTP Server Without Framework")
	mux := http.NewServeMux()
	userH := &userHandler{
		store: &datastore{
			m: map[string]user{
				"1": user{ID: "1", Name: "bob"},
			},
			RWMutex: &sync.RWMutex{},
		},
	}
	mux.Handle("/users/", userH)
	mux.Handle("/users", userH)
	http.ListenAndServe("localhost:8080", mux)
	// http://localhost:8080/users/1
	// http://localhost:8080/users
}
