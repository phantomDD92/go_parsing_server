package main

import (
	"github.com/PuerkitoBio/goquery"
)

const SELECTOR_QUESTION = "[jsname='yEVEwb']"
const SELECTOR_TOP_ADS = "#taw"
const SELECTOR_BOTTOM_ADS = "#bottomads"
const SELECTOR_SEARCH = "#search"
const SELECTOR_SEARCH_ITEM = ".MjjYud"
const SELECTOR_ADS_ITEM = ".uEierd"
const SELECTOR_ADS_IMAGE_CONTAINER = ".SuXxEf"
const SELECTOR_ADS_LINK_CONTAINER = ".v5yQqb"
const SELECTOR_ADS_BLOCK_CONTAINER = ".UBEOKe"
const SELECTOR_ADS_BLOCK = ".MhgNwc"
const SELECTOR_ADS_INLINE_CONTAINER = ".bOeY0b"

const PATTERN_TOTAL_RESULTS = `About (\d{1,3}(,\d{3})*) results`
const PATTERN_TIME_TAKEN = `\((\d+(\.\d+)?) seconds\)`

type GoogleLink map[string]string

type GooglePageUrl struct {
	Page int    `json:"page"`
	Url  string `json:"url"`
}

type GoogleSiteLinks struct {
	Inline []GoogleLink `json:"inline"`
	Block  []GoogleLink `json:"block"`
}

type GoogleKnowledgeGraph struct {
	Title        string              `json:"title"`
	Image        string              `json:"image"`
	Description  string              `json:"description"`
	Source       GoogleLink          `json:"source"`
	Related      []map[string]string `json:"related"`
	RelatedLink  string              `json:"related_link"`
	SocialMedia  []GoogleLink        `json:"social_media"`
	SeeMoreAbout []GoogleLink        `json:"see_more_about"`
}

type GoogleAdsInfo struct {
	Position      int             `json:"position"`
	BlockPosition string          `json:"block_position"`
	Title         string          `json:"title"`
	Link          string          `json:"link"`
	Thumbnail     string          `json:"thumbnail"`
	DisplayedLink string          `json:"displayed_link"`
	TrackingLink  string          `json:"tracking_link"`
	Description   string          `json:"description"`
	Sitelinks     GoogleSiteLinks `json:"sitelinks"`
}

type GoogleOrganicResult struct {
	Position      int             `json:"position"`
	Title         string          `json:"title"`
	Snippet       string          `json:"snippet"`
	Link          string          `json:"link"`
	Date          string          `json:"date"`
	DisplayedLink string          `json:"displayed_link"`
	Thumbnail     string          `json:"thumbnail"`
	SiteLinks     GoogleSiteLinks `json:"site_links"`
}

type GoogleQuestion struct {
	Question string `json:"question"`
}

type GoogleSearchResult struct {
	SearchInfo struct {
		TotalResults       int     `json:"total_results"`
		TimeTakenDisplayed float64 `json:"time_taken_displayed"`
		QueryDisplayed     string  `json:"query_displayed"`
	} `json:"search_information"`
	Ads              []GoogleAdsInfo       `json:"ads"`
	KnowledgeGraph   GoogleKnowledgeGraph  `json:"knowledge_graph"`
	RelatedQuestions []GoogleQuestion      `json:"related_questions"`
	AnswerBox        interface{}           `json:"answer_box"`
	OrganicResults   []GoogleOrganicResult `json:"organic_results"`
	RelatedSearchs   []GoogleLink          `json:"related_searches"`
	More             string                `json:"more"`
}

func parseAdsTag(tag *goquery.Selection, position int, block string) GoogleAdsInfo {
	var ads GoogleAdsInfo
	ads.Position = position
	ads.BlockPosition = block
	headTag := tag.Find(SELECTOR_ADS_LINK_CONTAINER).First()
	if headTag.Length() > 0 {
		linkTag := headTag.Find("a").First()
		if linkTag.Length() > 0 {
			ads.Link = linkTag.AttrOr("href", "")
			ads.TrackingLink = linkTag.AttrOr("data-rw", "")
			titleTag := linkTag.Find("span").First()
			if titleTag.Length() > 0 {
				ads.Title = normalizeText(titleTag.Text())
			}
			thumbTag := linkTag.Find("img").First()
			if thumbTag.Length() > 0 {
				ads.Thumbnail = thumbTag.AttrOr("src", "")
			}
			displayeLinkTag := linkTag.Find(".x2VHCd").First()
			if displayeLinkTag.Length() > 0 {
				ads.DisplayedLink = normalizeText(displayeLinkTag.Text())
			}
		}
	}
	descTag := headTag.Next()
	if descTag.Length() > 0 {
		ads.Description = normalizeText(descTag.Text())
	}
	blockLinkTags := tag.Find(SELECTOR_ADS_BLOCK)
	if blockLinkTags.Length() > 0 {
		blockLinkTags.Each(func(i int, s *goquery.Selection) {
			h3Tag := s.Find("h3").First()
			if h3Tag.Length() > 0 {
				link := make(GoogleLink)
				linkTag := h3Tag.Find("a").First()
				if linkTag.Length() > 0 {
					link["link"] = linkTag.AttrOr("href", "")
					link["title"] = normalizeText(linkTag.Text())
				}
				textTag := h3Tag.Next()
				if textTag.Length() > 0 {
					link["description"] = normalizeText(textTag.Text())
				}
				ads.Sitelinks.Block = append(ads.Sitelinks.Block, link)
			}
		})
	}
	inlineTags := tag.Find(SELECTOR_ADS_INLINE_CONTAINER).Find("a")
	if inlineTags.Length() > 0 {
		inlineTags.Each(func(i int, s *goquery.Selection) {
			link := make(GoogleLink)
			link["link"] = s.AttrOr("href", "")
			link["title"] = normalizeText(s.Text())
			ads.Sitelinks.Inline = append(ads.Sitelinks.Inline, link)
		})
	}
	return ads
}

func parseQuestionTag(tag *goquery.Selection) []GoogleQuestion {
	var questions []GoogleQuestion
	questionTags := tag.Find(SELECTOR_QUESTION)
	questionTags.Each(func(i int, s *goquery.Selection) {
		var question GoogleQuestion
		question.Question = normalizeText(s.Text())
		questions = append(questions, question)
	})
	return questions
}

func parseKnowledgeTag(tag *goquery.Selection, baseUrl string) GoogleKnowledgeGraph {
	var knowledge GoogleKnowledgeGraph
	titleTag := tag.Find("[data-attrid='title']").First()
	if titleTag.Length() > 0 {
		knowledge.Title = normalizeText(titleTag.Text())
	}
	imageTag := tag.Find("[data-attrid='image']").Find("img").First()
	if imageTag.Length() > 0 {
		knowledge.Image = imageTag.AttrOr("src", "")
	}
	descTag := tag.Find("[data-attrid='description']").First()
	if descTag.Length() > 0 {
		knowledge.Description = normalizeText(descTag.Find("span").First().Text())
		sourceTag := descTag.Find("a").First()
		if sourceTag.Length() > 0 {
			knowledge.Source = make(map[string]string)
			knowledge.Source["name"] = normalizeText(sourceTag.Text())
			knowledge.Source["link"] = sourceTag.AttrOr("href", "")
		}
	}
	relatedTags := tag.Find(".Z1hOCe")
	if relatedTags.Length() > 0 {
		relatedTags.Each(func(i int, s *goquery.Selection) {
			related := make(map[string]string)
			linkTag := s.Find("a").First()
			if linkTag.Length() > 0 {
				related["name"] = normalizeText(linkTag.Text())
				related["link"] = normalizeUrl(baseUrl, linkTag.AttrOr("href", ""))
			}
			contentTag := linkTag.Parent().Next()
			if contentTag.Length() > 0 {
				related["content"] = extractText(contentTag, " ")
			}
			imageTag := s.Find("img").First()
			if imageTag.Length() > 0 {
				related["image"] = imageTag.AttrOr("src", "")
			}
			knowledge.Related = append(knowledge.Related, related)
		})
	}
	socialTags := tag.Find("g-link.w23JUc")
	socialTags.Each(func(i int, s *goquery.Selection) {
		linkTag := s.Find("a").First()
		if linkTag.Length() > 0 {
			titleLink := make(GoogleLink)
			titleLink["title"] = normalizeText(linkTag.Text())
			titleLink["link"] = linkTag.AttrOr("href", "")
			knowledge.SocialMedia = append(knowledge.SocialMedia, titleLink)
		}
	})
	seeMoreTags := tag.Find("[data-reltype='sideways']")
	seeMoreTags.Each(func(i int, s *goquery.Selection) {
		linkTag := s.Find("a").First()
		if linkTag.Length() > 0 {
			titleLink := make(GoogleLink)
			titleLink["title"] = normalizeText(linkTag.Text())
			titleLink["link"] = normalizeUrl(baseUrl, linkTag.AttrOr("href", ""))
			knowledge.SeeMoreAbout = append(knowledge.SeeMoreAbout, titleLink)
		}
	})
	// Parse related-link tag
	return knowledge
}

func parseOrganicResult(tag *goquery.Selection, position int) GoogleOrganicResult {
	var result GoogleOrganicResult
	divTag := tag.Find("[jscontroller='yz368b']").First()
	result.Position = position
	if divTag.Length() > 0 {
		headTag := tag.Find("g-link").Find("a").First()
		if headTag.Length() > 0 {
			result.Link = headTag.AttrOr("href", "")
			titleTag := headTag.Find("h3").First()
			if titleTag.Length() > 0 {
				result.Title = normalizeText(titleTag.Text())
			}
		}
		thumbTag := headTag.Find("img.XNo5Ab").First()
		if thumbTag.Length() > 0 {
			result.Thumbnail = thumbTag.AttrOr("src", "")
		}
		displayedLinkTag := headTag.Find("cite").First()
		if displayedLinkTag.Length() > 0 {
			result.DisplayedLink = normalizeText(displayedLinkTag.Text())
		}
		blockTags := divTag.Find("g-inner-card")
		if blockTags.Length() > 0 {
			blockTags.Each(func(i int, s *goquery.Selection) {
				linkTag := s.Find(".tw-res").Find("a").First()
				if linkTag.Length() > 0 {
					blockLink := make(GoogleLink)
					blockLink["link"] = linkTag.AttrOr("href", "")
					snippetTag := linkTag.Next()
					if snippetTag.Length() > 0 {
						blockLink["snippet"] = normalizeText(snippetTag.Text())
					}
					dateTag := snippetTag.Next()
					if dateTag.Length() > 0 {
						blockLink["date"] = normalizeText(dateTag.Text())
					}
					result.SiteLinks.Block = append(result.SiteLinks.Block, blockLink)
				}
			})
		}
	} else {
		headTag := tag.Find(".yuRUbf")
		if headTag.Length() > 0 {
			linkTag := headTag.Find("a").First()
			if linkTag.Length() > 0 {
				titleTag := linkTag.Find("h3").First()
				result.Link = linkTag.AttrOr("href", "")
				if titleTag.Length() > 0 {
					result.Title = normalizeText(titleTag.Text())
				}
			}
			thumbTag := headTag.Find("img.XNo5Ab").First()
			if thumbTag.Length() > 0 {
				result.Thumbnail = thumbTag.AttrOr("src", "")
			}
			displayedLinkTag := headTag.Find("cite").First()
			if displayedLinkTag.Length() > 0 {
				result.DisplayedLink = normalizeText(displayedLinkTag.Text())
			}
		}
		snippetTag := tag.Find(".VwiC3b").First()
		if snippetTag.Length() > 0 {
			dateTag := snippetTag.Find(".LEwnzc").Find("span").First()
			if dateTag.Length() > 0 {
				result.Date = normalizeText(dateTag.Text())
			}
			result.Snippet = normalizeText(snippetTag.Text())
		}
		inlineTags := tag.Find(".HiHjCd").Find("a")
		if inlineTags.Length() > 0 {
			inlineTags.Each(func(i int, s *goquery.Selection) {
				inlineLink := make(GoogleLink)
				inlineLink["link"] = s.AttrOr("href", "")
				inlineLink["title"] = normalizeText(s.Text())
				result.SiteLinks.Inline = append(result.SiteLinks.Inline, inlineLink)
			})
		}
		blockTags := tag.Find("table.jmjoTe").Find("td.cIkxbf")
		if blockTags.Length() > 0 {
			blockTags.Each(func(i int, s *goquery.Selection) {
				h3Tag := s.Find("h3").First()
				if h3Tag.Length() > 0 {
					blockLink := make(GoogleLink)
					linkTag := h3Tag.Find("a").First()
					if linkTag.Length() > 0 {
						blockLink["link"] = linkTag.AttrOr("href", "")
						blockLink["title"] = normalizeText(linkTag.Text())
					}
					snippetTag := h3Tag.Next()
					if snippetTag.Length() > 0 {
						blockLink["snippet"] = normalizeText(snippetTag.Text())
					}
					result.SiteLinks.Block = append(result.SiteLinks.Block, blockLink)
				}
			})
		}
	}
	return result
}

func parseRelatedSearches(tag *goquery.Selection, baseUrl string) []GoogleLink {
	var relatedSearchs []GoogleLink
	searchTags := tag.Find("a")
	if searchTags.Length() > 0 {
		searchTags.Each(func(i int, s *goquery.Selection) {
			link := make(GoogleLink)
			link["link"] = normalizeUrl(baseUrl, s.AttrOr("href", ""))
			link["query"] = normalizeText(s.Text())
			relatedSearchs = append(relatedSearchs, link)
		})
	}
	return relatedSearchs
}

func parseGoogleSearch(doc *goquery.Document) GoogleSearchResult {
	var result GoogleSearchResult
	// Get base url
	baseUrl := "https://www.goolge.ca"
	searchForm := doc.Find("#searchform")
	if searchForm.Length() > 0 {
		homeLinkTag := searchForm.Find("a").First()
		if homeLinkTag.Length() > 0 {
			baseUrl = getBaseUrl(homeLinkTag.AttrOr("href", ""))
		}
	}
	// Parsing stats tag
	statsTag := doc.Find("#result-stats").First()
	if statsTag.Length() > 0 {
		statsText := normalizeText(statsTag.Text())
		total, err := extractIntFromPattern(statsText, PATTERN_TOTAL_RESULTS)
		if err == nil {
			result.SearchInfo.TotalResults = total
		}
		time, err1 := extractFloatFromPattern(statsText, PATTERN_TIME_TAKEN)
		if err1 == nil {
			result.SearchInfo.TimeTakenDisplayed = time
		}
	}
	// Parse query
	queryTag := doc.Find("textarea#APjFqb").First()
	if queryTag.Length() > 0 {
		result.SearchInfo.QueryDisplayed = normalizeText(queryTag.Text())
	}
	// parse knowlege
	knowledgeTag := doc.Find(".kp-wholepage").First()
	if knowledgeTag.Length() > 0 {
		result.KnowledgeGraph = parseKnowledgeTag(knowledgeTag, baseUrl)
	}
	// Parse Top Ads
	position := 1
	topAdsTags := doc.Find(SELECTOR_TOP_ADS).Find(SELECTOR_ADS_ITEM)
	if topAdsTags.Length() > 0 {
		topAdsTags.Each(func(i int, s *goquery.Selection) {
			ads := parseAdsTag(s, position, "top")
			position += 1
			result.Ads = append(result.Ads, ads)
		})
	}
	// Parse Bottom Ads
	bottomAdsTags := doc.Find(SELECTOR_BOTTOM_ADS).Find(SELECTOR_ADS_ITEM)
	if bottomAdsTags.Length() > 0 {
		bottomAdsTags.Each(func(i int, s *goquery.Selection) {
			ads := parseAdsTag(s, position, "bottom")
			position += 1
			result.Ads = append(result.Ads, ads)
		})
	}
	// parse search
	position = 1
	searchTags := doc.Find(SELECTOR_SEARCH).Find(SELECTOR_SEARCH_ITEM)
	if searchTags.Length() > 0 {
		searchTags.Each(func(i int, s *goquery.Selection) {
			questionTag := s.Find("[jscontroller='Da4hkd']").First()
			if questionTag.Length() > 0 {
				result.RelatedQuestions = parseQuestionTag(questionTag)
			} else {
				organicRes := parseOrganicResult(s, position)
				position += 1
				result.OrganicResults = append(result.OrganicResults, organicRes)
			}

		})
	}
	// Parse related-search
	relatedTag := doc.Find("#botstuff").Find(".y6Uyqe")
	if relatedTag.Length() > 0 {
		result.RelatedSearchs = parseRelatedSearches(relatedTag, baseUrl)
	}
	// Parse more results
	moreTag := doc.Find("#botstuff").Find("[jscontroller='bpec7b']").Find("a").First()
	if moreTag.Length() > 0 {
		result.More = normalizeUrl(baseUrl, moreTag.AttrOr("href", ""))
	}
	return result
}
