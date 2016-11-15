package clearbit_test

import (
  "fmt"
  "os"

  "github.com/clearbit/clearbit-go/clearbit"
)

func ExampleRevealService_Find() {
  client := clearbit.NewClient(nil, os.Getenv("CLEARBIT_KEY"))
  results, resp, _ := client.Reveal.Find(clearbit.RevealFindParams{IP: "104.193.168.24"})
  fmt.Println(results, resp)
}

func ExampleAutocompleteService_Suggest() {
  client := clearbit.NewClient(nil, os.Getenv("CLEARBIT_KEY"))
  results, resp, _ := client.Autocomplete.Suggest(clearbit.AutocompleteSuggestParams{Query: "clearbit"})
  fmt.Println(results, resp)
}
