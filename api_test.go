package gotrue_go

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const GOTRUE_JWT_SECRET = "37c304f8-51aa-419a-a1af-06154e63707a"

type AMREntry struct {
	Method    string `json:"method"`
	Timestamp int64  `json:"timestamp"`
}

type GoTrueClaims struct {
	jwt.StandardClaims
	Email                         string                 `json:"email,omitempty"`
	Phone                         string                 `json:"phone,omitempty"`
	AppMetaData                   map[string]interface{} `json:"app_metadata,omitempty"`
	UserMetaData                  map[string]interface{} `json:"user_metadata,omitempty"`
	Role                          string                 `json:"role"`
	AuthenticatorAssuranceLevel   string                 `json:"aal,omitempty"`
	AuthenticationMethodReference []AMREntry             `json:"amr,omitempty"`
	SessionId                     string                 `json:"session_id,omitempty"`
}

func NewTestServiceRoleClient(t *testing.T) *Client {
	claims := &GoTrueClaims{
		Role: "service_role",
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(GOTRUE_JWT_SECRET))
	if err != nil {
		assert.NoError(t, err)
	}
	return NewClient(token, "http://localhost:9999")
}

func NewTestAdminClient(t *testing.T) *Client {
	claims := &GoTrueClaims{
		StandardClaims: jwt.StandardClaims{
			Subject: "1234567890",
		},
		Role: "supabase_admin",
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(GOTRUE_JWT_SECRET))
	if err != nil {
		assert.NoError(t, err)
	}
	return NewClient(token, "http://localhost:9999")
}

func NewUnauthenticatedTestClient(t *testing.T) *Client {
	claims := &GoTrueClaims{
		StandardClaims: jwt.StandardClaims{
			Subject: "1234567890",
		},
		Role: "supabase_admin",
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("fooBar"))
	if err != nil {
		assert.NoError(t, err)
	}
	return NewClient(token, "http://localhost:9999")
}

func mockUpdateAdminUser() UpdateAdminUser {
	return customMockUpdateAdminUser(func(u UpdateAdminUser) UpdateAdminUser {
		return u
	})
}

func customMockUpdateAdminUser(options func(u UpdateAdminUser) UpdateAdminUser) UpdateAdminUser {
	password := gofakeit.Password(true, true, true, true, false, 32)
	user := UpdateAdminUser{
		Email:    gofakeit.Email(),
		Phone:    gofakeit.Phone(),
		Password: &password,
		UserMetaData: map[string]interface{}{
			"firstName": gofakeit.FirstName(),
			"lastName":  gofakeit.LastName(),
		},
		EmailConfirm: true,
		PhoneConfirm: true,
	}
	return options(user)
}
