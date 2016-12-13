package clearbit

import (
	"github.com/dghubble/sling"
	"net/http"
	"os"
	"time"
)

// Client is a Clearbit client for making Clearbit API requests.
type Client struct {
	sling *sling.Sling

	Autocomplete *AutocompleteService
	Person       *PersonService
	Company      *CompanyService
	Discovery    *DiscoveryService
	Prospector   *ProspectorService
	Reveal       *RevealService
}

// Config represents all the parameters available to configure a Clearbit
// client
type Config struct {
	apiKey     string
	httpClient *http.Client
	Timeout    time.Duration
}

// Option is an option passed to the NewClient function used to change
// the client configuration
type Option func(*Config)

// WithHTTPClient sets the optional http.Client we can use to make requests
func WithHTTPClient(httpClient *http.Client) Option {
	return func(config *Config) {
		config.httpClient = httpClient
	}
}

// WithAPIKey sets the Clearbit API key.
//
// When this is not provided we'll default to the `CLEARBIT_KEY` environment
// variable.
func WithAPIKey(apiKey string) func(*Config) {
	return func(config *Config) {
		config.apiKey = apiKey
	}
}

// WithTimeout sets the timeout in seconds directly
//
// This is just an easier way to set the timeout than directly setting it
// through the withHTTPClient option.
func WithTimeout(d time.Duration) func(*Config) {
	return func(config *Config) {
		config.Timeout = d
	}
}

// NewClient returns a new Client.
func NewClient(options ...Option) *Client {
	config := Config{
		apiKey:     os.Getenv("CLEARBIT_KEY"),
		httpClient: &http.Client{},
		Timeout:    10 * time.Second,
	}

	for _, option := range options {
		option(&config)
	}

	config.httpClient.Timeout = config.Timeout

	base := sling.New().Client(config.httpClient)
	base.SetBasicAuth(config.apiKey, "")

	return &Client{
		Autocomplete: newAutocompleteService(base.New()),
		Person:       newPersonService(base.New()),
		Company:      newCompanyService(base.New()),
		Discovery:    newDiscoveryService(base.New()),
		Prospector:   newProspectorService(base.New()),
		Reveal:       newRevealService(base.New()),
	}
}
