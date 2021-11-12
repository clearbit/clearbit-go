package clearbit_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/clearbit/clearbit-go/clearbit"
)

func handleError(err error, resp *http.Response) {
	fmt.Printf("%#v\n%s\n", err, resp.Status)
}

func mockClearbitServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// mock combined response
		if strings.Contains(r.URL.Path, "/v2/combined/find") {
			time.Sleep(5 * time.Second)
			_, _ = w.Write([]byte(`{
				"person": {
					"name": {
						"fullName": "Alex MacCaw"
					}
				},
				"company": {
					"name": "Clearbit"
				}
			  }`))
			return
		}

		// mock person response
		if strings.Contains(r.URL.Path, "/v2/people/find") {
			time.Sleep(5 * time.Second)
			_, _ = w.Write([]byte(`{
				"name": {
					"fullName": "Alex MacCaw"
				}
			  }`))
			return
		}

		// mock discovery response
		if strings.Contains(r.URL.Path, "/v1/companies/search") {
			time.Sleep(5 * time.Second)
			_, _ = w.Write([]byte(`{
				"results": [
					{
						"domain": "clearbit.com"
					}
				]
			  }`))
			return
		}

		// mock company response
		if strings.Contains(r.URL.Path, "/v2/companies/find") {
			time.Sleep(5 * time.Second)
			_, _ = w.Write([]byte(`{
				"name": "Clearbit"
			  }`))
			return
		}

		if strings.Contains(r.URL.Path, "/v1/people/search") {
			// mock prospector with roles param response
			if _, ok := r.URL.Query()["roles[]"]; ok {
				time.Sleep(5 * time.Second)
				_, _ = w.Write([]byte(`{
					"results": [
						{"role": "sales"},
						{"role": "sales"},
						{"role": "engineering"},
						{"role": "engineering"},
						{"role": "sales"}
					]
				  }`))
				return
			}

			// mock prospector without roles param response
			time.Sleep(5 * time.Second)
			_, _ = w.Write([]byte(`{
				"results": [
					{"email": "alex@clearbit.com"}
				]
			  }`))
			return
		}

		// mock autocomplete response
		if strings.Contains(r.URL.Path, "/v1/companies/suggest") {
			time.Sleep(5 * time.Second)
			_, _ = w.Write([]byte(`[
				{"domain": "clearbit.com"}
			  ]`))
			return
		}

		// mock name to domain response
		if strings.Contains(r.URL.Path, "/v1/domains/find") {
			time.Sleep(5 * time.Second)
			_, _ = w.Write([]byte(`{
				"domain": "uber.com"
			  }`))
			return
		}

		// mock reveal response
		if strings.Contains(r.URL.Path, "/v1/companies/find") {
			time.Sleep(5 * time.Second)
			_, _ = w.Write([]byte(`{
				"company": {
					"name": "Clearbit"
				}
			  }`))
			return
		}

		// mock risk response
		if strings.Contains(r.URL.Path, "/v1/calculate") {
			time.Sleep(5 * time.Second)
			_, _ = w.Write([]byte(`{
				"risk": {
					"score": 0
				}
			  }`))
			return
		}

		w.WriteHeader(http.StatusNotFound)
	}))
}

var clearbitServer = mockClearbitServer()

func ExampleNewClient_manuallyConfiguringEverything_output() {
	client := clearbit.NewClient(
		clearbit.WithHTTPClient(&http.Client{}),
		clearbit.WithTimeout(20*time.Second),
		clearbit.WithBaseURLs(map[string]string{"discovery": clearbitServer.URL}),
	)

	_, resp, _ := client.Discovery.Search(clearbit.DiscoverySearchParams{
		Query: "name:clearbit",
	})

	fmt.Println(resp.Status)

	// Output: 200 OK
}

func ExampleRiskService_Calculate_output() {
	client := clearbit.NewClient(clearbit.WithBaseURLs(map[string]string{"risk": clearbitServer.URL}))
	results, resp, err := client.Risk.Calculate(clearbit.RiskCalculateParams{
		Email: "alex@clearbit.com",
		Name:  "Alex MacCaw",
		IP:    "127.0.0.1",
	})

	if err == nil {
		fmt.Println(results.Risk.Score, resp.Status)
	} else {
		handleError(err, resp)
	}

	// Output: 0 200 OK
}

func ExampleRevealService_Find_output() {
	client := clearbit.NewClient(clearbit.WithBaseURLs(map[string]string{"reveal": clearbitServer.URL}))
	results, resp, err := client.Reveal.Find(clearbit.RevealFindParams{
		IP: "104.193.168.24",
	})

	if err == nil {
		fmt.Println(results.Company.Name, resp.Status)
	} else {
		handleError(err, resp)
	}

	// Output: Clearbit 200 OK
}

func ExampleAutocompleteService_Suggest_output() {
	client := clearbit.NewClient(clearbit.WithBaseURLs(map[string]string{"autocomplete": clearbitServer.URL}))
	results, resp, err := client.Autocomplete.Suggest(clearbit.AutocompleteSuggestParams{
		Query: "clearbit",
	})

	if err == nil {
		fmt.Println(results[0].Domain, resp.Status)
	} else {
		handleError(err, resp)
	}

	// Output: clearbit.com 200 OK
}

func ExampleNameToDomainService_Find_output() {
	client := clearbit.NewClient(clearbit.WithBaseURLs(map[string]string{"nameToDomain": clearbitServer.URL}))
	result, resp, err := client.NameToDomain.Find(clearbit.NameToDomainFindParams{
		Name: "Uber",
	})

	if err == nil {
		fmt.Println(result.Domain, resp.Status)
	} else {
		handleError(err, resp)
	}

	// Output: uber.com 200 OK
}

func ExampleProspectorService_Search_output() {
	client := clearbit.NewClient(clearbit.WithBaseURLs(map[string]string{"prospector": clearbitServer.URL}))
	results, resp, err := client.Prospector.Search(clearbit.ProspectorSearchParams{
		Domain: "clearbit.com",
	})

	if err == nil {
		fmt.Println(results.Results[0].Email, resp.Status)
	} else {
		handleError(err, resp)
	}

	// Output: alex@clearbit.com 200 OK
}

func ExampleProspectorService_Search_withRoles_Output() {
	client := clearbit.NewClient(clearbit.WithBaseURLs(map[string]string{"prospector": clearbitServer.URL}))
	results, resp, err := client.Prospector.Search(clearbit.ProspectorSearchParams{
		Domain: "clearbit.com",
		Roles:  []string{"sales", "engineering"},
	})

	if err == nil {
		fmt.Println(len(results.Results), resp.Status)
	} else {
		handleError(err, resp)
	}

	// Output: 5 200 OK
}

func ExampleCompanyService_Find_output() {
	client := clearbit.NewClient(clearbit.WithBaseURLs(map[string]string{"company": clearbitServer.URL}))
	results, resp, err := client.Company.Find(clearbit.CompanyFindParams{
		Domain: "clearbit.com",
	})

	if err == nil {
		fmt.Println(results.Name, resp.Status)
	} else {
		handleError(err, resp)
	}

	// Output: Clearbit 200 OK
}

func ExamplePersonService_Find_output() {
	client := clearbit.NewClient(clearbit.WithBaseURLs(map[string]string{"person": clearbitServer.URL}))
	results, resp, err := client.Person.Find(clearbit.PersonFindParams{
		Email: "alex@clearbit.com",
	})

	if err == nil {
		fmt.Println(results.Name.FullName, resp.Status)
	} else {
		handleError(err, resp)
	}

	// Output: Alex MacCaw 200 OK
}

func ExamplePersonService_FindCombined_output() {
	client := clearbit.NewClient(clearbit.WithBaseURLs(map[string]string{"person": clearbitServer.URL}))
	results, resp, err := client.Person.FindCombined(clearbit.PersonFindParams{
		Email: "alex@clearbit.com",
	})

	if err == nil {
		fmt.Println(results.Person.Name.FullName, results.Company.Name, resp.Status)
	} else {
		handleError(err, resp)
	}

	// Output: Alex MacCaw Clearbit 200 OK
}

func ExampleDiscoveryService_Search_output() {
	client := clearbit.NewClient(clearbit.WithBaseURLs(map[string]string{"discovery": clearbitServer.URL}))
	results, resp, err := client.Discovery.Search(clearbit.DiscoverySearchParams{
		Query: "name:clearbit",
	})

	if err == nil {
		fmt.Println(results.Results[0].Domain, resp.Status)
	} else {
		handleError(err, resp)
	}

	// Output: clearbit.com 200 OK
}
