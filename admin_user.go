package gotrue_go

import (
	"github.com/Linus-Boehm/gotrue-go/schema"
	"github.com/go-resty/resty/v2"
	"github.com/gofrs/uuid"
)

type UpdateAdminUser struct {
	Aud          string                 `json:"aud"`
	Role         string                 `json:"role"`
	Email        string                 `json:"email"`
	Phone        string                 `json:"phone"`
	Password     *string                `json:"password"`
	EmailConfirm bool                   `json:"email_confirm"`
	PhoneConfirm bool                   `json:"phone_confirm"`
	UserMetaData map[string]interface{} `json:"user_metadata"`
	AppMetaData  map[string]interface{} `json:"app_metadata"`
	BanDuration  string                 `json:"ban_duration"`
}

type AdminUsersApi struct {
	client *Client
}

func (a AdminUsersApi) ListUsers() (*schema.ListUserResponse, *schema.APIError, error) {
	r := a.client.PrepareRequest()
	r.SetResult(schema.ListUserResponse{})

	response, err := a.client.GetRequest(r, "/admin/users")
	if err != nil {
		return nil, nil, err
	}
	listUserResponse, apiError := parseResult[schema.ListUserResponse](response)

	return listUserResponse, apiError, nil
}

func (a AdminUsersApi) CreateUser(updateUser UpdateAdminUser) (*schema.User, *schema.APIError, error) {
	r := a.client.PrepareRequest()

	r.SetResult(schema.User{})
	r.SetBody(updateUser)

	response, err := a.client.PostRequest(r, "/admin/users")
	if err != nil {
		return nil, nil, err
	}
	user, apiError := parseResult[schema.User](response)

	return user, apiError, nil
}

func (a AdminUsersApi) UpdateUserByID(id uuid.UUID, updateUser UpdateAdminUser) (*schema.User, *schema.APIError, error) {
	r := a.client.PrepareRequest()

	r.SetResult(schema.User{})
	r.SetBody(updateUser)

	response, err := a.client.PutRequestWithParam(r, "/admin/users", id)
	if err != nil {
		return nil, nil, err
	}
	user, apiError := parseResult[schema.User](response)

	return user, apiError, nil
}

func parseResult[T any](r *resty.Response) (*T, *schema.APIError) {
	var apiError *schema.APIError
	err := r.Error()
	if err != nil {
		apiError = err.(*schema.APIError)
	}

	result := r.Result()
	var resultT *T
	if result != nil {
		resultT = result.(*T)
	}

	return resultT, apiError
}
