package routers

import (
	"errors"

	"github.com/gorilla/mux"
	"github.com/hendralatumeten/vehicles_rental/src/database/orm"
	"github.com/hendralatumeten/vehicles_rental/src/modules/v1/history"
	"github.com/hendralatumeten/vehicles_rental/src/modules/v1/users"
	"github.com/hendralatumeten/vehicles_rental/src/modules/v1/vehicles"
)

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()
	db, err := orm.New()
	if err != nil {
		return nil, errors.New("gagal init database")
	}
	users.New(mainRoute, db)
	vehicles.New(mainRoute, db)
	history.New(mainRoute, db)

	return mainRoute, nil
}
