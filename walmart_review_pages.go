package main

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type WrProduct struct {
	Name       string    `json:"name"`
	Url        string    `json:"url"`
	Type       string    `json:"type"`
	Seller     string    `json:"seller"`
	Categories []WStrMap `json:"categories"`
}

type WrMedia struct {
	MediaType string `json:"media_type"`
	Rating    int    `json:"rating"`
	Url       string `json:"url"`
	Caption   string `json:"caption"`
}

type WrReview struct {
	Title            string    `json:"title"`
	Text             string    `json:"text"`
	SubmissionTime   string    `json:"time"`
	Rating           int       `json:"rating"`
	NegativeFeedback int       `json:"negative_feedbacks"`
	PositiveFeedback int       `json:"positive_feedbacks"`
	UserName         string    `json:"username"`
	Badges           []string  `json:"badges"`
	Media            []WrMedia `json:"media"`
}

type WrFrequentMention struct {
	Name         string `json:"name"`
	Score        int    `json:"score"`
	SnippetCount int    `json:"snippet_count"`
}

type WrPagination struct {
	PageCount   int       `json:"page_count"`
	CurrentPage int       `json:"current_page"`
	CurrentSpan string    `json:"current_span"`
	PageLinks   []WStrMap `json:"page_links"`
}

type WrData struct {
	AvarageRating            float64             `json:"average_rating"`
	TotalReviewCount         int                 `json:"total_review_count"`
	ReviewWithTextCount      int                 `json:"review_withtext_count"`
	AspectReviewCount        int                 `json:"aspect_review_count"`
	NegativeCount            int                 `json:"negative_count"`
	PositiveCount            int                 `json:"positive_count"`
	TotalMediaCount          int                 `json:"total_media_count"`
	PercentageFiveCount      int                 `json:"5_star_percent"`
	RatingValueFiveCount     int                 `json:"5_star_count"`
	ReviewWithTextFiveCount  int                 `json:"5_star_withtext_count"`
	PercentageFourCount      int                 `json:"4_star_percent"`
	RatingValueFourCount     int                 `json:"4_star_count"`
	ReviewWithTextFourCount  int                 `json:"4_star_withtext_count"`
	PercentageThreeCount     int                 `json:"3_star_percent"`
	RatingValueThreeCount    int                 `json:"3_star_count"`
	ReviewWithTextThreeCount int                 `json:"3_star_withtext_count"`
	PercentageTwoCount       int                 `json:"2_star_percent"`
	RatingValueTwoCount      int                 `json:"2_star_count"`
	ReviewWithTextTwoCount   int                 `json:"2_star_withtext_count"`
	PercentageOneCount       int                 `json:"1_star_percent"`
	RatingValueOneCount      int                 `json:"1_star_count"`
	ReviewWithTextOneCount   int                 `json:"1_star_withtext_count"`
	Reviews                  []WrReview          `json:"reviews"`
	TopNegativeReview        WrReview            `json:"top_negative_review"`
	TopPositiveReview        WrReview            `json:"top_positive_review"`
	Product                  WrProduct           `json:"product"`
	FrequentMetions          []WrFrequentMention `json:"frequent_mentions"`
	Pagination               WrPagination        `json:"pagination"`
}

type WrResult struct {
	Data   WrData `json:"data"`
	Status string `json:"status"`
	URL    string `json:"url"`
}

type WrrReview struct {
	ReviewId             string `json:"reviewId"`
	Rating               int    `json:"rating"`
	ReviewSubmissionTime string `json:"reviewSubmissionTime"`
	ReviewText           string `json:"reviewText"`
	ReviewTitle          string `json:"reviewTitle"`
	NegativeFeedback     int    `json:"negativeFeedback"`
	PositiveFeedback     int    `json:"positiveFeedback"`
	UserNickname         string `json:"userNickname"`
	Badges               []struct {
		Id         string `json:"id"`
		GlassBadge struct {
			Text string `json:"text"`
		} `json:"glassBadge"`
	} `json:"badges"`
	Media []struct {
		MediaType string `json:"mediaType"`
		NormalUrl string `json:"normalUrl"`
		Caption   string `json:"caption"`
		Rating    int    `json:"rating"`
	} `json:"media"`
}

type WrrPage struct {
	Num    int    `json:"num"`
	Active bool   `json:"active"`
	Url    string `json:"url"`
	Gap    bool   `json:"gap"`
}

type WrrData struct {
	Reviews struct {
		Aspects []struct {
			Name         string `json:"name"`
			Score        int    `json:"score"`
			SnippetCount int    `json:"snippetCount"`
		} `json:"aspects"`
		CustomerReviews                 []WrrReview `json:"customerReviews"`
		TopNegativeReview               WrrReview   `json:"topNegativeReview"`
		TopPositiveReview               WrrReview   `json:"topPositiveReview"`
		NegativeCount                   int         `json:"negativeCount"`
		PositiveCount                   int         `json:"positiveCount"`
		AverageRating                   float64     `json:"averageOverallRating"`
		PercentageFiveCount             int         `json:"percentageFiveCount"`
		PercentageFourCount             int         `json:"percentageFourCount"`
		PercentageThreeCount            int         `json:"percentageThreeCount"`
		PercentageTwoCount              int         `json:"percentageTwoCount"`
		PercentageOneCount              int         `json:"percentageOneCount"`
		FilteredReviewsCount            int         `json:"filteredReviewsCount"`
		RatingValueFiveCount            int         `json:"ratingValueFiveCount"`
		RatingValueFourCount            int         `json:"ratingValueFourCount"`
		RatingValueThreeCount           int         `json:"ratingValueThreeCount"`
		RatingValueTwoCount             int         `json:"ratingValueTwoCount"`
		RatingValueOneCount             int         `json:"ratingValueOneCount"`
		RecommendedPercentage           int         `json:"recommendedPercentage"`
		TotalMediaCount                 int         `json:"totalMediaCount"`
		TotalReviewCount                int         `json:"totalReviewCount"`
		ReviewsWithTextCount            int         `json:"reviewsWithTextCount"`
		ReviewsWithTextRatingOneCount   int         `json:"reviewsWithTextRatingOneCount"`
		ReviewsWithTextRatingTwoCount   int         `json:"reviewsWithTextRatingTwoCount"`
		ReviewsWithTextRatingThreeCount int         `json:"reviewsWithTextRatingThreeCount"`
		ReviewsWithTextRatingFourCount  int         `json:"reviewsWithTextRatingFourCount"`
		ReviewsWithTextRatingFiveCount  int         `json:"reviewsWithTextRatingFiveCount"`
		AspectReviewsCount              int         `json:"aspectReviewsCount"`
		Pagination                      struct {
			CurrentSpan string    `json:"currentSpan"`
			Prev        WrrPage   `json:"prev"`
			Next        WrrPage   `json:"next"`
			Pages       []WrrPage `json:"pages"`
			Total       int       `json:"total"`
		} `json:"pagination"`
	} `json:"reviews"`

	Product struct {
		Name         string `json:"name"`
		CanonicalUrl string `json:"canonicalUrl"`
		SellerName   string `json:"sellerName"`
		Category     struct {
			Path []struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"path"`
		} `json:"category"`
		Type string `json:"typee"`
	} `json:"product"`
}

type WrrResult struct {
	Props struct {
		PageProps struct {
			InitialData struct {
				Data WrrData `json:"data"`
			} `json:"initialData"`
		} `json:"pageProps"`
	} `json:"props"`
	Page          string  `json:"page"`
	Query         WStrMap `json:"query"`
	RuntimeConfig struct {
		Host struct {
			Wmt string `json:"wmt"`
		} `json:"host"`
	} `json:"runtimeConfig"`
}

func WR_MakeUrl(baseUrl string, path string, params map[string]string) string {
	// Create a Values instance to encode the query parameters
	mainUrl := strings.ReplaceAll(baseUrl+path, "[id]", params["id"])
	values := url.Values{}
	for key, value := range params {
		if key != "id" {
			values.Add(key, value)
		}
	}
	// Construct the query string
	queryString := values.Encode()
	if len(values) > 0 {
		return mainUrl + "?" + queryString
	}
	return mainUrl
}

func WR_MakePageUrl(mainUrl string, page int) string {
	pageUrl := mainUrl
	if page != 1 {
		if strings.Contains(mainUrl, "?") {
			pageUrl += "&page=" + strconv.Itoa(page)
		} else {
			pageUrl += "?page=" + strconv.Itoa(page)
		}
	}
	return pageUrl
}

func WR_ParsePagination(raw WrrResult, mainUrl string) WrPagination {
	var pagination WrPagination
	rawPagination := raw.Props.PageProps.InitialData.Data.Reviews.Pagination
	pagination.CurrentSpan = rawPagination.CurrentSpan
	for idx, item := range rawPagination.Pages {
		if !item.Gap && item.Num != 0 {
			page := make(WStrMap)
			page["label"] = strconv.Itoa(item.Num)
			page["url"] = WR_MakePageUrl(mainUrl, item.Num)
			pagination.PageLinks = append(pagination.PageLinks, page)
			if item.Active {
				pagination.CurrentPage = item.Num
			}
			if idx == len(rawPagination.Pages)-1 {
				pagination.PageCount = item.Num
			}
		}
	}
	return pagination
}

func WR_ParseProduct(rawData WrrData, baseUrl string) WrProduct {
	var product WrProduct
	rawProduct := rawData.Product
	product.Name = rawProduct.Name
	product.Type = rawProduct.Type
	product.Seller = rawProduct.SellerName
	product.Url = baseUrl + rawProduct.CanonicalUrl
	for _, item := range rawProduct.Category.Path {
		cat := make(WStrMap)
		cat["name"] = item.Name
		cat["url"] = baseUrl + item.Url
		product.Categories = append(product.Categories, cat)
	}
	return product
}

func WR_ParseReview(rawReview WrrReview) WrReview {
	var review WrReview
	review.Text = rawReview.ReviewText
	review.Title = rawReview.ReviewTitle
	for _, item := range rawReview.Badges {
		review.Badges = append(review.Badges, item.Id)
	}
	review.NegativeFeedback = rawReview.NegativeFeedback
	review.PositiveFeedback = rawReview.PositiveFeedback
	review.Rating = rawReview.Rating
	review.SubmissionTime = rawReview.ReviewSubmissionTime
	review.UserName = rawReview.UserNickname
	for _, item := range rawReview.Media {
		var media WrMedia
		media.MediaType = item.MediaType
		media.Caption = item.Caption
		media.Rating = item.Rating
		media.Url = item.NormalUrl
		review.Media = append(review.Media, media)
	}
	return review
}

func WR_ParseData(rawData WrrData, baseUrl string) WrData {
	var data WrData
	data.AvarageRating = rawData.Reviews.AverageRating
	data.AspectReviewCount = rawData.Reviews.AspectReviewsCount
	for _, item := range rawData.Reviews.Aspects {
		var mention WrFrequentMention
		mention.Name = item.Name
		mention.Score = item.Score
		mention.SnippetCount = item.SnippetCount
		data.FrequentMetions = append(data.FrequentMetions, mention)
	}
	data.NegativeCount = rawData.Reviews.NegativeCount
	data.PercentageFiveCount = rawData.Reviews.PercentageFiveCount
	data.PercentageFourCount = rawData.Reviews.PercentageFourCount
	data.PercentageThreeCount = rawData.Reviews.PercentageThreeCount
	data.PercentageTwoCount = rawData.Reviews.PercentageTwoCount
	data.PercentageOneCount = rawData.Reviews.PercentageOneCount
	data.PositiveCount = rawData.Reviews.PositiveCount
	data.RatingValueFiveCount = rawData.Reviews.RatingValueFiveCount
	data.RatingValueFourCount = rawData.Reviews.RatingValueFourCount
	data.RatingValueThreeCount = rawData.Reviews.RatingValueThreeCount
	data.RatingValueTwoCount = rawData.Reviews.RatingValueTwoCount
	data.RatingValueOneCount = rawData.Reviews.RatingValueOneCount
	data.ReviewWithTextCount = rawData.Reviews.ReviewsWithTextCount
	data.ReviewWithTextFiveCount = rawData.Reviews.ReviewsWithTextRatingFiveCount
	data.ReviewWithTextFourCount = rawData.Reviews.ReviewsWithTextRatingFourCount
	data.ReviewWithTextThreeCount = rawData.Reviews.ReviewsWithTextRatingThreeCount
	data.ReviewWithTextTwoCount = rawData.Reviews.ReviewsWithTextRatingTwoCount
	data.ReviewWithTextOneCount = rawData.Reviews.ReviewsWithTextRatingOneCount
	data.TotalMediaCount = rawData.Reviews.TotalMediaCount
	data.TotalReviewCount = rawData.Reviews.TotalReviewCount
	for _, item := range rawData.Reviews.CustomerReviews {
		review := WR_ParseReview(item)
		data.Reviews = append(data.Reviews, review)
	}
	data.TopNegativeReview = WR_ParseReview(rawData.Reviews.TopNegativeReview)
	data.TopPositiveReview = WR_ParseReview(rawData.Reviews.TopPositiveReview)
	data.Product = WR_ParseProduct(rawData, baseUrl)
	return data
}

func Walmart_ReviewPageScraper(jsonTag *goquery.Selection) WrResult {
	var raw WrrResult
	var result WrResult
	baseUrl := "https://www.walmart.com"
	// dataTag := doc.Find("script#__NEXT_DATA__").First()
	// if dataTag.Length() > 0 {
	json.Unmarshal([]byte(jsonTag.Text()), &raw)
	baseUrl = raw.RuntimeConfig.Host.Wmt
	result.Data = WR_ParseData(raw.Props.PageProps.InitialData.Data, baseUrl)
	result.Data.Pagination = WR_ParsePagination(raw, baseUrl)
	result.Status = "parse_successful"
	// reviewUrl := strings.ReplaceAll(baseUrl+raw.Page, "[id]", raw.Query["id"])
	result.URL = WR_MakeUrl(baseUrl, raw.Page, raw.Query)
	return result
}

func Walmart_IsReviewPage(jsonTag *goquery.Selection) bool {
	var page WRawResult
	json.Unmarshal([]byte(jsonTag.Text()), &page)
	return page.Page != "" && strings.Split(page.Page, "/")[1] == "reviews"
}
