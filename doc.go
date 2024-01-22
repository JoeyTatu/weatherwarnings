package weatherwarnings

// Package weatherwarnings provides a utility for retrieving weather warnings from a specified URL.
// It includes a function GetWarnings that returns a slice of Warning structs.
//
// Example usage:
// package main
//
// import (
//     "bufio"
//     "fmt"
//     "log"
//
//     "github.com/JoeyTatu/weatherwarnings"
// )
//
// func main() {
//     // Valid requests:
//     // weatherwarnings.GetWarnings("today")
//     // weatherwarnings.GetWarnings("tomorrow")
//     // weatherwarnings.GetWarnings("dayAfterTomorrow")
//     // weatherwarnings.GetWarnings() // defaults to "today"
//
//     weatherWarnings, err := weatherwarnings.GetWarnings("today")
//     if err != nil {
//         log.Fatal(err)
//     }
//
//     // Print or use the warnings as needed
//     for _, w := range weatherWarnings {
//         fmt.Printf("Title: %s\n", w.Title)
//         fmt.Printf("Description: %s\n", w.Description)
//         fmt.Printf("Valid: %s\n", w.Valid)
//         fmt.Printf("Issued: %s\n", w.Issued)
//         fmt.Println("-----------")
//     }
// }
