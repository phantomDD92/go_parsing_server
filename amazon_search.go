package main

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type AmazonAdsInfo struct {
	ASIN           string  `json:"asin"`
	HasPrime       bool    `json:"has_prime"`
	Image          string  `json:"image"`
	IsAmazonChoice bool    `json:"is_amazon_choice"`
	IsBestSeller   bool    `json:"is_best_seller"`
	IsLimitedDeal  bool    `json:"is_limited_deal"`
	Name           string  `json:"name"`
	Stars          float64 `json:"stars"`
	TotalReviews   int     `json:"total_reviews"`
	Type           string  `json:"type"`
	URL            string  `json:"url"`
}

type AmazonSearchInfo struct {
	ASIN           string  `json:"asin"`
	HasPrime       bool    `json:"has_prime"`
	Image          string  `json:"image"`
	IsAmazonChoice bool    `json:"is_amazon_choice"`
	IsBestSeller   bool    `json:"is_best_seller"`
	IsLimitedDeal  bool    `json:"is_limited_deal"`
	Name           string  `json:"name"`
	Position       int     `json:"position"`
	Price          float64 `json:"price"`
	PriceString    string  `json:"price_string"`
	PriceSymbol    string  `json:"price_symbol"`
	Stars          float64 `json:"stars"`
	TotalReviews   int     `json:"total_reviews"`
	Type           string  `json:"type"`
	URL            string  `json:"url"`
}

type AmazonSearchData struct {
	Ads              []AmazonAdsInfo    `json:"ads"`
	ExploreMoreItems []interface{}      `json:"explore_more_items"`
	NextPages        []string           `json:"next_pages"`
	Results          []AmazonSearchInfo `json:"results"`
}

type AmazonSearchResult struct {
	Data   AmazonSearchData `json:"data"`
	Status string           `json:"status"`
	URL    string           `json:"url"`
}

func replacePage(path string, page int) string {
	pathsegs := strings.Split(path, "&")
	for i := 0; i < len(pathsegs); i++ {
		if len(pathsegs[i]) > 5 && pathsegs[i][:5] == "page=" {
			pathsegs[i] = "page=" + strconv.Itoa(page)
		} else if len(pathsegs[i]) > 4 && pathsegs[i][:4] == "ref=" {
			pathsegs[i] = "ref=sr_pg_" + strconv.Itoa(page)
		}
	}
	return strings.Join(pathsegs, "&")
}

func getAdsItem(s *goquery.Selection) AmazonAdsInfo {
	var record AmazonAdsInfo
	asin := s.AttrOr("data-asin", "")
	record.ASIN = asin
	if asin != "" {
		record.HasPrime = s.Find("span.s-prime").Length() > 0
		titleTag := s.Find("span.a-truncate-full").First()
		if titleTag.Length() > 0 {
			record.Name = normalizeText(titleTag.Text())
		}
		record.IsAmazonChoice = false
		record.IsLimitedDeal = false
		record.IsBestSeller = s.Find("span.sx-bestseller-component").Length() > 0
		rateTag := s.Find("span[data-rt]").First()
		if rateTag.Length() > 0 {
			starText := rateTag.AttrOr("data-rt", "")
			num, err := strconv.ParseFloat(starText, 64)
			if err == nil {
				record.Stars = num
			}
			reviewText := normalizeText(rateTag.Text())
			review, err := strconv.Atoi(strings.ReplaceAll(reviewText, ",", ""))
			if err == nil {
				record.TotalReviews = review
			}
		}
		record.Type = "top_stripe_ads"
	}
	return record
}

func getSearchItem(s *goquery.Selection, baseUrl string, pos int) AmazonSearchInfo {
	var record AmazonSearchInfo
	asin := s.AttrOr("data-asin", "")
	record.ASIN = asin
	if asin != "" {
		// has_prime
		record.HasPrime = s.Find("span.s-prime").Length() > 0
		imgTag := s.Find("img").First()
		// image
		if imgTag.Length() > 0 {
			record.Image = normalizeImage(imgTag.AttrOr("src", ""))
		}
		record.IsAmazonChoice = false
		record.IsLimitedDeal = false
		record.IsBestSeller = s.Find("span.sx-bestseller-component").Length() > 0
		titleTag := s.Find("h2").First()
		if titleTag.Length() > 0 {
			record.Name = normalizeText(titleTag.Text())
		}
		record.Position = pos
		priceTag := s.Find(".s-price-instructions-style").First()
		if priceTag.Length() > 0 {
			priceText := getPrice(normalizeText(priceTag.Text()))
			record.PriceString = priceText
			if priceText != "" {
				price, err := strconv.ParseFloat(priceText[1:], 64)
				if err == nil {
					record.PriceSymbol = priceText[:1]
					record.Price = price
				} else {
					price, err := strconv.ParseFloat(priceText[2:], 64)
					if err == nil {
						record.PriceSymbol = priceText[:2]
						record.Price = price
					}
				}
			}
		}
		rateTag := s.Find(".s-title-instructions-style").First().Next()
		if rateTag.Length() > 0 {
			starTag := rateTag.Find("span").First()
			if starTag.Length() > 0 {
				starText := strings.Split(normalizeText(starTag.Text()), " ")[0]
				num, err := strconv.ParseFloat(starText, 64)
				if err == nil {
					record.Stars = num
				}
			}
			reviewTag := starTag.Next()
			if reviewTag.Length() > 0 {
				reviewText := normalizeText(reviewTag.Text())
				review, err := strconv.Atoi(strings.ReplaceAll(reviewText, ",", ""))
				if err == nil {
					record.TotalReviews = review
				}
			}
		}
		record.Type = "search_product"
		urlTag := s.Find("h2").Find("a").First()
		if urlTag.Length() > 0 {
			record.URL = baseUrl + urlTag.AttrOr("href", "")
		}
	}
	return record
}

func parseSearch(doc *goquery.Document) AmazonSearchResult {
	var result AmazonSearchResult
	var data AmazonSearchData
	baseUrl := "https://www.amazon.com"
	pos := 1
	resultsTag := doc.Find(".s-search-results").First()
	resultsTag.ChildrenFiltered("div.s-result-item").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			s.Find("div[data-asin]").Each(func(i int, node *goquery.Selection) {
				ads := getAdsItem(node)
				data.Ads = append(data.Ads, ads)
			})
		} else {
			item := getSearchItem(s, baseUrl, pos)
			if item.ASIN != "" {
				data.Results = append(data.Results, item)
				pos += 1
			}
		}
	})
	paginationTag := doc.Find(".s-pagination-container").First()
	if paginationTag.Length() > 0 {
		lastNumStr := normalizeText(paginationTag.Find("span.s-pagination-item").Last().Text())
		num, err := strconv.Atoi(lastNumStr)
		if err != nil {
			num = 0
		}
		path := paginationTag.Find("a").First().AttrOr("href", "")
		for i := 1; i <= num; i++ {
			nextPage := baseUrl + replacePage(path, i)
			data.NextPages = append(data.NextPages, nextPage)
		}
	}
	result.Data = data
	result.Status = "parse_successful"
	return result
}

func isSearchPage(doc *goquery.Document) bool {
	var mainElement = doc.Find("#search")
	return mainElement.Length() > 0
}
