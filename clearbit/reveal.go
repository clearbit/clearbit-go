package clearbit

import (
	"github.com/dghubble/sling"
	"net/http"
)

const (
	revealBase = "https://reveal.clearbit.com"
)

type Reveal struct {
	IP    string `json:"ip"`
	Fuzzy bool   `json:"fuzzy"`

	Domain  string `json:"domain"`
	Company Company
}

type RevealFindParams struct {
	IP string `url:"ip,omitempty"`
}

type RevealService struct {
	baseSling *sling.Sling
	sling     *sling.Sling
}

func newRevealService(sling *sling.Sling) *RevealService {
	return &RevealService{
		baseSling: sling.New(),
		sling:     sling.Base(revealBase).Path("/v1/companies/"),
	}
}

// Find takes an IP address, and returns the company associated with that IP
func (s *RevealService) Find(params RevealFindParams) (*Reveal, *http.Response, error) {
	item := new(Reveal)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("find").QueryStruct(params).Receive(item, apiError)
	return item, resp, relevantError(err, *apiError)
}
