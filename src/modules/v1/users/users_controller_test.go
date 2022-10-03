package users

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/hendralatumeten/vehicles_rental/src/database/orm/models"
	"github.com/hendralatumeten/vehicles_rental/src/libs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repos = RepoMock{mock.Mock{}}
var service = NewService(&repos)
var ctrl = NewCtrl(service)

var dataMock = models.User{
	Username: "Hendra",
	Email:    "Hendra@gmail.com",
}

func TestCtrlGetAll(t *testing.T) {
	w := httptest.NewRecorder()
	mux := mux.NewRouter()

	repos.mock.On("FindAll").Return(&dataMock, nil)

	mux.HandleFunc("/test/users", ctrl.GetAll)

	mux.ServeHTTP(w, httptest.NewRequest("GET", "/test/users", nil))

	var users models.User
	response := libs.Response{
		Data: &users,
	}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, w.Code, "status is not 200")
	assert.Nil(t, response.IsError, "error is not nil")
}
