package ctrl

import (
	"encoding/json"
	"github.com/davewhit3/std-server/model"
	"github.com/davewhit3/std-server/server"
	"github.com/samber/lo"
	"net/http"
	"slices"
	"strconv"
	"sync"
)

var mu sync.Mutex

var UserData []*model.User = []*model.User{
	{ID: 1, FirstName: "Wye", LastName: "Cana", Email: "wcana0@archive.org"},
	{ID: 2, FirstName: "Pamela", LastName: "Scain", Email: "pscain1@vk.com"},
	{ID: 3, FirstName: "Haroun", LastName: "Pendlington", Email: "hpendlington2@fotki.com"},
	{ID: 4, FirstName: "Amity", LastName: "Reddel", Email: "areddel3@answers.com"},
	{ID: 5, FirstName: "Kane", LastName: "Bruckmann", Email: "kbruckmann4@comsenz.com"},
	{ID: 6, FirstName: "Chrissie", LastName: "Culverhouse", Email: "cculverhouse5@go.com"},
	{ID: 7, FirstName: "Sol", LastName: "Keddle", Email: "skeddle6@twitter.com"},
}

func SetupUserHandlers(srv *http.ServeMux) {
	srv.HandleFunc("POST /users", createUser)
	srv.HandleFunc("GET /users", listUsers)
	srv.HandleFunc("GET /users/{id}", listUser)
	srv.HandleFunc("PUT /users/{id}", updateUser)
	srv.HandleFunc("DELETE /users/{id}", deleteUser)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var u model.User

	if r.ContentLength > 0 {
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			server.ResponseError(w, err)
			return
		}
		mu.Lock()
		defer mu.Unlock()
		u.ID = len(UserData) + 1

		if u.Valid() {

			UserData = append(UserData, &u)

			server.ResponseCode(w, http.StatusCreated)
			return
		}
	}

	server.ResponseError(w, server.ErrInputNotValid)
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	server.ResponseData(w, UserData)
}

func listUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		server.ResponseError(w, server.ErrInputNotValid)
		return
	}

	u, ok := lo.Find(UserData, func(u *model.User) bool {
		return u.ID == id
	})

	if ok {
		server.ResponseData(w, u)
		return
	}

	server.ResponseCode(w, http.StatusNotFound)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || r.ContentLength == 0 {
		server.ResponseError(w, server.ErrInputNotValid)
		return
	}

	var u model.User

	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		server.ResponseError(w, err)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	_, i, ok := lo.FindIndexOf(UserData, func(d *model.User) bool {
		return d.ID == id
	})

	if !ok {
		server.ResponseCode(w, http.StatusNotFound)
		return
	}

	u.ID = id
	if !u.Valid() {
		server.ResponseError(w, server.ErrInputNotValid)
		return
	}

	UserData[i] = &u

	server.ResponseCode(w, http.StatusNoContent)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		server.ResponseError(w, server.ErrInputNotValid)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	_, i, ok := lo.FindIndexOf(UserData, func(d *model.User) bool {
		return d.ID == id
	})

	if !ok {
		server.ResponseCode(w, http.StatusNotFound)
		return
	}

	UserData = slices.Replace(UserData, i, i+1)

	server.ResponseCode(w, http.StatusAccepted)
}
