package main

import (
	"encoding/json"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type WStrMap map[string]string
type WpFulfillment struct {
	Type              string `json:"type"`
	AvailableStatus   bool   `json:"available_status"`
	Location          string `json:"location"`
	MaxOrderQuantity  int    `json:"max_order_quantity"`
	AvailableQuantity int    `json:"available_quantity"`
	Badge             string `json:"badge"`
	Date              string `json:"date"`
	Price             string `json:"price"`
}
type WpInfo struct {
	Categories      []WStrMap `json:"categories"`
	Name            string    `json:"name"`
	Brand           string    `json:"brand"`
	BrandUrl        string    `json:"brand_url"`
	Images          []string  `json:"images"`
	Thumbnail       string    `json:"thumbnail"`
	AverageRating   float64   `json:"average_rating"`
	NumberOfReviews int       `json:"number_of_reviews"`
	AvailableStatus string    `json:"available_status"`
	Model           string    `json:"model"`
	PriceInfo       struct {
		Price           float64 `json:"price"`
		PriceString     string  `json:"price_string"`
		WasPrice        string  `json:"was_price"`
		UnitPrice       string  `json:"unit_price"`
		ShipPrice       string  `json:"ship_price"`
		ListPrice       string  `json:"list_price"`
		ComparisonPrice string  `json:"comparison_price"`
		SavingsAmount   string  `json:"savings_amount"`
		PriceRange      string  `json:"price_range"`
		AdditionalFees  struct {
			DutyFee              string `json:"duty_fee"`
			ShippingAndImportFee string `json:"shipping_import_fee"`
			EstimatedTotalPrice  string `json:"estimated_total_price"`
		} `json:"additional_fees"`
		SecondaryOfferPrice string `json:"secondary_offer_price"`
	} `json:"price_info"`
	Fulfillments []WpFulfillment `json:"fulfillments"`
	EbtEligible  bool            `json:"is_ebt_eligible"`
	Badges       struct {
		Flags  []string `json:"flags"`
		Labels []string `json:"labels"`
		Groups []string `json:"groups"`
		Tags   []string `json:"tags"`
	} `json:"badges"`
	Seller struct {
		Name          string  `json:"name"`
		DisplayName   string  `json:"display_name"`
		StoreUrl      string  `json:"store_url"`
		ReviewCount   int     `json:"review_count"`
		AverageRating float64 `json:"average_rating"`
	} `json:"seller"`
	ReturnPolicy string `json:"return_policy"`
	Location     struct {
		PostalCode          string `json:"postal_code"`
		StateOrProvinceCode string `json:"state_code"`
		City                string `json:"city"`
	} `json:"location"`
	SalesUnit               string `json:"sales_unit"`
	ItemId                  string `json:"item_id"`
	OfferType               string `json:"offer_type"`
	TransactableOfferCount  string `json:"transactable_offer_count"`
	InteractiveProductVideo string `json:"interactive_product_video"`
}

type WpMedia struct {
	MediaType string `json:"media_type"`
	Rating    int    `json:"rating"`
	Url       string `json:"url"`
	Caption   string `json:"caption"`
}

type WpReview struct {
	ReviewId         string    `json:"id"`
	Rating           int       `json:"rating"`
	SubmissionTime   string    `json:"time"`
	Title            string    `json:"title"`
	Text             string    `json:"text"`
	NegativeFeedback int       `json:"negative_feedbacks"`
	PositiveFeedback int       `json:"positive_feedbacks"`
	UserName         int       `json:"username"`
	Badges           []string  `json:"badges"`
	Media            []WpMedia `json:"media"`
}

type WpAbort struct {
	NutritionInfo interface{} `json:"nutrition_information"`
	Details       struct {
		ShortDescription string `json:"short_description"`
		LongDescription  string `json:"long_description"`
	} `json:"product_details"`
	Specifications []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"specifications"`
	Warnings []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"warnings"`
	Directions []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"directions"`
	Highlights []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"highlights"`
	Ingredients []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"ingredients"`
	Warranty string    `json:"warranty"`
	Videos   []WStrMap `json:"videos"`
}

type WpReviewInfo struct {
	AverageOverallRating  float64    `json:"average_rating"`
	TotalReviewCount      int        `json:"total_review_count"`
	TotalMediaCount       int        `json:"total_media_count"`
	RatingValueFiveCount  int        `json:"5_star_rating"`
	PercentageFiveCount   int        `json:"5_star_percent"`
	RatingValueFourCount  int        `json:"4_star_rating"`
	PercentageFourCount   int        `json:"4_star_percent"`
	RatingValueThreeCount int        `json:"3_star_rating"`
	PercentageThreeCount  int        `json:"3_star_percent"`
	RatingValueTwoCount   int        `json:"2_star_rating"`
	PercentageTwoCount    int        `json:"2_star_percent"`
	RatingValueOneCount   int        `json:"1_star_rating"`
	PercentageOneCount    int        `json:"1_star_percent"`
	CustomerReviews       []WpReview `json:"customer_reviews"`
	TopNegativeReview     WpReview   `json:"top_negative_review"`
	TopPositiveReview     WpReview   `json:"top_positive_review"`
}
type WpData struct {
	Product       WpInfo       `json:"product"`
	About         WpAbort      `json:"about"`
	RelatedSearch []WStrMap    `json:"related_search"`
	Review        WpReviewInfo `json:"review_infomation"`
	RelatedPages  []WStrMap    `json:"related_pages"`
}

type WpResult struct {
	Data   WpData `json:"data"`
	Status string `json:"status"`
	URL    string `json:"url"`
}

type WprModule struct {
	Type    string `json:"type"`
	Configs struct {
		SeoItemRelmData struct {
			Relm []struct {
				Id   string `json:"id"`
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"relm"`
		} `json:"seoItemRelmData,omitempty"`
		RelatedSearch struct {
			RelatedSearch []struct {
				Title    string `json:"title"`
				ImageUrl string `json:"imageUrl"`
				Url      string `json:"url"`
			} `json:"relatedSearch"`
		} `json:"relatedSearch,omitempty"`
	} `json:"configs"`
}

type WprReview struct {
	ReviewId             string `json:"reviewId"`
	Rating               int    `json:"rating"`
	ReviewSubmissionTime string `json:"reviewSubmissionTime"`
	ReviewText           string `json:"reviewText"`
	ReviewTitle          string `json:"reviewTitle"`
	NegativeFeedback     int    `json:"negativeFeedback"`
	PositiveFeedback     int    `json:"positiveFeedback"`
	UserNickname         int    `json:"userNickname"`
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

type WprResult struct {
	Props struct {
		PageProps struct {
			InitialData struct {
				Data struct {
					Idml struct {
						ProductHighlights []struct {
							Name  string `json:"name"`
							Value string `json:"value"`
						} `json:"productHighlights"`
						DrugGuide interface{} `json:"drugGuide"`
						Warranty  struct {
							Information string `json:"information"`
						} `json:"warranty"`
						Ingredients struct {
							Ingredients []struct {
								Name  string `json:"name"`
								Value string `json:"value"`
							} `json:"ingredients"`
						} `json:"ingredients"`
						ShortDescription string `json:"shortDescription"`
						LongDescription  string `json:"longDescription"`
						Directions       []struct {
							Name  string `json:"name"`
							Value string `json:"value"`
						} `json:"directions"`
						Specifications []struct {
							Name  string `json:"name"`
							Value string `json:"value"`
						} `json:"specifications"`
						Warnings []struct {
							Name  string `json:"name"`
							Value string `json:"value"`
						} `json:"warnings"`
						NutritionFacts interface{} `json:"nutritionFacts"`
						Videos         []struct {
							Poster   string `json:"poster"`
							Title    string `json:"title"`
							Versions struct {
								Small string `json:"small"`
								Large string `json:"large"`
							} `json:"versions"`
						} `json:"videos"`
					} `json:"idml"`
					Reviews struct {
						AverageOverallRating  float64 `json:"averageOverallRating"`
						RatingValueFiveCount  int     `json:"ratingValueFiveCount"`
						PercentageFiveCount   int     `json:"percentageFiveCount"`
						RatingValueFourCount  int     `json:"ratingValueFourCount"`
						PercentageFourCount   int     `json:"percentageFourCount"`
						RatingValueThreeCount int     `json:"ratingValueThreeCount"`
						PercentageThreeCount  int     `json:"percentageThreeCount"`
						RatingValueTwoCount   int     `json:"ratingValueTwoCount"`
						PercentageTwoCount    int     `json:"percentageTwoCount"`
						RatingValueOneCount   int     `json:"ratingValueOneCount"`
						PercentageOneCount    int     `json:"percentageOneCount"`
						TotalMediaCount       int     `json:"totalMediaCount"`
						TotalReviewCount      int     `json:"totalReviewCount"`
						Aspects               []struct {
							Name         string `json:"name"`
							Score        string `json:"score"`
							SnippetCount string `json:"snippetCount"`
						} `json:"aspects"`
						CustomerReviews   []WprReview `json:"customerReviews"`
						TopNegativeReview WprReview   `json:"topNegativeReview"`
						TopPositiveReview WprReview   `json:"topPositiveReview"`
					} `json:"reviews"`
					Product struct {
						AvailabilityStatus string  `json:"availabilityStatus"`
						AverageRating      float64 `json:"averageRating"`
						Brand              string  `json:"brand"`
						BrandUrl           string  `json:"brandUrl"`
						Badges             struct {
							Flags []struct {
								Text string `json:"text"`
								Key  string `json:"key"`
							} `json:"flags"`
							Labels []interface{} `json:"labels"`
							Tags   []interface{} `json:"tags"`
							Groups []interface{} `json:"groups"`
						} `json:"badges"`
						ManufacturerProductId   string  `json:"manufacturerProductId"`
						InteractiveProductVideo string  `json:"interactiveProductVideo"`
						ProductTypeId           string  `json:"productTypeId"`
						Model                   string  `json:"model"`
						BuyNowEligible          bool    `json:"buyNowEligible"`
						OfferType               string  `json:"offerType"`
						SnapEligible            bool    `json:"snapEligible"`
						IsWplusMember           bool    `json:"isWplusMember"`
						ShowBuyWithWplus        bool    `json:"showBuyWithWplus"`
						CanonicalUrl            string  `json:"canonicalUrl"`
						SellerName              string  `json:"sellerName"`
						SellerDisplayName       string  `json:"sellerDisplayName"`
						SellerStoreFrontURL     string  `json:"sellerStoreFrontURL"`
						SellerReviewCount       int     `json:"sellerReviewCount"`
						SellerAverageRating     float64 `json:"sellerAverageRating"`
						ReturnPolicy            struct {
							ReturnPolicyText string `json:"returnPolicyText"`
						} `json:"returnPolicy"`
						ShortDescription string `json:"shortDescription"`
						FulfillmentLabel []struct {
							CheckStoreAvailability bool   `json:"checkStoreAvailability"`
							Nessage                string `json:"message"`
							ShippingText           string `json:"shippingText"`
							FulfillmentText        string `json:"fulfillmentText"`
							LocationText           string `json:"locationText"`
							FulfillmentMethod      string `json:"fulfillmentMethod"`
							AddressEligibility     string `json:"addressEligibility"`
							FulfillmentType        string `json:"fulfillmentType"`
							PostalCode             string `json:"postalCode"`
						} `json:"fulfillmentLabel"`
						Location struct {
							StateOrProvinceCode string `json:"stateOrProvinceCode"`
							City                string `json:"city"`
							PostalCode          string `json:"postalCode"`
						} `json:"location"`
						FulfillmentOptions []struct {
							TypeName          string `json:"__typename"`
							Type              string `json:"type"`
							Intent            bool   `json:"intent"`
							AvailableQuantity int    `json:"availableQuantity"`
							MaxOrderQuantity  int    `json:"maxOrderQuantity"`
							OrderLimit        int    `json:"orderLimit"`
							SpeedDetails      struct {
								FulfillmentBadge string `json:"fulfillmentBadge,omitempty"`
								DeliveryDate     string `json:"deliveryDate,omitempty"`
								FulfillmentPrice struct {
									Price       float64 `json:"price"`
									PriceString string  `json:"priceString"`
								} `json:"fulfillmentPrice,omitempty"`
								FreeFulfillment bool `json:"freeFulfillment"`
								WPlusEligible   bool `json:"wPlusEligible,omitempty"`
							} `json:"speedDetails,omitempty"`
							LocationText           string `json:"locationText"`
							AvailabilityStatus     string `json:"availabilityStatus"`
							StoreName              string `json:"storeName,omitempty"`
							SubscriptionSubmessage string `json:"subscriptionSubmessage,omitempty"`
							InventoryStatus        string `json:"inventoryStatus,omitempty"`
							ProductLocation        string `json:"productLocation,omitempty"`
						} `json:"fulfillmentOptions"`
						TransactableOfferCount int  `json:"transactableOfferCount"`
						HasSellerBadge         bool `json:"hasSellerBadge"`
						Id                     bool `json:"id"`
						ImageInfo              struct {
							AllImages []struct {
								Url string `json:"url"`
							} `json:"allImages"`
							ThumbnailUrl string `json:"thumbnailUrl"`
						} `json:"imageInfo"`
						Name               string `json:"name"`
						SalesUnit          string `json:"salesUnit"`
						ItemId             string `json:"usItemId"`
						Personalizable     bool   `json:"personalizable"`
						GiftingEligibility bool   `json:"giftingEligibility"`
						PreOrder           struct {
							IsPreOrder  bool   `json:"isPreOrder"`
							ReleaseDate string `json:"releaseDate"`
						} `json:"preOrder"`
						OrderMinLimit       int `json:"orderMinLimit"`
						OrderLimit          int `json:"orderLimit"`
						NumberOfReviews     int `json:"numberOfReviews"`
						SecondaryOfferPrice struct {
							CurrentPrice struct {
								PriceString string  `json:"priceString"`
								Price       float64 `json:"price"`
							} `json:"currentPrice"`
						} `json:"secondaryOfferPrice"`
						PriceInfo struct {
							CurrentPrice struct {
								Price       float64 `json:"price"`
								PriceString string  `json:"priceString"`
							} `json:"currentPrice"`
							WasPrice struct {
								Price       float64 `json:"price"`
								PriceString string  `json:"priceString"`
							} `json:"wasPrice"`
							UnitPrice struct {
								Price       float64 `json:"price"`
								PriceString string  `json:"priceString"`
							} `json:"unitPrice"`
							SavingsAmount struct {
								Amount      float64 `json:"price"`
								PriceString string  `json:"priceString"`
							} `json:"savingsAmount"`
							ComparisonPrice struct {
								Price       float64 `json:"price"`
								PriceString string  `json:"priceString"`
							} `json:"comparisonPrice"`
							ShipPrice struct {
								Price       float64 `json:"price"`
								PriceString string  `json:"priceString"`
							} `json:"shipPrice"`
							IsPriceReduced bool `json:"isPriceReduced"`
							ListPrice      struct {
								Price       float64 `json:"price"`
								PriceString string  `json:"priceString"`
							} `json:"listPrice"`
							PriceRange struct {
								MinPrice    float64 `json:"minPrice"`
								MaxPrice    float64 `json:"maxPrice"`
								PriceString string  `json:"priceString"`
							} `json:"priceRange"`
							AdditionalFees struct {
								DutyFee struct {
									Price       float64 `json:"price"`
									PriceString string  `json:"priceString"`
								} `json:"dutyFee"`
								ShippingAndImportFee struct {
									Price       float64 `json:"price"`
									PriceString string  `json:"priceString"`
								} `json:"shippingAndImportFee"`
								EstimatedTotalPrice struct {
									Price       float64 `json:"price"`
									PriceString string  `json:"priceString"`
								} `json:"estimatedTotalPrice"`
							} `json:"additionalFees"`
						} `json:"priceInfo"`
						Category struct {
							Path []struct {
								Name string `json:"name"`
								Url  string `json:"url"`
							} `json:"path"`
						} `json:"category"`
					} `json:"product"`
					ContentLayout struct {
						Modules []WprModule `json:"modules"`
					} `json:"contentLayout"`
				} `json:"data"`
			} `json:"initialData"`
		} `json:"pageProps"`
	} `json:"props"`
	Page  string `json:"page"`
	Query struct {
		ItemParams []string `json:"itemParams"`
	} `json:"query"`
	RuntimeConfig struct {
		Host struct {
			Wmt string `json:"wmt"`
		} `json:"host"`
	} `json:"runtimeConfig"`
}

func Walmart_ParseProductInfo(raw WprResult, baseUrl string) WpInfo {
	var product WpInfo
	rawProduct := raw.Props.PageProps.InitialData.Data.Product
	product.InteractiveProductVideo = rawProduct.InteractiveProductVideo
	product.AverageRating = rawProduct.AverageRating
	product.AvailableStatus = rawProduct.AvailabilityStatus
	if rawProduct.Badges.Flags != nil {
		for _, item := range rawProduct.Badges.Flags {
			product.Badges.Flags = append(product.Badges.Flags, item.Text)
		}
	}
	product.Brand = rawProduct.Brand
	product.BrandUrl = rawProduct.BrandUrl
	for _, item := range rawProduct.Category.Path {
		category := make(WStrMap)
		category["name"] = item.Name
		category["url"] = baseUrl + item.Url
		product.Categories = append(product.Categories, category)
	}
	product.EbtEligible = rawProduct.SnapEligible
	for _, item := range rawProduct.FulfillmentOptions {
		var fulfillment WpFulfillment
		fulfillment.AvailableQuantity = item.AvailableQuantity
		fulfillment.AvailableStatus = item.AvailabilityStatus != "NOT_AVAILABLE"
		fulfillment.Badge = item.SpeedDetails.FulfillmentBadge
		fulfillment.Date = item.SpeedDetails.DeliveryDate
		fulfillment.Location = item.LocationText
		fulfillment.MaxOrderQuantity = item.MaxOrderQuantity
		fulfillment.Price = item.SpeedDetails.FulfillmentPrice.PriceString
		fulfillment.Type = item.Type
		product.Fulfillments = append(product.Fulfillments, fulfillment)
	}
	for _, item := range rawProduct.ImageInfo.AllImages {
		product.Images = append(product.Images, item.Url)
	}
	product.Thumbnail = rawProduct.ImageInfo.ThumbnailUrl
	product.Model = rawProduct.Model
	product.Name = rawProduct.Name
	product.NumberOfReviews = rawProduct.NumberOfReviews
	product.PriceInfo.Price = rawProduct.PriceInfo.CurrentPrice.Price
	product.PriceInfo.PriceString = rawProduct.PriceInfo.CurrentPrice.PriceString
	product.PriceInfo.AdditionalFees.DutyFee = rawProduct.PriceInfo.AdditionalFees.DutyFee.PriceString
	product.PriceInfo.AdditionalFees.EstimatedTotalPrice = rawProduct.PriceInfo.AdditionalFees.EstimatedTotalPrice.PriceString
	product.PriceInfo.AdditionalFees.ShippingAndImportFee = rawProduct.PriceInfo.AdditionalFees.ShippingAndImportFee.PriceString
	product.PriceInfo.ComparisonPrice = rawProduct.PriceInfo.ComparisonPrice.PriceString
	product.PriceInfo.ListPrice = rawProduct.PriceInfo.ListPrice.PriceString
	product.PriceInfo.PriceRange = rawProduct.PriceInfo.PriceRange.PriceString
	product.PriceInfo.SecondaryOfferPrice = rawProduct.SecondaryOfferPrice.CurrentPrice.PriceString
	product.PriceInfo.ShipPrice = rawProduct.PriceInfo.ShipPrice.PriceString
	product.PriceInfo.UnitPrice = rawProduct.PriceInfo.UnitPrice.PriceString
	product.PriceInfo.WasPrice = rawProduct.PriceInfo.WasPrice.PriceString
	product.PriceInfo.SavingsAmount = rawProduct.PriceInfo.SavingsAmount.PriceString
	product.Seller.AverageRating = rawProduct.SellerAverageRating
	product.Seller.Name = rawProduct.SellerName
	product.Seller.DisplayName = rawProduct.SellerDisplayName
	product.Seller.StoreUrl = rawProduct.SellerStoreFrontURL
	product.ReturnPolicy = rawProduct.ReturnPolicy.ReturnPolicyText
	product.Location.PostalCode = rawProduct.Location.PostalCode
	product.Location.StateOrProvinceCode = rawProduct.Location.StateOrProvinceCode
	product.Location.City = rawProduct.Location.City
	return product
}

func WP_ParseReview(rawReview WprReview) WpReview {
	var review WpReview
	review.ReviewId = rawReview.ReviewId
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
		var media WpMedia
		media.MediaType = item.MediaType
		media.Caption = item.Caption
		media.Rating = item.Rating
		media.Url = item.NormalUrl
		review.Media = append(review.Media, media)
	}
	return review
}

func WP_ParseReviewInfo(raw WprResult) WpReviewInfo {
	var reviewInfo WpReviewInfo
	rawReviews := raw.Props.PageProps.InitialData.Data.Reviews
	reviewInfo.AverageOverallRating = rawReviews.AverageOverallRating
	reviewInfo.PercentageFiveCount = rawReviews.PercentageFiveCount
	reviewInfo.RatingValueFiveCount = rawReviews.RatingValueFiveCount
	reviewInfo.PercentageFourCount = rawReviews.PercentageFourCount
	reviewInfo.RatingValueFourCount = rawReviews.RatingValueFourCount
	reviewInfo.PercentageThreeCount = rawReviews.PercentageThreeCount
	reviewInfo.RatingValueThreeCount = rawReviews.RatingValueThreeCount
	reviewInfo.PercentageTwoCount = rawReviews.PercentageTwoCount
	reviewInfo.RatingValueTwoCount = rawReviews.RatingValueTwoCount
	reviewInfo.PercentageOneCount = rawReviews.PercentageOneCount
	reviewInfo.RatingValueOneCount = rawReviews.RatingValueOneCount
	reviewInfo.TotalReviewCount = rawReviews.TotalReviewCount
	reviewInfo.TotalMediaCount = rawReviews.TotalMediaCount
	reviewInfo.TopNegativeReview = WP_ParseReview(rawReviews.TopNegativeReview)
	reviewInfo.TopPositiveReview = WP_ParseReview(rawReviews.TopPositiveReview)
	for _, item := range rawReviews.CustomerReviews {
		reviewInfo.CustomerReviews = append(reviewInfo.CustomerReviews, WP_ParseReview(item))
	}
	return reviewInfo
}

func WP_ParseAbout(raw WprResult) WpAbort {
	var about WpAbort
	rawAbout := raw.Props.PageProps.InitialData.Data.Idml
	about.Details.ShortDescription = rawAbout.ShortDescription
	about.Details.LongDescription = rawAbout.LongDescription
	about.Directions = rawAbout.Directions
	about.Ingredients = rawAbout.Ingredients.Ingredients
	about.NutritionInfo = rawAbout.NutritionFacts
	about.Specifications = rawAbout.Specifications
	about.Warnings = rawAbout.Warnings
	for _, item := range rawAbout.Videos {
		video := make(WStrMap)
		video["poster"] = item.Poster
		video["title"] = item.Title
		video["small_url"] = item.Versions.Small
		video["large_url"] = item.Versions.Large
		about.Videos = append(about.Videos, video)
	}
	about.Warranty = rawAbout.Warranty.Information
	about.Highlights = rawAbout.ProductHighlights
	return about
}

func WP_ParseRelatedSearch(module WprModule, baseUrl string) []WStrMap {
	var relatedSearch []WStrMap
	for _, item := range module.Configs.RelatedSearch.RelatedSearch {
		search := make(WStrMap)
		search["title"] = item.Title
		search["url"] = baseUrl + "/" + item.Url
		relatedSearch = append(relatedSearch, search)
	}
	return relatedSearch
}

func WP_ParseRelatedPages(module WprModule, baseUrl string) []WStrMap {
	var relatedPages []WStrMap
	for _, item := range module.Configs.SeoItemRelmData.Relm {
		page := make(WStrMap)
		page["title"] = item.Name
		page["url"] = baseUrl + item.Url
		relatedPages = append(relatedPages, page)
	}
	return relatedPages
}

func Walmart_ProductPageScraper(jsonTag *goquery.Selection) WpResult {
	var result WpResult
	var data WpData
	var raw WprResult
	// parse product info
	json.Unmarshal([]byte(jsonTag.Text()), &raw)
	baseUrl := raw.RuntimeConfig.Host.Wmt
	data.Product = Walmart_ParseProductInfo(raw, baseUrl)
	data.About = WP_ParseAbout(raw)
	data.Review = WP_ParseReviewInfo(raw)
	for _, module := range raw.Props.PageProps.InitialData.Data.ContentLayout.Modules {
		if module.Type == "RelatedSearch" {
			data.RelatedSearch = WP_ParseRelatedSearch(module, baseUrl)
		} else if module.Type == "ItemRelatedShelves" {
			data.RelatedPages = WP_ParseRelatedPages(module, baseUrl)
		}
	}
	result.Data = data
	result.Status = "parse_successful"
	result.URL = baseUrl + raw.Props.PageProps.InitialData.Data.Product.CanonicalUrl
	return result
}

func Walmart_IsProductPage(jsonTag *goquery.Selection) bool {
	var page WRawResult
	json.Unmarshal([]byte(jsonTag.Text()), &page)
	return page.Page != "" && strings.Split(page.Page, "/")[1] == "ip"
}
