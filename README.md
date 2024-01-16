# Weather Warnings Scraper

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/JoeyTatu/weatherwarnings)
![GitHub](https://img.shields.io/github/license/JoeyTatu/weatherwarnings)

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
	// Day should equal 'today', 'tomorrow' or 'dayAfterTomorrow'
	// If no day entered, defaults to 'today'
	fmt.Print("Enter the day:\n> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	day := scanner.Text()
	
	weatherWarnings, err := weatherwarnings.GetWarnings(day)
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
