package gotrue_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdminUsersApi_CreateUser_HappyPath(t *testing.T) {
	c := NewTestAdminClient(t)
	updateUser := mockUpdateAdminUser()

	user, apiError, err := c.AdminUsersApi.CreateUser(updateUser)

	assert.NoError(t, err)
	assert.NoError(t, apiError)
	assert.Equal(t, updateUser.Email, user.Email)
	assert.Equal(t, updateUser.Role, user.Role)
	assert.Equal(t, updateUser.Aud, user.Aud)
}

func TestAdminUsersApi_CreateUser_UniqueEMail(t *testing.T) {
	c := NewTestAdminClient(t)
	updateUser := mockUpdateAdminUser()

	_, apiError, err := c.AdminUsersApi.CreateUser(updateUser)
	_, apiError2, err2 := c.AdminUsersApi.CreateUser(updateUser)

	assert.NoError(t, err)
	assert.Nil(t, apiError)
	assert.NoError(t, err2)
	assert.Equal(t, 422, apiError2.StatusCode)
	assert.Equal(t, "Email address already registered by another user", apiError2.Error())
}

func TestAdminUsersApi_CreateUser_WrongCredentials(t *testing.T) {
	c := NewUnauthenticatedTestClient(t)
	updateUser := mockUpdateAdminUser()

	_, apiError, err := c.AdminUsersApi.CreateUser(updateUser)
	assert.NoError(t, err)
	assert.Equal(t, 401, apiError.StatusCode)
	assert.Equal(t, "Invalid token: signature is invalid", apiError.Error())
}

func TestAdminUsersApi_GetUsers(t *testing.T) {
	c := NewTestServiceRoleClient(t)
	updateUser := mockUpdateAdminUser()

	wantUser, apiError, err := c.AdminUsersApi.CreateUser(updateUser)
	assert.NoError(t, err)
	assert.Nil(t, apiError)

	response, apiError2, err2 := c.AdminUsersApi.ListUsers()
	assert.NoError(t, err2)
	assert.Nil(t, apiError2)

	assert.NotNil(t, response)
	assert.GreaterOrEqual(t, len(response.Users), 1)
	found := false
	for _, user := range response.Users {
		if user.ID.String() != wantUser.ID.String() {
			continue
		}
		found = true
		assert.Equal(t, wantUser.Email, user.Email)
		assert.Equal(t, wantUser.Role, user.Role)
		assert.Equal(t, wantUser.Aud, user.Aud)
		assert.Equal(t, wantUser.UserMetaData, user.UserMetaData)
		assert.Equal(t, wantUser.Phone, user.Phone)
	}
	assert.True(t, found)
}
