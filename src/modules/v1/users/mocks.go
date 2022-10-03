package users

import (
	"github.com/hendralatumeten/vehicles_rental/src/database/orm/models"
	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock mock.Mock
}

func (m *RepoMock) FindAll() (*models.Users, error) {
	args := m.mock.Called()
	//return modals
	return args.Get(0).(*models.Users), nil
}
func (m *RepoMock) Save(data *models.User) (*models.User, error) {
	args := m.mock.Called(data)
	return args.Get(0).(*models.User), nil
}

func (m *RepoMock) FindByUsername(data string) (*models.User, error) {

	args := m.mock.Called(data)
	//return modals
	return args.Get(0).(*models.User), nil
}
func (m *RepoMock) FindByRole(role string) (*models.User, error) {
	args := m.mock.Called()

	//return modals
	return args.Get(0).(*models.User), nil
}

func (m *RepoMock) UserExsist(username, email string) bool {

	//return modals
	return true
}

func (m *RepoMock) Delete(data *models.User, params string) (*models.User, error) {
	args := m.mock.Called()
	//return modals
	return args.Get(0).(*models.User), nil
}
func (m *RepoMock) Update(data *models.User, params string, filename string) (*models.User, error) {
	args := m.mock.Called()
	//return modals
	return args.Get(0).(*models.User), nil
}
