package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Pierre Done", "pierre.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user.ID)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Pierre Done", user.Name)
	assert.Equal(t, "pierre.com", user.Email)
}

func TestUserValidatePassword(t *testing.T) {
	user, err := NewUser("Pierre Done", "pierre.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassWord("123456"))
	assert.False(t, user.ValidatePassWord("1234567"))
	assert.NotEqual(t, "123456", user.Password)
}
