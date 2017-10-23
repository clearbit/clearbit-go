package clearbit_test

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/clearbit/clearbit-go/clearbit"
)

func handleError(err error, resp *http.Response) {
	fmt.Printf("%#v\n%s\n", err, resp.Status)
}

func ExampleNewClient_manuallyConfiguringEverything_output() {
	var clearbitApiKey = os.Getenv("CLEARBIT_KEY")
	client := clearbit.NewClient(
		clearbit.WithHTTPClient(&http.Client{}),
		clearbit.WithAPIKey(clearbitApiKey),
		clearbit.WithTimeout(20*time.Second),
	)

	_, resp, _ := client.Discovery.Search(clearbit.DiscoverySearchParams{
		Query: "name:clearbit",
	})

	fmt.Println(resp.Status)

	// Output: 200 OK
}

func ExampleRiskService_Calculate_output() {
	client := clearbit.NewClient()
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
	client := clearbit.NewClient()
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
	client := clearbit.NewClient()
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

func ExampleProspectorService_Search_output() {
	client := clearbit.NewClient()
	results, resp, err := client.Prospector.Search(clearbit.ProspectorSearchParams{
		Domain: "clearbit.com",
	})

	if err == nil {
		fmt.Println(results[0].Email, resp.Status)
	} else {
		handleError(err, resp)
	}

	// Output: alex@clearbit.com 200 OK
}

func ExampleProspectorService_Search_withRoles_Output() {
	client := clearbit.NewClient()
	results, resp, err := client.Prospector.Search(clearbit.ProspectorSearchParams{
		Domain: "clearbit.com",
		Roles:  []string{"sales", "engineering"},
	})

	if err == nil {
		fmt.Println(len(results), resp.Status)
	} else {
		handleError(err, resp)
	}

	// Output: 5 200 OK
}

func ExampleCompanyService_Find_output() {
	client := clearbit.NewClient()
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
	client := clearbit.NewClient()
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
	client := clearbit.NewClient()
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
	client := clearbit.NewClient()
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
