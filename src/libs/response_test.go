package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRespone(t *testing.T) {
	res := Respone(nil, 200, false)

	assert.Nil(t, nil, res.Data, "retrun must be nil")
	assert.True(t, true, res.IsError, "IsError must be true")
	assert.Equal(t, 200, res.Code, "must be status code 200")
}
