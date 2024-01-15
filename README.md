# Weather Warnings Scraper

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/JoeyTatu/weatherwarnings)
![GitHub](https://img.shields.io/github/license/JoeyTatu/weatherwarnings)

## Overview

`weatherwarningsscraper` is a Go package that provides a utility for retrieving weather warnings from Met Éireann. It includes a function `GetWarnings` that returns a slice of `Warning` structs.

## Installation

To use this package in your Go project, run:

```go get -u github.com/JoeyTatu/weatherwarningsscraper```

# Example usage:
```
package main

import (
	"fmt"
	"log"

	"github.com/JoeyTatu/weatherwarningsscraper"
)

func main() {
	weatherWarnings, err := weatherwarningsscraper.GetWarnings()
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
