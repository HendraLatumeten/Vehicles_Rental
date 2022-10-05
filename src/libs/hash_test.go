package libs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {

	hashPass, err := HashPassword("mypassword")
	fmt.Println(err)

	assert.True(t, true, hashPass, "return must be mypassword")
	assert.IsType(t, "string", hashPass, "return must be mypassword")

}

func TestCheckPassword(t *testing.T) {
	CheckPass := CheckPassword("123", "123")

	assert.Equal(t, false, CheckPass, "return must be 123")
	assert.True(t, true, CheckPass)

}
