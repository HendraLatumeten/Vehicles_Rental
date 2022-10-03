package users

import (
	"testing"

	"github.com/hendralatumeten/vehicles_rental/src/database/orm/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMock = models.Users{
		{Username: "admin", Role: "admin"},
		{Username: "user", Role: "user"},
	}

	//repo.Mock.On()
	repo.mock.On("FindAll").Return(&dataMock, nil)

	data := service.GetAll()

	//pecah data respone
	result := data.Data.(*models.Users)

	for i, v := range *result {
		assert.Equal(t, dataMock[i].Username, v.Username, "expect username from data mock")
		assert.Equal(t, dataMock[i].Role, v.Role, "expect Role from data mock")
	}
}

// func TestAdd(t *testing.T) {
// 	repo := RepoMock{mock.Mock{}}
// 	service := NewService(&repo)

// 	var dataMock = models.User{Username: "Hendra", Email: "hendra@gmail.com", Image: "hendra.jpg"}

// 	repo.mock.On("Save", &dataMock).Return(&dataMock, nil)
// 	data := service.Add(&dataMock, dataMock.Image)

// 	var expectName = "hendra.jpg"
// 	result, err := data.Data.(string)

// 	if err != true {
// 		fmt.Println(err)
// 	}

// 	assert.Equal(t, expectName, result, "name is not hendra")
// 	//	assert.IsType(t, "string", result, "result is not string")
// }

// func TestGetByUsername(t *testing.T) {
// 	repo := RepoMock{mock.Mock{}}
// 	service := NewService(&repo)

// 	var dataMock = models.User{Username: "hendra"}

// 	//repo.Mock.On()
// 	repo.mock.On("FindByUsername").Return(&dataMock, nil)

// 	data := service.GetByUsername("hendra")

// 	var expectName = "hendra"
// 	//pecah data respone
// 	result := data.Data.(string)

// 	assert.Equal(t, expectName, result, "name is not hendra")

// }
