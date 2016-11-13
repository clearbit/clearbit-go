package clearbit

import (
	"github.com/dghubble/sling"
	"net/http"
)

// Client is a Clearbit client for making Clearbit API requests.
type Client struct {
	apiKey string
	sling  *sling.Sling

	Autocomplete *AutocompleteService
	Person       *PersonService
	Company      *CompanyService
	Discovery    *DiscoveryService
	Prospector   *ProspectorService
	Reveal       *RevealService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client, apiKey string) *Client {
	base := sling.New().Client(httpClient)
	base.SetBasicAuth(apiKey, "")

	return &Client{
		apiKey:       apiKey,
		Autocomplete: newAutocompleteService(base.New()),
		Person:       newPersonService(base.New()),
		Company:      newCompanyService(base.New()),
		Discovery:    newDiscoveryService(base.New()),
		Prospector:   newProspectorService(base.New()),
		Reveal:       newRevealService(base.New()),
	}
}
