# Clearbit Go Bindings
[![Build Status](https://travis-ci.org/clearbit/clearbit-go.svg?branch=master)](https://travis-ci.org/clearbit/clearbit-go) [![GoDoc](https://godoc.org/github.com/clearbit/clearbit-go?status.svg)](https://godoc.org/github.com/clearbit/clearbit-go/clearbit)

Package clearbit provides a client for using the Clearbit API.

## Maintenance Status

This repository is currently not actively maintained. If you're looking to integrate with Clearbit's API we recommend looking at the HTTP requests available in our documentation at [clearbit.com/docs](https://clearbit.com/docs)

## Usage

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

If you use the httpClient just to set the timeout you can instead use WithTimeout:

```go
  client := clearbit.NewClient(clearbit.WithTimeout(20 * time.Second))
```

All options can be combined and the order is not important.

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

      if err != nil {
        fmt.Println(results, resp)
      }
  }
```

Please see [the examples](https://godoc.org/github.com/clearbit/clearbit-go/clearbit#pkg-examples) for more details.

## License

clearbit-go is copyright © 2016 Clearbit. It is free software, and may
be redistributed under the terms specified in the [`LICENSE`] file.

[`LICENSE`]: /LICENSE
