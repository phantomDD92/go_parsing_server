package main

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func normalizeText(text string) string {
	text = strings.ReplaceAll(text, "‎", "")
	re := regexp.MustCompile(`\s+`)
	replacedStr := re.ReplaceAllString(strings.TrimSpace(text), " ")
	newStr := strings.ReplaceAll(replacedStr, "\u0026", "&")
	return newStr
}

func normalizeImage(url string) string {
	if url[:15] != "https://m.media" {
		return ""
	}
	segs := strings.Split(url, "/")
	last := len(segs) - 1
	imageName := segs[last]
	nameSegs := strings.Split(imageName, ".")
	imageName = nameSegs[0] + "." + nameSegs[len(nameSegs)-1]
	segs[last] = imageName
	return strings.Join(segs, "/")
}

func getPrice(text string) string {
	for _, price := range strings.Split(text, " ") {
		pattern := `(.\d+(\.\d+)?)`
		regex := regexp.MustCompile(pattern)
		matches := regex.FindAllString(price, -1)
		if len(matches) > 0 {
			return matches[0]
		}
	}
	return ""
}

func getBaseUrl(url string) string {
	parts := strings.Split(url, "/")[:3]
	return strings.Join(parts, "/")
}

func normalizeUrl(url string) string {
	return strings.ReplaceAll(url, "\u0026", "&")
}

func extractText(s *goquery.Selection, splitter string) string {
	var allText string = ""
	if s.Get(0).Data != "span" && s.Children().Length() > 0 {
		s.Children().Each(func(index int, item *goquery.Selection) {
			allText += extractText(item, splitter)
		})
	} else {
		if s.Get(0).Type == html.TextNode || s.Get(0).Data == "p" || s.Get(0).Data == "span" || s.Get(0).Data == "h2" || s.Get(0).Data == "h3" || s.Get(0).Data == "h4" {
			allText += normalizeText(s.Text()) + splitter
		}
	}
	return allText
}

func convertTableToMap(s *goquery.Selection) map[string]string {
	dataMap := make(map[string]string)

	// Iterate over each table row
	s.Find("tbody tr").Each(func(i int, row *goquery.Selection) {
		// Extract the key and value from each row
		key := normalizeText(row.Find("th").Text())
		value := normalizeText(row.Find("td").Text())
		// Add the key-value pair to the map
		dataMap[key] = value
	})
	// Print the map
	return dataMap
}
