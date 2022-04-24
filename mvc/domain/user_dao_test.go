package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	// Initialization:

	// Execution:
	user, err := GetUser(0)

	// Validation:
	assert.Nil(t, user, "we are not expecting a user with user_id = 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 is not found", err.Message)

}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.NotNil(t, user, "we expect a user with user_id = 123")
	assert.Nil(t, err)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Fede", user.FirstName)
	assert.EqualValues(t, "Leon", user.LastName)
	assert.EqualValues(t, "tm1@gmail.com", user.Email)

}
