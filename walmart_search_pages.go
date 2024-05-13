package main

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type WalmartSearchRawItem struct {
	TypeName                  string `json:"__typename"`
	Type                      string `json:"type"`
	Id                        string `json:"id,omitempty"`
	UsItemId                  string `json:"usItemId,omitempty"`
	Name                      string `json:"name,omitempty"`
	CheckStoreAvailabilityATC bool   `json:"checkStoreAvailabilityATC,omitempty"`
	SeeShippingEligibility    bool   `json:"seeShippingEligibility,omitempty"`
	Brand                     string `json:"brand,omitempty"`
	ShortDescription          string `json:"shortDescription,omitempty"`
	WeightIncrement           int    `json:"weightIncrement,omitempty"`
	ImageInfo                 struct {
		Id           string `json:"id"`
		ThumbnailUrl string `json:"thumbnailUrl"`
		Size         string `json:"size"`
	} `json:"imageInfo,omitempty"`
	CanonicalUrl string `json:"canonicalUrl,omitempty"`
	Badges       struct {
		Flags []struct {
			Key  string `json:"key"`
			Text string `json:"text"`
			Type string `json:"type"`
		} `json:"flags"`
		Tags []struct {
			Key  string `json:"key"`
			Text string `json:"text"`
			Type string `json:"type"`
		} `json:"tags"`
		Groups []struct {
			Name    string        `json:"name"`
			Members []interface{} `json:"members"`
		} `json:"groups"`
	} `json:"badges,omitempty"`
	SnapEligible       bool    `json:"snapEligible,omitempty"`
	BuyNowEligible     bool    `json:"buyNowEligible,omitempty"`
	ClassType          string  `json:"classType,omitempty"`
	AverageRating      float64 `json:"averageRating,omitempty"`
	NumberOfReviews    int     `json:"numberOfReviews,omitempty"`
	SalesUnitType      string  `json:"salesUnitType,omitempty"`
	SellerName         string  `json:"sellerName,omitempty"`
	AvailabilityStatus struct {
		Display string `json:"display"`
		Value   string `json:"value"`
	} `json:"availabilityStatusV2,omitempty"`
	PriceInfo struct {
		ItemPrice        string `json:"itemPrice"`
		LinePrice        string `json:"linePrice"`
		UnitPrice        string `json:"unitPrice"`
		WasPrice         string `json:"wasPrice"`
		ShipPrice        string `json:"shipPrice"`
		PriceRangeString string `json:"priceRangeString"`
	} `json:"priceInfo,omitempty"`
	FulfillmentType                string `json:"fulfillmentType,omitempty"`
	ManufacturerName               string `json:"manufacturerName,omitempty"`
	ShowAtc                        bool   `json:"showAtc,omitempty"`
	ShowOptions                    bool   `json:"showOptions,omitempty"`
	ShowBuyNow                     bool   `json:"showBuyNow,omitempty"`
	AvailabilityStatusDisplayValue string `json:"availabilityStatusDisplayValue,omitempty"`
	ProductLocationDisplayValue    string `json:"productLocationDisplayValue,omitempty"`
	CanAddToCart                   bool   `json:"canAddToCart,omitempty"`
	Flag                           string `json:"flag,omitempty"`
	Badge                          struct {
		Key  string `json:"key"`
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"badge,omitempty"`
	FulfillmentBadgeGroups []struct {
		Text    string `json:"text"`
		SlaText string `json:"slaText"`
	} `json:"fulfillmentBadgeGroups,omitempty"`
	SpecialBuy   bool    `json:"specialBuy,omitempty"`
	PriceFlip    bool    `json:"priceFlip,omitempty"`
	Image        string  `json:"image,omitempty"`
	ImageSize    string  `json:"imageSize,omitempty"`
	IsOutOfStock bool    `json:"isOutOfStock,omitempty"`
	Price        float64 `json:"price,omitempty"`
	Rating       struct {
		AverageRating   float64 `json:"averageRating"`
		NumberOfReviews int     `json:"numberOfReviews"`
	} `json:"rating,omitempty"`
	SalesUnit           string        `json:"salesUnit,omitempty"`
	VariantList         []interface{} `json:"variantList,omitempty"`
	IsVariantTypeSwatch bool          `json:"isVariantTypeSwatch,omitempty"`
	IsSponsoredFlag     bool          `json:"isSponsoredFlag,omitempty"`
	TileTakeOverTile    struct {
		Title    string `json:"title"`
		Subtitle string `json:"subtitle"`
		Image    struct {
			Src string `json:"src"`
			Alt string `json:"alt"`
		} `json:"image"`
		TileCta []struct {
			CtaLink struct {
				ClickThrough struct {
					Value string `json:"value"`
				} `json:"clickThrough"`
				LinkText string `json:"linkText"`
				Title    string `json:"title"`
			} `json:"ctaLink"`
			CtaType string `json:"ctaType"`
		} `json:"tileCta"`
		AdsEnabled string `json:"adsEnabled"`
	} `json:"tileTakeOverTile,omitempty"`
}

type WalmartConfig map[string]interface{}

type WalmartRawPill struct {
	Title string `json:"title"`
	Url   string `json:"url"`
	Image struct {
		Src string `json:"src"`
		Alt string `json:"alt"`
	} `json:"image"`
}

type WalmartRawConfigs struct {
	ModuleType string `json:"moduleType,omitempty"`
	ViewConfig struct {
		Title       string `json:"title"`
		Image       string `json:"image"`
		DisplayName string `json:"displayName"`
		Description string `json:"description"`
		Url         string `json:"url"`
		UrlAlt      string `json:"urlAlt"`
	} `json:"viewConfig,omitempty"`
	Pills []struct {
		Title string `json:"title"`
		Url   string `json:"url"`
		Image struct {
			Src string `json:"src"`
			Alt string `json:"alt"`
		} `json:"image"`
	} `json:"pillsV2,omitempty"`
}

type WsrResult struct {
	Props struct {
		PageProps struct {
			InitialData struct {
				SearchResult struct {
					Title           string `json:"title"`
					AggregatedCount int    `json:"aggregatedCount"`
					ItemStacks      []struct {
						Title                 string                 `json:"title"`
						Description           string                 `json:"description"`
						TotalItemCountDisplay string                 `json:"totalItemCountDisplay"`
						Count                 int                    `json:"count"`
						Items                 []WalmartSearchRawItem `json:"items"`
					} `json:"itemStacks"`
					Pagination struct {
						MaxPage    int `json:"maxPage"`
						Properties struct {
							Ps               string `json:"ps"`
							AffinityOverride string `json:"affinityOverride"`
							Stores           string `json:"stores"`
							Grid             string `json:"grid"`
							Query            string `json:"query"`
							CatId            string `json:"cat_id"`
							Sort             string `json:"sort"`
							Page             int    `json:"page"`
							PageType         string `json:"pageType"`
						} `json:"pageProperties"`
					} `json:"paginationV2"`
					RelatedSearch []struct {
						Title string `json:"title"`
						Url   string `json:"url"`
						Image string `json:"imageUrl"`
					} `json:"relatedSearch"`
					Count          int  `json:"count"`
					GridItemsCount int  `json:"gridItemsCount"`
					HasMorePages   bool `json:"hasMorePages"`
					CatInfo        struct {
						CatId string `json:"catId"`
						Name  string `json:"name"`
					} `json:"catInfo"`
				} `json:"searchResult"`
				ContentLayout struct {
					Modules []struct {
						Type    string            `json:"type"`
						Name    string            `json:"name"`
						Version int               `json:"version"`
						Configs WalmartRawConfigs `json:"configs"`
					} `json:"modules"`
				} `json:"contentLayout"`
			} `json:"initialData"`
		} `json:"pageProps"`
	} `json:"props"`
	Page          string            `json:"page"`
	Query         map[string]string `json:"query"`
	RuntimeConfig struct {
		Host struct {
			Wmt string `json:"wmt"`
		} `json:"host"`
	} `json:"runtimeConfig"`
	DynamicIds []int  `json:"dynamicIds"`
	AppGip     bool   `json:"appGip"`
	Locale     string `json:"locale"`
}

type WalmartSearchProduct struct {
	Position         int      `json:"position"`
	ID               string   `json:"id"`
	ItemID           string   `json:"item_id"`
	Name             string   `json:"name"`
	Type             string   `json:"type"`
	Brand            string   `json:"brand"`
	ShortDescription string   `json:"short_description"`
	AverageRating    float64  `json:"average_rating"`
	NumberOfReviews  int      `json:"number_of_reviews"`
	SalesUnit        string   `json:"sales_unit"`
	SellerName       string   `json:"seller_name"`
	Image            string   `json:"image"`
	ImageSize        string   `json:"image_size"`
	Price            float64  `json:"price"`
	LinePrice        string   `json:"line_price"`
	WasPrice         string   `json:"was_price"`
	UnitPrice        string   `json:"unit_price"`
	ItemPrice        string   `json:"item_price"`
	PriceRange       string   `json:"price_range"`
	IsEligible       bool     `json:"is_eligible"`
	IsShowATC        bool     `json:"is_show_atc"`
	IsShowOptions    bool     `json:"is_show_options"`
	IsBuyNow         bool     `json:"is_buy_now"`
	IsSponsored      bool     `json:"is_sponsored"`
	IsOutofStock     bool     `json:"is_outofstock"`
	Availability     string   `json:"availability"`
	ProductLocation  string   `json:"product_location"`
	Flag             string   `json:"flag"`
	Fulfillment      []string `json:"fulfillment"`
	URL              string   `json:"url"`
}

type WalmartLink map[string]string

type WalmartSearchPill struct {
	Title string `json:"title"`
	Image string `json:"image"`
	URL   string `json:"url"`
}

type WalmartSearchBanner struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	URL         string `json:"url"`
	Button      string `json:"button"`
}

type WalmartSearchTile struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Image    string `json:"image"`
	Link     struct {
		Title string `json:"title"`
		Url   string `json:"url"`
	} `json:"site_link"`
}

type WsData struct {
	TotalCount        int                    `json:"total_count"`
	TotalCountDisplay string                 `json:"total_count_display"`
	Query             map[string]string      `json:"search_query"`
	RelatedSearch     []WalmartLink          `json:"related_search"`
	Results           []WalmartSearchProduct `json:"results"`
	Tiles             []WalmartSearchTile    `json:"takeover_tiles"`
	Pills             []WalmartSearchPill    `json:"search_pills"`
	Banner            WalmartSearchBanner    `json:"search_banner"`
	Pagination        struct {
		PageCount   int      `json:"page_count"`
		CurrentPage int      `json:"current_page"`
		PageLinks   []string `json:"page_links"`
	} `json:"pagination"`
}

type WsResult struct {
	Data   WsData `json:"data"`
	Status string `json:"status"`
	URL    string `json:"url"`
}

func normalizeWalmartImage(url string) string {
	return strings.Split(url, "?")[0]
}

func makeUrl(path string, params map[string]string) string {
	// Create a Values instance to encode the query parameters
	values := url.Values{}
	for key, value := range params {
		values.Add(key, value)
	}
	// Construct the query string
	queryString := values.Encode()
	return path + "?" + queryString
}

func parseWalmartSearchTile(item WalmartSearchRawItem) WalmartSearchTile {
	var tile WalmartSearchTile
	tile.Title = item.TileTakeOverTile.Title
	tile.Subtitle = item.TileTakeOverTile.Subtitle
	tile.Image = normalizeWalmartImage(item.TileTakeOverTile.Image.Src)
	if len(item.TileTakeOverTile.TileCta) > 0 {
		tile.Link.Title = item.TileTakeOverTile.TileCta[0].CtaLink.Title
		tile.Link.Url = item.TileTakeOverTile.TileCta[0].CtaLink.ClickThrough.Value
	}
	return tile
}
func parseWalmartSearchProduct(item WalmartSearchRawItem, position int, baseUrl string) WalmartSearchProduct {
	var product WalmartSearchProduct
	product.Position = position
	product.ID = item.Id
	product.ItemID = item.UsItemId
	product.Availability = item.AvailabilityStatusDisplayValue
	product.AverageRating = item.AverageRating
	product.Brand = item.Brand
	product.Flag = item.Flag
	for _, value := range item.FulfillmentBadgeGroups {
		product.Fulfillment = append(product.Fulfillment, value.Text+value.SlaText)
	}
	product.Image = normalizeWalmartImage(item.Image)
	product.ImageSize = item.ImageSize
	product.IsBuyNow = item.ShowBuyNow
	product.IsEligible = item.SnapEligible
	product.IsOutofStock = item.IsOutOfStock
	product.IsShowATC = item.ShowAtc
	product.IsShowOptions = item.ShowOptions
	product.IsSponsored = item.IsSponsoredFlag
	product.ItemPrice = item.PriceInfo.ItemPrice
	product.LinePrice = item.PriceInfo.LinePrice
	product.Name = item.Name
	product.NumberOfReviews = item.NumberOfReviews
	product.Price = item.Price
	product.PriceRange = item.PriceInfo.PriceRangeString
	product.ProductLocation = item.ProductLocationDisplayValue
	product.SalesUnit = item.SalesUnit
	product.SellerName = item.SellerName
	product.ShortDescription = item.ShortDescription
	product.Type = item.Type
	product.URL = baseUrl + item.CanonicalUrl
	product.UnitPrice = item.PriceInfo.UnitPrice
	product.WasPrice = item.PriceInfo.WasPrice
	return product
}

func parseWalmartPills(config WalmartRawConfigs) []WalmartSearchPill {
	var pills []WalmartSearchPill
	for _, value := range config.Pills {
		var pill WalmartSearchPill
		pill.Image = normalizeWalmartImage(value.Image.Src)
		pill.Title = value.Title
		pill.URL = value.Url
		pills = append(pills, pill)
	}
	return pills
}

func parseWalmartBanner(config WalmartRawConfigs) WalmartSearchBanner {
	var banner WalmartSearchBanner
	banner.Title = config.ViewConfig.DisplayName
	banner.URL = config.ViewConfig.Url
	banner.Button = config.ViewConfig.UrlAlt
	banner.Description = config.ViewConfig.Description
	banner.Image = normalizeWalmartImage(config.ViewConfig.Image)
	return banner
}

func Walmart_SearchPageScraper(jsonTag *goquery.Selection) WsResult {
	var raw WsrResult
	var result WsResult
	var data WsData
	baseUrl := "https://www.walmart.com"
	// dataTag := doc.Find("script#__NEXT_DATA__").First()
	// if dataTag.Length() > 0 {
	json.Unmarshal([]byte(jsonTag.Text()), &raw)
	baseUrl = raw.RuntimeConfig.Host.Wmt
	for _, item := range raw.Props.PageProps.InitialData.SearchResult.RelatedSearch {
		queryLink := make(WalmartLink)
		queryLink["link"] = baseUrl + raw.Page + "?" + item.Url
		queryLink["query"] = item.Title
		data.RelatedSearch = append(data.RelatedSearch, queryLink)
	}
	position := 1
	itemStack := raw.Props.PageProps.InitialData.SearchResult.ItemStacks[0]
	for _, item := range itemStack.Items {
		if item.TypeName == "Product" {
			product := parseWalmartSearchProduct(item, position, baseUrl)
			data.Results = append(data.Results, product)
			position += 1
		} else if item.TypeName == "TileTakeOverProductPlaceholder" {
			tile := parseWalmartSearchTile(item)
			data.Tiles = append(data.Tiles, tile)
		}
	}
	modules := raw.Props.PageProps.InitialData.ContentLayout.Modules
	for _, module := range modules {
		if module.Type == "PillsModule" {
			data.Pills = parseWalmartPills(module.Configs)
		} else if module.Type == "SearchBanner" {
			data.Banner = parseWalmartBanner(module.Configs)
		}
	}
	data.TotalCount = itemStack.Count
	data.TotalCountDisplay = itemStack.TotalItemCountDisplay
	data.Query = raw.Query
	pagination := raw.Props.PageProps.InitialData.SearchResult.Pagination
	data.Pagination.CurrentPage = pagination.Properties.Page
	data.Pagination.PageCount = pagination.MaxPage
	result.URL = makeUrl(baseUrl+raw.Page, raw.Query)
	for i := 1; i <= pagination.MaxPage; i++ {
		link := result.URL + "&affinityOverride=" + pagination.Properties.AffinityOverride
		if i != 1 {
			link += "&page=" + strconv.Itoa(i)
		}
		data.Pagination.PageLinks = append(data.Pagination.PageLinks, link)
	}
	// }
	result.Data = data
	result.Status = "parse_successful"
	return result
}

func Walmart_IsSearchPage(jsonTag *goquery.Selection) bool {
	var page WRawResult
	json.Unmarshal([]byte(jsonTag.Text()), &page)
	return page.Page != "" && strings.Split(page.Page, "/")[1] == "search"
}
