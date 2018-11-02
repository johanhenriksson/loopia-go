package loopia

import (
	"github.com/kolo/xmlrpc"
)

// APIURL - where API is hosted
const APIURL string = "https://api.loopia.se/RPCSERV"

// API Struct to store runtime info
type API struct {
	Username string
	Password string
	RCPEndpoint string
	CustomerNumber string
}

// XMLRPCClient to interact with Loopia XMLRPC
func (api *API) XMLRPCClient() *xmlrpc.Client {
	client, _ := xmlrpc.NewClient(api.RCPEndpoint, nil)
	return client
}

// New returns a loopia.API instance
func New(username string, password string) (*API, error) {
	api := &API{
		RCPEndpoint: APIURL,
	}

	api.Username = username
	api.Password = password
	api.CustomerNumber = ""

	return api, nil
}
