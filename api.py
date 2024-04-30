import requests
import sys
import os
import json

PRODUCT_OVERALL_FILEDS = [
    "aplus_present", 
    "availability_status", 
    "average_rating",
    "brand",
    "brand_url",
    "customization_options",
    "feature_bullets",
    "full_description",
    "is_coupon_exists",
    "model",
    "name",
    "price",
    "price_string",
    "price_currency",
    "product_category",
    "product_information",
    "shipping_price",
    "ships_from",
    "sold_by",
    "total_reviews",
]

PRODUCT_CRITICAL_FILEDS = [
    "name",
    "brand",
    "brand_url",
    "model",
    "full_description",
    "price",
    "price_string",
    "availability_status", 
    "product_category",
    "average_rating",
    "total_reviews",
    "sold_by",
]

REVIEW_OVERALL_FILEDS = [
    "average_rating", 
    "total_reviews", 
    "5_star_ratings",
    "5_star_percentage",
    "4_star_ratings",
    "4_star_percentage",
    "3_star_ratings",
    "3_star_percentage",
    "2_star_ratings",
    "2_star_percentage",
    "1_star_ratings",
    "1_star_percentage",
    "product",
    "top_positive_review",
    "top_critical_review",
    "reviews",
    "pagination",
]

REVIEW_CRITICAL_FILEDS = [
    "average_rating", 
    "total_reviews", 
    "product",
    "reviews",
]

SEARCH_OVERALL_FILEDS = [
    "ads",
    "explore_more_items",
    "next_pages",
    "results",
]

SEARCH_RESULT_OVERALL_FILEDS = [
    "asin",
	"has_prime",
	"image",
	"is_amazon_choice",
	"is_best_seller",
	"is_limited_deal",
	"name",
    "position",
    "price",
    "price_string",
    "price_symbol",
    "stars",
    "total_reviews",
    "type",
    "url",
]

SEARCH_RESULT_CRITICAL_FILEDS = [
    "asin",
    "image",
    "name",
    "stars",
    "price",
    "price_string",
    "total_reviews",
]


class ApiTester():
    def __init__(self) -> None:
        pass
    
    def saveResponse(self, filename, result):
        jsonname = ".".join(filename.split(".")[:-1]) + ".json"
        with open(jsonname, "w", encoding="utf-8") as f:
            json.dump(result, f, indent="\t")

    def isEmpty(self, value):
        return value == None or value == ""
    
    def isSearch(self, data):
        return data.get("ads", None) != None

    def isReview(self, data):
        return data.get("reviews", None) != None

    def scoreForProduct(self, data, fields):
        count = 0
        score = 0
        for key in fields:
            count += 1
            if not self.isEmpty(data.get(key, None)):
                score += 1
        return score / count
    
    def scoreForReview(self, data, fields):
        count = 0
        score = 0
        for key in fields:
            count += 1
            if not self.isEmpty(data.get(key, None)):
                score += 1
        return score / count
    
    def scoreForSearch(self, data, fields):
        if self.isEmpty(data.get("results", None)) or len(data["results"]) == 0:
            return 0
        count = 0
        score = 0
        for record in data["results"]:
            for key in fields:
                count += 1
                if not self.isEmpty(record.get(key, None)):
                    score += 1
        return score / count


    def overallScore(self, data):
        if self.isSearch(data):
            return self.scoreForSearch(data, SEARCH_RESULT_OVERALL_FILEDS)
        elif self.isReview(data):
            return self.scoreForReview(data, REVIEW_OVERALL_FILEDS)
        else:
            return self.scoreForProduct(data, PRODUCT_OVERALL_FILEDS)
        

    def criticalScore(self, data):
        if self.isSearch(data):
            return self.scoreForSearch(data, SEARCH_RESULT_CRITICAL_FILEDS)
        elif self.isReview(data):
             return self.scoreForReview(data, REVIEW_CRITICAL_FILEDS)
        else:
            return self.scoreForProduct(data, PRODUCT_CRITICAL_FILEDS)   
        
    def hasData(self, result):
        return result.get("data", None) != None
    
    def start(self, filepath):
        try:
            with open(filepath, 'r', encoding="utf-8") as f:
                html = f.read()
            resp = requests.post("http://localhost:8080/post", json={
                "url": "url",
                "html": html
            })        
            result = resp.json()
            self.saveResponse(filepath, result)
            if not self.hasData(result):
                print(f"ERROR : Html Parsing is failed")
                return
            print(f"Overall Coverage Score(0~1) : {self.overallScore(result["data"]):.2f}")
            print(f"Critical Coverage Score(0~1) : {self.criticalScore(result["data"]):.2f}")

        except Exception as e:
            print(f"ERROR : Api Test is failed due to {str(e)}")
        
def main():
    if len(sys.argv) != 2:
        print("USAGE : $ python apitester.py <htmlfile>")
        exit()
    filepath = sys.argv[1]
    if not os.path.exists(filepath):
        print(f"ERROR : ${filepath} does not exist.")
        exit()
    tester = ApiTester()
    tester.start(filepath)
    
if __name__ == "__main__":
    main()