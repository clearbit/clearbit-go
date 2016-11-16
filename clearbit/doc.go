/*
Package clearbit provides a client for using the Clearbit API.

Usage:

  package main

  import (
      "fmt"
      "github.com/clearbit/clearbit-go/clearbit"
      "os"
  )

  func main() {
      client := clearbit.NewClient(clearbit.SetAPIKey("sk_1234567890123123"))

      results, resp, _ := client.Reveal.Find(clearbit.RevealFindParams{
            IP: "104.193.168.24",
      })

      fmt.Println(results, resp)
  }

*/
package clearbit
