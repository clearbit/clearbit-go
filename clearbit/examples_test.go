package clearbit_test

import (
	"fmt"
	"net/http"
	"os"

	"github.com/clearbit/clearbit-go/clearbit"
)

//In this example we manually set the Clearbit key but notice that when
//WithApiKey is not called we will fallback to the CLEARBIT_KEY environment
//variable.
func ExampleNewClient_manuallyConfiguringEverything_output() {
	yourApiKey := os.Getenv("CLEARBIT_KEY")

	client := clearbit.NewClient(
		clearbit.WithHTTPClient(&http.Client{}),
		clearbit.WithAPIKey(yourApiKey),
	)

	_, resp, _ := client.Discovery.Search(clearbit.DiscoverySearchParams{
		Query: "name:clearbit",
	})

	fmt.Println(resp.Status)

	// Output: 200 OK
}

func ExampleRevealService_Find_output() {
	client := clearbit.NewClient()
	results, resp, err := client.Reveal.Find(clearbit.RevealFindParams{
		IP: "104.193.168.24",
	})

	if err == nil {
		fmt.Println(results.Company.Name, resp.Status)
	} else {
		fmt.Println(results, resp.Status)
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
		fmt.Println(err, resp.Status)
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
		fmt.Println(err, resp.Status)
	}

	// Output: chris@clearbit.com 200 OK
}

func ExampleCompanyService_Find_output() {
	client := clearbit.NewClient()
	results, resp, err := client.Company.Find(clearbit.CompanyFindParams{
		Domain: "clearbit.com",
	})

	if err == nil {
		fmt.Println(results.Name, resp.Status)
	} else {
		fmt.Println(err, resp.Status)
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
		fmt.Println(err, resp.Status)
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
		fmt.Println(err, resp.Status)
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
		fmt.Println(err, resp.Status)
	}

	// Output: clearbit.com 200 OK
}
