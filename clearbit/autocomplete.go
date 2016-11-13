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

type autocompleteSuggestParams struct {
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

func (s *AutocompleteService) Suggest(query string) ([]AutocompleteItem, *http.Response, error) {
	params := &autocompleteSuggestParams{Query: query}
	items := new([]AutocompleteItem)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("suggest").QueryStruct(params).Receive(items, apiError)
	return *items, resp, relevantError(err, *apiError)
}
