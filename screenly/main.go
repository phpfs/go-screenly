package screenly

import (
	"fmt"
	rclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	sclient "github.com/phpfs/go-screenly/screenly/client"
)

const headerAuthName = "Authorization"
const headerAuthFormat = "Token %v"

func NewScreenlyAPI(apiToken string) *sclient.ScreenlyAPI {
	t := rclient.New(sclient.DefaultHost, sclient.DefaultBasePath, sclient.DefaultSchemes)
	t.DefaultAuthentication = rclient.APIKeyAuth(headerAuthName, "header", fmt.Sprintf(headerAuthFormat, apiToken))
	return sclient.New(t, strfmt.Default)
}
