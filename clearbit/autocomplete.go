package clearbit

import (
	"github.com/dghubble/sling"
	"net/http"
)

const (
	autoCompleteBase = "https://autocomplete.clearbit.com"
)

type AutocompleteItem struct {
	Domain string `json:"domain"`
	Logo   string `json:"logo"`
	Name   string `json:"name"`
}

type AutocompleteSuggestParams struct {
	Query string `url:"query"`
}

type AutocompleteService struct {
	baseSling *sling.Sling
	sling     *sling.Sling
}

func newAutocompleteService(sling *sling.Sling) *AutocompleteService {
	return &AutocompleteService{
		baseSling: sling.New(),
		sling:     sling.Base(autoCompleteBase).Path("/v1/companies/").Set("Authorization", ""),
	}
}

func (s *AutocompleteService) Suggest(params AutocompleteSuggestParams) ([]AutocompleteItem, *http.Response, error) {
	items := new([]AutocompleteItem)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("suggest").QueryStruct(params).Receive(items, apiError)
	return *items, resp, relevantError(err, *apiError)
}
