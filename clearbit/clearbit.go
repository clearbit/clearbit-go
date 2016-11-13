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
	// Enrichment   *EnrichmentService
	// Discovery    *DiscoveryService
	// Logo         *LogoService
	// Prospector   *ProspectorService
	// Resource     *ResourceService
	// Reveal       *RevealService
	// Risk         *RiskService
	// Watchlist    *WatchlistService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client, apiKey string) *Client {
	base := sling.New().Client(httpClient)
	base.SetBasicAuth(apiKey, "")

	return &Client{
		apiKey:       apiKey,
		Autocomplete: newAutocompleteService(base.New()),
		Person:       newPersonService(base.New()),
		// Enrichment:   newEnrichmentService(base.New()),
		// Discovery:    newDiscoveryService(base.New()),
		// Logo:         newLogoService(base.New()),
		// Prospector:   newProspectorService(base.New()),
		// Resource:     newResourceService(base.New()),
		// Reveal:       newRevealService(base.New()),
		// Risk:         newRiskService(base.New()),
		// Watchlist:    newWatchlistService(base.New()),
	}
}

// Bool returns a new pointer to the given bool value.
func Bool(v bool) *bool {
	ptr := new(bool)
	*ptr = v
	return ptr
}

// Float returns a new pointer to the given float64 value.
func Float(v float64) *float64 {
	ptr := new(float64)
	*ptr = v
	return ptr
}
