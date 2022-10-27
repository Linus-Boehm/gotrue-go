package gotrue_go

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	serviceToken  string
	instanceUrl   string
	resty         *resty.Client
	AdminUsersApi *AdminUsersApi
}

func NewClient(serviceToken string, instanceUrl string) *Client {
	r := resty.New()
	client := &Client{
		serviceToken:  serviceToken,
		instanceUrl:   instanceUrl,
		resty:         r,
		AdminUsersApi: &AdminUsersApi{},
	}
	client.AdminUsersApi.client = client

	return client
}

func (c Client) PrepareRequest() *resty.Request {
	return c.resty.
		R().
		SetHeader("Accept", "application/json").
		SetAuthToken(c.serviceToken)
}

func (c Client) GetRequest(request *resty.Request, path string) (*resty.Response, error) {
	return request.Get(fmt.Sprintf("%s%s", c.instanceUrl, path))
}
