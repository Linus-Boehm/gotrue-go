package gotrue_go

import (
	jwt "github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type AMREntry struct {
	Method    string `json:"method"`
	Timestamp int64  `json:"timestamp"`
}

type GoTrueClaims struct {
	jwt.StandardClaims
	Email                         string                 `json:"email"`
	Phone                         string                 `json:"phone"`
	AppMetaData                   map[string]interface{} `json:"app_metadata"`
	UserMetaData                  map[string]interface{} `json:"user_metadata"`
	Role                          string                 `json:"role"`
	AuthenticatorAssuranceLevel   string                 `json:"aal,omitempty"`
	AuthenticationMethodReference []AMREntry             `json:"amr,omitempty"`
	SessionId                     string                 `json:"session_id,omitempty"`
}

func NewTestClient(t *testing.T) *Client {
	claims := &GoTrueClaims{
		Role: "supabase_admin",
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("testsecret"))
	if err != nil {
		assert.NoError(t, err)
	}
	return NewClient(token, "http://localhost:9999")
}

func TestAdminUsersApi_GetUsers(t *testing.T) {

	tests := []struct {
		name    string
		want    []User
		wantErr bool
	}{
		{
			name:    "happy path",
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewTestClient(t)
			got, err := c.AdminUsersApi.GetUser()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}
