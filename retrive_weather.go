package weatherwarnings

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Warning represents the information for each warning
type Warning struct {
	Title       string
	Description string
	Valid       string
	Issued      string
}

func GetWarnings() ([]Warning, error) {
	url := "https://www.met.ie/warnings/today/"

	// Make a GET request to the URL
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Parse the HTML document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	// Initialize a slice of Warning
	warnings := []Warning{}

	// Iterate over each h2 element
	doc.Find("h2").Each(func(i int, s *goquery.Selection) {
		// Get the text of the h2 element
		title := strings.TrimSpace(s.Text())

		// Skip certain titles
		if title == "Search" || title == "Connect" || title == "Download our apps" {
			return
		}

		// Create a new Warning instance
		warning := Warning{}

		// Set the Title from the text of the h2 element
		warning.Title = title

		// Find the next span with class "sr-only"
		statusSpan := s.Next().Find("span.sr-only")
		if statusSpan != nil {
			// Set the Description from the text of the span
			warning.Description = strings.TrimSpace(strings.ReplaceAll(statusSpan.Text(), "<br />", "\n"))
		}

		// Find the next strong element with class "sr-only"
		strongElement := s.Next().Find("strong.sr-only")
		if strongElement != nil {
			// Set the Description from the text of the strong element
			warning.Description = strings.TrimSpace(strings.ReplaceAll(strongElement.Text(), "<br />", "\n"))
		}

		// Find the next p elements
		s.Next().Find("p").Each(func(j int, p *goquery.Selection) {
			// Set the Valid and Issued fields based on the content of p elements
			text := strings.TrimSpace(strings.ReplaceAll(p.Text(), "<br />", "\n"))
			switch {
			case strings.Contains(text, "Valid"):
				warning.Valid = strings.TrimPrefix(text, "Valid: ")
			case strings.Contains(text, "Issued"):
				warning.Issued = strings.TrimPrefix(text, "Issued: ")
			}
		})

		// Append the warning to the slice
		warnings = append(warnings, warning)
	})

	// Build the result string
	var result string
	for _, w := range warnings {
		result += fmt.Sprintf("Title: %s\n", w.Title)
		result += fmt.Sprintf("Description: %s\n", w.Description)
		result += fmt.Sprintf("Valid: %s\n", w.Valid)
		result += fmt.Sprintf("Issued: %s\n", w.Issued)
		result += "-----------\n"
	}

	return warnings, nil
}
