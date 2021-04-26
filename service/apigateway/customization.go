package apigateway

import (
	"github.com/ClearcodeHQ/aws-sdk-go/aws/client"
	"github.com/ClearcodeHQ/aws-sdk-go/aws/request"
)

func init() {
	initClient = func(c *client.Client) {
		c.Handlers.Build.PushBack(func(r *request.Request) {
			r.HTTPRequest.Header.Add("Accept", "application/json")
		})
	}
}
