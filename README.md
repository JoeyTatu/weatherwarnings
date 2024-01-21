# Weather Warnings Scraper

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/JoeyTatu/weatherwarnings)
![GitHub](https://img.shields.io/github/license/JoeyTatu/weatherwarnings)
[![Go Report Card](https://goreportcard.com/badge/github.com/JoeyTatu/weatherwarnings)](https://goreportcard.com/report/github.com/JoeyTatu/weatherwarnings)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-chi/render.svg)](https://pkg.go.dev/github.com/JoeyTatu/weatherwarnings)

## Overview

`weatherwarnings` is a Go package that provides a utility for retrieving weather warnings from **[Met Éireann](https://met.ie)**. It includes a function `GetWarnings` that returns a slice of `Warning` structs.

## Installation

To use this package in your Go project, run:

```go get -u github.com/JoeyTatu/weatherwarnings```

# Example usage:
```
package main

	import (
		"bufio"
		"fmt"
		"log"

		"github.com/JoeyTatu/weatherwarnings"
	)

func main() {
	
	// Valid requests:
	// weatherwarnings.GetWarnings("today")
	// weatherwarnings.GetWarnings("tomorrow")
	// weatherwarnings.GetWarnings("dayAfterTomorrow")
	// weatherwarnings.GetWarnings() // defaults to "today"

	weatherWarnings, err := weatherwarnings.GetWarnings("today")
	if err != nil {
		log.Fatal(err)
	}
	
	// Print or use the warnings as needed
	for _, w := range weatherWarnings {
		fmt.Printf("Title: %s\n", w.Title)
		fmt.Printf("Description: %s\n", w.Description)
		fmt.Printf("Valid: %s\n", w.Valid)
		fmt.Printf("Issued: %s\n", w.Issued)
		fmt.Println("-----------")
	}
}
```
 
