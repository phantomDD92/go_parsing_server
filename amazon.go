package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func handleAmazonHtml(filename string) bool {
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
	if isAmazonSearch(doc) {
		result := parseAmazonSearch(doc)
		return saveJsonFile(result, filename)
	} else if isAmazonReview(doc) {
		result := parseAmazonReview(doc)
		return saveJsonFile(result, filename)
	} else {
		result := parseAmazonProduct(doc)
		return saveJsonFile(result, filename)
	}
}

func handleAmazonPost(c *gin.Context) {
	var postData PostData
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
	if isAmazonSearch(doc) {
		result = parseAmazonSearch(doc)
	} else if isAmazonReview(doc) {
		result = parseAmazonReview(doc)
	} else {
		result = parseAmazonProduct(doc)
	}
	// saveJsonFile(result, "result")
	c.JSON(http.StatusOK, result)
}
