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

var dataMock = models.Users{
	{Username: "hendra", Role: "admin"},
	{Username: "user", Role: "user"},
}

func TestCtrlGetAll(t *testing.T) {

	//get respon
	w := httptest.NewRecorder()
	mux := mux.NewRouter()

	repos.mock.On("FindAll").Return(&dataMock, nil)
	mux.HandleFunc("/test/user", ctrl.GetAll)

	mux.ServeHTTP(w, httptest.NewRequest("GET", "/test/user", nil))

	var users models.Users
	respone := libs.Response{
		Data: &users,
	}

	if err := json.Unmarshal(w.Body.Bytes(), &respone); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, w.Code, "status code must be 200")
	assert.False(t, respone.IsError, "isError must be false")

}
