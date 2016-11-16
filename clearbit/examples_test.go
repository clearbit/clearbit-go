package clearbit_test

import (
	"fmt"
	"os"

	"github.com/clearbit/clearbit-go/clearbit"
)

func ExampleRevealService_Find_output() {
	client := clearbit.NewClient(nil, os.Getenv("CLEARBIT_KEY"))
	results, resp, err := client.Reveal.Find(clearbit.RevealFindParams{
		IP: "104.193.168.24",
	})

	if err == nil {
		fmt.Println(results, resp.Status)
	}

	// Output: PENDING 200 OK
}

func ExampleAutocompleteService_Suggest_output() {
	client := clearbit.NewClient(nil, os.Getenv("CLEARBIT_KEY"))
	results, resp, err := client.Autocomplete.Suggest(clearbit.AutocompleteSuggestParams{
		Query: "clearbit",
	})

	if err == nil {
		fmt.Println(results[0].Domain, resp.Status)
	}

	// Output: clearbit.com 200 OK
}

func ExampleProspectorService_Search_output() {
	client := clearbit.NewClient(nil, os.Getenv("CLEARBIT_KEY"))
	results, resp, err := client.Prospector.Search(clearbit.ProspectorSearchParams{
		Domain: "clearbit.com",
	})

	if err == nil {
		fmt.Println(results[0].Email, resp.Status)
	}

	// Output: chris@clearbit.com 200 OK
}

func ExampleCompanyService_Find_output() {
	client := clearbit.NewClient(nil, os.Getenv("CLEARBIT_KEY"))
	results, resp, err := client.Company.Find(clearbit.CompanyFindParams{
		Domain: "clearbit.com",
	})

	if err == nil {
		fmt.Println(results.Name, resp.Status)
	}

	// Output: Clearbit 200 OK
}

func ExamplePersonService_Find_output() {
	client := clearbit.NewClient(nil, os.Getenv("CLEARBIT_KEY"))
	results, resp, err := client.Person.Find(clearbit.PersonFindParams{
		Email: "alex@clearbit.com",
	})

	if err == nil {
		fmt.Println(results.Name.FullName, resp.Status)
	}

	// Output: Alex MacCaw 200 OK
}

func ExamplePersonService_FindCombined_output() {
	client := clearbit.NewClient(nil, os.Getenv("CLEARBIT_KEY"))
	results, resp, err := client.Person.FindCombined(clearbit.PersonFindParams{
		Email: "alex@clearbit.com",
	})

	if err == nil {
		fmt.Println(results.Person.Name.FullName, results.Company.Name, resp.Status)
	}

	// Output: Alex MacCaw Clearbit 200 OK
}
