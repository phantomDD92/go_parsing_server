package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func Amazon_ParseHtml(filename string) bool {
	println("### ", filename)
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
	if Amazon_IsSearchPage(doc) {
		result := Amazon_SearchPagesScraper(doc)
		return saveJsonFile(result, filename)
	} else if Amazon_IsReviewPage(doc) {
		result := Amazon_ReviewPagesScraper(doc)
		return saveJsonFile(result, filename)
	} else {
		result := Amazon_ProductPagesScraper(doc)
		return saveJsonFile(result, filename)
	}
}

func Amazon_PostRequest(c *gin.Context) {
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
	var result interface{}
	if Amazon_IsSearchPage(doc) {
		result = Amazon_SearchPagesScraper(doc)
	} else if Amazon_IsReviewPage(doc) {
		result = Amazon_ReviewPagesScraper(doc)
	} else {
		result = Amazon_ProductPagesScraper(doc)
	}
	// saveJsonFile(result, "result")
	c.JSON(http.StatusOK, result)
}
