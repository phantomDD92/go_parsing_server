package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

type PostData struct {
	Url  string `json:"url"`
	Html string `json:"html"`
}

func saveJsonFile(result interface{}, filename string) bool {
	// Marshal the struct into JSON
	jsonData, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return false
	}
	// Write the JSON data to a file
	file, err := os.Create("./data/" + filename + ".json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return false
	}
	defer file.Close()
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return false
	}
	return true
}

func parseAndSave(filename string) bool {
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
	if !isSearchPage(doc) {
		result := parseProduct(doc)
		return saveJsonFile(result, filename)
	} else {
		result := parseSearch(doc)
		return saveJsonFile(result, filename)
	}
}

func handlePost(c *gin.Context) {
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
	if isSearchPage(doc) {
		var result SearchResult = parseSearch(doc)
		saveJsonFile(result, "result")
		c.JSON(http.StatusOK, result)
		return
	} else {
		var result ProductResult = parseProduct(doc)
		fmt.Println(result.Data.PriceCurrency, result.Data.Price)
		c.JSON(http.StatusOK, result)
		return
	}
}

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 512 << 20 // 8 MiB
	// Define a route and its handler
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin!",
		})
	})

	// Define a POST route and its handler
	r.POST("/post", handlePost)

	srv := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 16 << 20, // 1 MiB
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

}
