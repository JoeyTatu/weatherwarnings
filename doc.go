// Package mypackage provides a utility for retrieving weather warnings from a specified URL.
// It includes a function GetWarnings that returns a slice of Warning structs.
//
// Example usage:
// package main
//
//		import (
//
//	    	"bufio"
//			"fmt"
//			"log"
//
//			"github.com/JoeyTatu/weatherwarnings"
//
//		)
//
//	func main() {
//	    // Day should equal 'today', 'tomorrow' or 'dayAfterTomorrow'
//	    fmt.Print("Enter the day:\n> ")
//		scanner := bufio.NewScanner(os.Stdin)
//		scanner.Scan()
//		day := scanner.Text()
//
//		weatherWarnings, err := weatherwarnings.GetWarnings(day)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		// Print or use the warnings as needed
//		for _, w := range weatherWarnings {
//			fmt.Printf("Title: %s\n", w.Title)
//			fmt.Printf("Description: %s\n", w.Description)
//			fmt.Printf("Valid: %s\n", w.Valid)
//			fmt.Printf("Issued: %s\n", w.Issued)
//			fmt.Println("-----------")
//		}
//	}
package weatherwarnings
