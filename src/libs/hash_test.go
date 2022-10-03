package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestHashPassword(t *testing.T) {

// 	hashPass := HashPassword(pass)

// 	assert.True(t, true, hassPass, "return must be 123")

// }
func TestCheckPassword(t *testing.T) {
	hassPass := CheckPassword("123", "123")

	assert.Equal(t, false, hassPass, "return must be 123")
	assert.True(t, true, hassPass)

}
