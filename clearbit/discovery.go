package clearbit

import (
	"github.com/dghubble/sling"
	"net/http"
)

const (
	discoveryBase = "https://discovery.clearbit.com"
)

type DiscoverySearchParams struct {
	Page     int    `url:"page,omitempty"`
	PageSize int    `url:"page_size,omitempty"`
	Limit    int    `url:"limit,omitempty"`
	Sort     int    `url:"sort,omitempty"`
	Query    string `url:"query,omitempty"`
}

type DiscoveryResults struct {
	Total   int       `json:"total"`
	Page    int       `json:"page"`
	Results []Company `json:"results"`
}

type DiscoveryService struct {
	baseSling *sling.Sling
	sling     *sling.Sling
}

func newDiscoveryService(sling *sling.Sling) *DiscoveryService {
	return &DiscoveryService{
		baseSling: sling.New(),
		sling:     sling.Base(discoveryBase).Path("/v1/companies/"),
	}
}

func (s *DiscoveryService) Search(params DiscoverySearchParams) (*DiscoveryResults, *http.Response, error) {
	item := new(DiscoveryResults)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("search").QueryStruct(params).Receive(item, apiError)
	return item, resp, relevantError(err, *apiError)
}
