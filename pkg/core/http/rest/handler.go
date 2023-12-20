package rest

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"hexagonal-structure/pkg/domain/adding"
	"hexagonal-structure/pkg/domain/deleting"
	"hexagonal-structure/pkg/domain/listing"
	"net/http"
)

func Handler(a adding.Service, l listing.Service, d deleting.Service) http.Handler {
	router := httprouter.New()

	router.GET("/users", getUsers(l))
	router.GET("/users/:id", getUser(l))

	router.POST("/users", addUser(a))

	router.DELETE("/users/:id", deleteUser(d))

	return router
}

// addUser returns a handler for POST /users requests
func addUser(s adding.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		var newUser adding.User
		err := decoder.Decode(&newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		s.AddUser(newUser)
		// error handling omitted for simplicity

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New user added.")
	}
}

// getUsers returns a handler for GET /users requests
func getUsers(s listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		list := s.GetUsers()
		json.NewEncoder(w).Encode(list)
	}
}

// getUser returns a handler for GET /users/:id requests
func getUser(s listing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		user, err := s.GetUser(p.ByName("id"))
		if err == listing.ErrNotFound {
			http.Error(w, "The user you requested does not exist.", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

// deleteUser returns a handler for DELETE /users/:id requests
func deleteUser(s deleting.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		err := s.DeleteUser(p.ByName("id"))

		if err == deleting.ErrNotFound {
			http.Error(w, "The user you requested does not exist.", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("User deleted successfully.")
	}
}
