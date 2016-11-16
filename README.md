#Clearbit Go Bindings

##Usage

To use one of the Clearbit APIs you'll first need to create a client by calling the `NewClient` function.
By default `NewClient` will use a new `http.Client` and will fetch the Clearbit API key from the `CLEARBIT_KEY` environment variable.

The Clearbit API key can be changed with:

```go
  client := clearbit.NewClient(clearbit.WithAPIKey("sk_1234567890123123"))
```

You can tap another `http.Client` with:

```go
  client := clearbit.NewClient(clearbit.WithHTTPClient(&http.Client{}))
```

Both can be combined and the order is not important.

Once the client is created you can use any of the Clearbit APIs

```go
	client.Autocomplete
	client.Company     
	client.Discovery   
	client.Person      
	client.Prospector  
	client.Reveal      
```

Example:

```go
  package main

  import (
      "fmt"
      "github.com/clearbit/clearbit-go/clearbit"
  )

  func main() {
      client := clearbit.NewClient(clearbit.WithAPIKey("sk_1234567890123123"))

      results, resp, err := client.Reveal.Find(clearbit.RevealFindParams{
            IP: "104.193.168.24",
      })

      if err != nil {}
        fmt.Println(results, resp)
      }
  }
```

Please see `examples_test.go` for more details.
