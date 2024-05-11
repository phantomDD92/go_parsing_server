package main

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

type WRawResult struct {
	Page string `json:"page"`
}

func Walmart_ParseHtml(filename string) bool {
	content, err := os.ReadFile("./data/" + filename + ".html")
	if err != nil {
		println(err)
		return false
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(content)))
	if err != nil {
		println(err)
		return false
	}
	json, err := Walmart_ExtractJson(doc)
	if err != nil {
		println(err)
		return false
	}
	var result interface{}
	if Walmart_IsSearchPage(json) {
		result = Walmart_SearchPageScraper(json)
	} else if Walmart_IsProductPage(json) {
		result = Walmart_ProductPageScraper(json)
	} else {
		return false
	}
	// return true
	return saveJsonFile(result, filename)
}

func Walmart_ExtractJson(doc *goquery.Document) (*goquery.Selection, error) {
	jsonTag := doc.Find("script#__NEXT_DATA__").First()
	if jsonTag.Length() > 0 {
		return jsonTag, nil
	}
	return nil, errors.New("parsing failed")
}

func Walmart_PostRequest(c *gin.Context) {
	var postData RequestData
	// Get post data
	if err := c.BindJSON(&postData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// parse html using goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(postData.Html))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	json, err := Walmart_ExtractJson(doc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var result interface{}
	if Walmart_IsSearchPage(json) {
		result = Walmart_SearchPageScraper(json)
	} else if Walmart_IsProductPage(json) {
		result = Walmart_ProductPageScraper(json)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported page"})
		return
	}
	// id := uuid.New()
	// saveJsonFile(result, id.String())
	c.JSON(http.StatusOK, result)
}
