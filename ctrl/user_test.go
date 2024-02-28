package ctrl

import (
	"encoding/json"
	"github.com/davewhit3/std-server/model"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	UserInput = `{
		"first_name": "Dave",
		"last_name": "White",
		"email": "dave@white.com"
	}`
)

func TestCreateUser(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPost,
		"/users",
		strings.NewReader(UserInput),
	)
	w := httptest.NewRecorder()

	usersCounter := len(UserData)
	createUser(w, req)
	resp := w.Result()

	require.Equal(t, http.StatusCreated, resp.StatusCode)
	require.Equal(t, usersCounter+1, len(UserData))
}

func TestListUsers(t *testing.T) {
	var data []*model.User

	req := httptest.NewRequest(
		http.MethodGet,
		"/users",
		nil,
	)
	w := httptest.NewRecorder()
	listUsers(w, req)

	resp := w.Result()

	require.Nil(t, json.NewDecoder(resp.Body).Decode(&data))
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, len(UserData), len(data))
}

func TestListUser(t *testing.T) {
	var data *model.User

	req := httptest.NewRequest(
		http.MethodGet,
		"/users/1",
		nil,
	)
	req.SetPathValue("id", "1")

	w := httptest.NewRecorder()
	listUser(w, req)

	resp := w.Result()

	require.Nil(t, json.NewDecoder(resp.Body).Decode(&data))
	require.Equal(t, http.StatusOK, resp.StatusCode)

	require.Equal(t, UserData[0].ID, data.ID)
	require.Equal(t, UserData[0].FirstName, data.FirstName)
	require.Equal(t, UserData[0].LastName, data.LastName)
	require.Equal(t, UserData[0].Email, data.Email)
}

func TestNotFoundUser(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/users/9999",
		nil,
	)
	req.SetPathValue("id", "9999")

	w := httptest.NewRecorder()
	listUser(w, req)

	resp := w.Result()

	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestUpdateUser(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPut,
		"/users/1",
		strings.NewReader(UserInput),
	)
	req.SetPathValue("id", "1")

	w := httptest.NewRecorder()
	updateUser(w, req)

	resp := w.Result()

	require.Equal(t, http.StatusNoContent, resp.StatusCode)

	require.Equal(t, 1, UserData[0].ID)
	require.Equal(t, "Dave", UserData[0].FirstName)
	require.Equal(t, "White", UserData[0].LastName)
	require.Equal(t, "dave@white.com", UserData[0].Email)
}

func TestDeleteUser(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodDelete,
		"/users/1",
		nil,
	)
	req.SetPathValue("id", "1")

	w := httptest.NewRecorder()
	usersCounter := len(UserData)
	deleteUser(w, req)

	resp := w.Result()

	require.Equal(t, http.StatusAccepted, resp.StatusCode)
	require.Equal(t, usersCounter-1, len(UserData))
}
