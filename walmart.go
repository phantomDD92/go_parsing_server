package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func handleWalmartHtml(filename string) bool {
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
	result := Walmart_SearchPageScraper(doc)
	// return true
	return saveJsonFile(result, filename)
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
	var result interface{}
	result = Walmart_SearchPageScraper(doc)
	// id := uuid.New()
	// saveJsonFile(result, id.String())
	c.JSON(http.StatusOK, result)
}
