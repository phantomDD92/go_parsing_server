package main

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ReviewProduct struct {
	URL       string            `json:"url"`
	Name      string            `json:"name"`
	Brand     string            `json:"brand"`
	Price     float64           `json:"price"`
	Image     string            `json:"image"`
	Variation map[string]string `json:"variation"`
}

type ReviewInfo struct {
	Stars               int               `json:"stars"`
	Date                string            `json:"date"`
	VerifiedPurchase    bool              `json:"verified_purchase"`
	ManufacturerReplied bool              `json:"manufacturer_replied"`
	Username            string            `json:"username"`
	UserURL             string            `json:"userUrl"`
	Title               string            `json:"title"`
	Review              string            `json:"review"`
	ReviewURL           string            `json:"reviewUrl"`
	TotalFoundHelpful   int               `json:"total_found_helpful"`
	Images              []string          `json:"images"`
	Variation           map[string]string `json:"variation"`
	VideoURL            string            `json:"videoUrl"`
}

type ReviewData struct {
	AverageRating       float64       `json:"average_rating"`
	TotalReviews        int           `json:"total_reviews"`
	FiveStarRatings     int           `json:"5_star_ratings"`
	FiveStarPercentage  float64       `json:"5_star_percentage"`
	FourStarRatings     int           `json:"4_star_ratings"`
	FourStarPercentage  float64       `json:"4_star_percentage"`
	ThreeStarRatings    int           `json:"3_star_ratings"`
	ThreeStarPercentage float64       `json:"3_star_percentage"`
	TwoStarRatings      int           `json:"2_star_ratings"`
	TwoStarPercentage   float64       `json:"2_star_percentage"`
	OneStarRatings      int           `json:"1_star_ratings"`
	OneStarPercentage   float64       `json:"1_star_percentage"`
	Product             ReviewProduct `json:"product"`
	TopPositiveReview   ReviewInfo    `json:"top_positive_review"`
	TopCriticalReview   ReviewInfo    `json:"top_critical_review"`
	Reviews             []ReviewInfo  `json:"reviews"`
	Pagination          []string      `json:"pagination"`
}

type ReviewResult struct {
	Data   ReviewData `json:"data"`
	Status string     `json:"status"`
	URL    string     `json:"url"`
}

func extractVariation(variationTag *goquery.Selection) map[string]string {
	variationMap := make(map[string]string)
	variationText := normalizeTextWithReturn(variationTag.Text())
	variations := strings.Split(variationText, "\n")
	for _, variation := range variations {
		vals := strings.Split(variation, ":")
		if len(vals) > 1 {
			key := strings.TrimSpace(vals[0])
			value := strings.TrimSpace(vals[1])
			variationMap[key] = value
		}
	}
	return variationMap
}

func parseReviewSummary(reviewTag *goquery.Selection, baseUrl string) ReviewInfo {
	var info ReviewInfo
	profileTag := reviewTag.Find(".a-profile").First()
	if profileTag.Length() > 0 {
		info.Username = normalizeText(profileTag.Find("span").First().Text())
		info.UserURL = normalizeUrl(baseUrl, profileTag.AttrOr("href", ""))
	}
	reviewRatingTag := reviewTag.Find("i[data-hook='review-star-rating-view-point']").First()
	if reviewRatingTag.Length() > 0 {
		starText := normalizeText(reviewRatingTag.Find("span").First().Text())
		num, err := strconv.Atoi(strings.Split(starText, ".")[0])
		if err == nil {
			info.Stars = num
		}
	}
	reviewTitleTag := reviewTag.Find(".review-title").First()
	if reviewTitleTag.Length() > 0 {
		titleTexts := strings.Split(extractText(reviewTitleTag, "|"), "|")
		if len(titleTexts) >= 2 && titleTexts[len(titleTexts)-1] == "" {
			info.Title = titleTexts[len(titleTexts)-2]
		} else {
			info.Title = titleTexts[len(titleTexts)-1]
		}
	}
	reviewDateTag := reviewTag.Find(".review-date").First()
	if reviewDateTag.Length() > 0 {
		dateText := strings.Split(normalizeText(reviewDateTag.Text()), "on")[1]
		info.Date = strings.Trim(dateText, " ")
	}
	reviewContentTag := reviewDateTag.Parent().Next()
	if reviewContentTag.Length() > 0 {
		reviewText := extractText(reviewContentTag, " ")
		info.Review = reviewText
	}
	readMoreTag := reviewTag.Find(".readMore").Find("a").First()
	if readMoreTag.Length() > 0 {
		info.ReviewURL = normalizeUrl(baseUrl, readMoreTag.AttrOr("href", ""))
	}
	reviewVoteTag := reviewTag.Find(".review-votes").First()
	if reviewVoteTag.Length() > 0 {
		reviewVoteText := strings.Split(normalizeText(reviewVoteTag.Text()), " ")[0]
		if reviewVoteText == "One" {
			info.TotalFoundHelpful = 1
		} else {
			num, err := strconv.Atoi(strings.Split(reviewVoteText, " ")[0])
			if err == nil {
				info.TotalFoundHelpful = num
			}
		}
	}
	return info
}

func parseReviewInfo(reviewTag *goquery.Selection, baseUrl string) ReviewInfo {
	var info ReviewInfo
	profileTag := reviewTag.Find(".a-profile").First()
	if profileTag.Length() > 0 {
		info.Username = normalizeText(profileTag.Find("span").First().Text())
		info.UserURL = normalizeUrl(baseUrl, profileTag.AttrOr("href", ""))
	}
	reviewRatingTag := reviewTag.Find(".review-rating").First()
	if reviewRatingTag.Length() > 0 {
		starText := normalizeText(reviewRatingTag.Find("span").First().Text())
		num, err := strconv.Atoi(strings.Split(starText, ".")[0])
		if err == nil {
			info.Stars = num
		}
	}
	reviewTitleTag := reviewTag.Find(".review-title").First()
	if reviewTitleTag.Length() > 0 {
		titleTexts := strings.Split(extractText(reviewTitleTag, "|"), "|")
		if len(titleTexts) >= 2 && titleTexts[len(titleTexts)-1] == "" {
			info.Title = titleTexts[len(titleTexts)-2]
		} else {
			info.Title = titleTexts[len(titleTexts)-1]
		}
		info.ReviewURL = normalizeUrl(baseUrl, reviewTitleTag.AttrOr("href", ""))
	}
	reviewDateTag := reviewTag.Find(".review-date").First()
	if reviewDateTag.Length() > 0 {
		// println(normalizeText(reviewTag.Text()))
		dateText := strings.Split(normalizeText(reviewDateTag.Text()), "on")[1]
		info.Date = strings.Trim(dateText, " ")
	}
	reviewContentTag := reviewTag.Find(".review-text-content").First()
	if reviewContentTag.Length() > 0 {
		reviewText := extractText(reviewContentTag, " ")
		info.Review = reviewText
	}
	reviewVideoTag := reviewTag.Find("input.video-url").First()
	if reviewVideoTag.Length() > 0 {
		info.VideoURL = reviewVideoTag.AttrOr("value", "")
	}
	reviewVoteTag := reviewTag.Find("span[data-hook='helpful-vote-statement']").First()
	if reviewVoteTag.Length() > 0 {
		reviewVoteText := strings.Split(normalizeText(reviewVoteTag.Text()), " ")[0]
		if reviewVoteText == "One" {
			info.TotalFoundHelpful = 1
		} else {
			num, err := strconv.Atoi(strings.Split(reviewVoteText, " ")[0])
			if err == nil {
				info.TotalFoundHelpful = num
			}
		}
	}
	reviewImageTags := reviewTag.Find(".review-image-container").Find("img")
	if reviewImageTags.Length() > 0 {
		reviewImageTags.Each(func(i int, s *goquery.Selection) {
			info.Images = append(info.Images, normalizeImage(s.AttrOr("src", "")))
		})
	}
	// info.Variation = make(map[string]string)
	reviewVariationTag := reviewTag.Find("[data-hook='format-strip']").First()
	if reviewVariationTag.Length() > 0 {
		info.Variation = extractVariation(reviewVariationTag)
	}
	reviewBadgeTags := reviewTag.Find("[data-hook='avp-badge']")
	if reviewBadgeTags.Length() > 0 {
		reviewBadgeTags.Each(func(i int, s *goquery.Selection) {
			reviewBadgeText := normalizeText(s.Text())
			if reviewBadgeText == "Verified Purchase" {
				info.VerifiedPurchase = true
			}
		})
	}
	return info
}

func parseReview(doc *goquery.Document) ReviewResult {
	var result ReviewResult
	var data ReviewData
	baseUrl := "https://www.amazon.com"
	// url
	href, exists := doc.Find("link[rel='canonical']").Attr("href")
	if exists {
		result.URL = href
		baseUrl = getBaseUrl(href)
	}
	// parse review summary
	reviewSummaryTag := doc.Find("#cm_cr-product_info").Find(".reviewNumericalSummary").First()
	if reviewSummaryTag.Length() > 0 {
		averageRatingTag := reviewSummaryTag.Find(".averageStarRatingIconAndCount").First()
		if averageRatingTag.Length() > 0 {
			starText := strings.Split(normalizeText(averageRatingTag.Text()), " ")[0]
			starRate, err := strconv.ParseFloat(starText, 64)
			if err == nil {
				data.AverageRating = starRate
			}
		}
		totalReviewTag := reviewSummaryTag.Find(".averageStarRatingNumerical").First()
		if totalReviewTag.Length() > 0 {
			totalReviewText := normalizeText(totalReviewTag.Text())
			totalReviewText = strings.ReplaceAll(strings.Split(totalReviewText, " ")[0], ",", "")
			num, err := strconv.Atoi(totalReviewText)
			if err == nil {
				data.TotalReviews = num
			}
		}
	}
	starHistogramTags := doc.Find("#histogramTable").Find("tr")
	if starHistogramTags.Length() > 0 {
		starHistogramTags.Each(func(i int, s *goquery.Selection) {
			starText := strings.Split(s.AttrOr("aria-label", ""), "%")[0]
			starRate, err := strconv.ParseFloat(starText, 64)
			if err == nil {
				if i == 0 {
					data.FiveStarPercentage = starRate / 100
					data.FiveStarRatings = int(starRate * float64(data.TotalReviews) / 100)
				} else if i == 1 {
					data.FourStarPercentage = starRate / 100
					data.FourStarRatings = int(starRate * float64(data.TotalReviews) / 100)
				} else if i == 2 {
					data.ThreeStarPercentage = starRate / 100
					data.ThreeStarRatings = int(starRate * float64(data.TotalReviews) / 100)
				} else if i == 3 {
					data.TwoStarPercentage = starRate / 100
					data.TwoStarRatings = int(starRate * float64(data.TotalReviews) / 100)
				} else if i == 4 {
					data.OneStarPercentage = starRate / 100
					data.OneStarRatings = int(starRate * float64(data.TotalReviews) / 100)
				} else {

				}
			}
		})
	}
	productInfoTag := reviewSummaryTag.Next()
	if productInfoTag.Length() > 0 {
		productTitleTag := productInfoTag.Find(".product-title").First()
		if productTitleTag.Length() > 0 {
			data.Product.Name = normalizeText(productTitleTag.Text())
			productLinkTag := productTitleTag.Find("a").First()
			if productLinkTag.Length() > 0 {
				data.Product.URL = normalizeUrl(baseUrl, productLinkTag.AttrOr("href", ""))
			}
		}
		productBrandTag := productInfoTag.Find(".product-by-line").Find("a").First()
		if productBrandTag.Length() > 0 {
			data.Product.Brand = normalizeText(productBrandTag.Text())
		}
		productImageTag := doc.Find("img").First()
		if productImageTag.Length() > 0 {
			data.Product.Image = normalizeImage(productImageTag.AttrOr("src", ""))
		}
		productVariationTag := productInfoTag.Find(".product-variation-strip").Find("span").First()
		if productVariationTag.Length() > 0 {
			data.Product.Variation = extractVariation(productVariationTag)
		}
	}
	reviewViewpointTag := doc.Find("#cm_cr-rvw_summary-viewpoints").First()
	if reviewViewpointTag.Length() > 0 {
		positiveReviewTag := reviewViewpointTag.Find(".positive-review").First()
		if positiveReviewTag.Length() > 0 {
			data.TopPositiveReview = parseReviewSummary(positiveReviewTag, baseUrl)
		}
		criticalReviewTag := reviewViewpointTag.Find(".critical-review").First()
		if criticalReviewTag.Length() > 0 {
			data.TopCriticalReview = parseReviewSummary(criticalReviewTag, baseUrl)
		}
	}
	reviewListTag := doc.Find("#cm_cr-review_list").Find("div[data-hook='review']")
	if reviewListTag.Length() > 0 {
		reviewListTag.Each(func(i int, s *goquery.Selection) {
			reviewInfo := parseReviewInfo(s, baseUrl)
			data.Reviews = append(data.Reviews, reviewInfo)
		})
	}
	result.Data = data
	result.Status = "parse_successful"
	return result
}

func isReviewPage(doc *goquery.Document) bool {
	var mainElement = doc.Find("#cm_cr-product_info")
	return mainElement.Length() > 0
}
