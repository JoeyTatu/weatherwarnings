package weatherwarnings

import (
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Warning struct {
	Title       string
	Description string
	Valid       string
	Issued      string
}

func getDayAfterTomorrow() (string, error) {
	currentTime := time.Now()
	dayAfterTomorrow := currentTime.Add(48 * time.Hour)
	nameOfDay := dayAfterTomorrow.Weekday().String()
	nameOfDay = strings.ToLower(nameOfDay)
	return nameOfDay, nil
}

func getURL(day string) (string, error) {
	switch day {
	case "today":
		return "https://www.met.ie/warnings-today.html", nil
	case "tomorrow":
		return "https://www.met.ie/warnings-tomorrow.html", nil
	case "dayAfterTomorrow":
		result, err := getDayAfterTomorrow()
		if err != nil {
			return "", err
		}
		return "https://www.met.ie/warnings-" + result + ".html", nil
	default:
		return "https://www.met.ie/warnings-today.html", nil
	}
}

func GetWarnings(day ...string) ([]Warning, error) {
	nameOfDay, err := getURL(day[0])
	if err != nil {
		return nil, err
	}

	response, err := http.Get(nameOfDay)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	warnings := []Warning{}

	doc.Find("h2").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Text())

		if title == "Search" || title == "Connect" || title == "Download our apps" {
			return
		}

		warning := Warning{}

		warning.Title = title

		statusSpan := s.Next().Find("span.sr-only")
		if statusSpan != nil {
			warning.Description = strings.TrimSpace(strings.ReplaceAll(statusSpan.Text(), "<br />", "\n"))
		}

		strongElement := s.Next().Find("strong.sr-only")
		if strongElement != nil {
			warning.Description = strings.TrimSpace(strings.ReplaceAll(strongElement.Text(), "<br />", "\n"))
		}

		s.Next().Find("p").Each(func(j int, p *goquery.Selection) {
			text := strings.TrimSpace(strings.ReplaceAll(p.Text(), "<br />", "\n"))
			switch {
			case strings.Contains(text, "Valid"):
				warning.Valid = strings.TrimPrefix(text, "Valid: ")
			case strings.Contains(text, "Issued"):
				warning.Issued = strings.TrimPrefix(text, "Issued: ")
			}
		})

		warnings = append(warnings, warning)
	})

	return warnings, nil
}
