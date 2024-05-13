import requests
import json
import numpy as np
import datetime

TEST_URLS = [
    { "url" : 
        "https://www.walmart.com/reviews/product/172844767",
        # "https://www.walmart.com/search?catId=976759&facet=brand%3AMarketside&q=cake&sort=best_seller",
        # "https://www.walmart.com/search?q=metal%20garden%20hose&sort=best_match&affinityOverride=default&page=3",
        # "https://www.walmart.com/search?q=travel%20bags%20for%20luggage&sort=best_match&affinityOverride=default&page=2",
    }
]

CURRENT_TEST = "walmart-review"
# CURRENT_TEST = "walmart-product"
# CURRENT_TEST = "walmart-search"
# CURRENT_TEST = "google-search"
SAMPLE_COUNT = 1
# SAMPLE_COUNT = 1000

### Settings for api key
API_ENDPOINT_URL = "https://proxy.scrapeops.io/v1/sample-urls"
API_KEY_URL = "81a3edae-ece8-47ed-a672-13e7961a6fcc"

API_ENDPOINT_PROXY = 'https://proxy.scrapeops.io/v1/'
API_KEY_PROXY = "d0f35644-3095-48f9-b109-9146894ae581"

### Settings for test config
TEST_CONFIG = {
    "walmart-search": {
        "url": "*walmart.com/search*",
        "domain": "walmart.com",
        "endpoint": "walmart",
    },
    "walmart-product": {
        "url": "*walmart.com/ip/*",
        "domain": "walmart.com",
        "endpoint": "walmart",
    },
    "walmart-review": {
        "url": "*walmart.com/reviews/*",
        "domain": "walmart.com",
        "endpoint": "walmart",
    },
    "google-search": {
        "url": "*google.com/search?*",
        "domain": "google.com",
        "endpoint": "google",
    }
}

# API_ENDPOINT_PARSE = f'http://localhost:8080/v2/{TEST_CONFIG[CURRENT_TEST]["endpoint"]}'
API_ENDPOINT_PARSE = f'http://localhost:8080/{TEST_CONFIG[CURRENT_TEST]["endpoint"]}'
### Settings for coverage calculation
# name(*required) is the field name
# critical(default = False) means that this field attends to calculate critical coverage score
# weight(default = 1) means the importance of this field
# children(default = 1) means that this field contains subfield.

# Walmart Search Fields
WALMART_SEARCH_FIELDS = [
    {"name": "data", "critical": True, "weight" : 5, "children" : [
        {"name": "results", "critical": True, "weight" : 10, "children" : [
            {"name": "position", "critical": True },
            {"name": "id", "critical": True},
            {"name": "item_id"},
            {"name": "name", "critical": True, "weight" : 3},
            {"name": "type" },
            {"name": "brand" },
            {"name": "short_description" },
            {"name": "average_rating", "critical": True, "weight" : 2},
            {"name": "number_of_reviews", "critical": True, "weight" : 2},
            {"name": "sales_unit"},
            {"name": "seller_name", "critical": True},
            {"name": "image", "critical": True, "weight" : 3},
            {"name": "image_size"},
            {"name": "price", "critical": True, "weight" : 3},
            {"name": "line_price" },
            {"name": "was_price" },
            {"name": "unit_price" },
            {"name": "item_price" },
            {"name": "price_range" },
            {"name": "availability" },
            {"name": "product_location" },
            {"name": "flag" },
            {"name": "fulfillment" },
            {"name": "url", "critical": True, "weight" : 3},
        ]},
        {"name": "total_count", "critical": True, "weight" : 3},
        {"name": "search_query", "critical": True, "weight" : 2},
        {"name": "total_count_display"},
        {"name": "pagination", "critical": True, "weight" : 3, "children": [
            {"name": "page_count", "critical": True },
            {"name": "current_page", "critical": True },
            {"name": "page_links", "critical": True },
        ]},
        {"name": "takeover_tiles", "weight" : 2, "children": [
            {"name": "title" },
            {"name": "subtitle" },
            {"name": "image" },
            {"name": "site_link", "children": [
                {"name": "title"},
                {"name": "url"},
            ]},
        ]},
        {"name": "search_pills", "children": [
            {"name": "title" },
            {"name": "image" },
            {"name": "url" },
        ]},
        {"name": "search_banner", "children": [
            {"name": "title" },
            {"name": "image" },
            {"name": "url" },
            {"name": "description" },
            {"name": "button" },
        ]},
    ]}
]

# Walmart Product Fields
WALMART_PRODUCT_FIELDS = [
    {"name": "data", "critical": True, "children" : [
        {"name": "product", "critical": True, "weight" : 10, "children" : [
            {"name": "categories", "weight" : 3, "critical": True },
            {"name": "name", "weight" : 3, "critical": True },
            {"name": "brand", "critical": True},
            {"name": "brand_url" },
            {"name": "images", "weight" : 2, "critical": True },
            {"name": "thumbnail" },
            {"name": "average_rating", "critical": True },
            {"name": "number_of_reviews", "critical": True },
            {"name": "model"},
            {"name": "available_status", "critical": True},
            {"name": "sales_unit"},
            {"name": "price_info", "weight" : 2, "critical": True, "children" : [
                { "name": "price", "critical": True, "weight" : 3 },
                { "name": "price_string", "critical": True, "weight" : 3 },
                { "name": "was_price" },
                { "name": "unit_price" },
                { "name": "was_price" },
                { "name": "ship_price" },
                { "name": "list_price" },
                { "name": "comparison_price" },
                { "name": "savings_amount" },
                { "name": "price_range" },
                { "name": "additional_fees" },
            ]},
            {"name": "fulfillments"},
            {"name": "badges", "children" : [
                { "name": "flags" },
                { "name": "labels" },
                { "name": "groups" },
                { "name": "tags" },
            ]},
            {"name": "seller", "critical": True, "weight" : 2, "children" : [
                { "name": "name", "weight" : 3, "critical": True },
                { "name": "display_name" },
                { "name": "store_url" },
                { "name": "review_count" },
                { "name": "average_rating" },
            ]},
            { "name": "return_policy" },
            { "name": "location", "children" : [
                { "name": "postal_code" },
                { "name": "state_code" },
                { "name": "city" },
            ]},
            {"name": "item_id" },
            {"name": "offer_type" },
            {"name": "transactable_offer_count" },
            {"name": "interactive_product_video" },
        ]},
        {"name": "about", "critical": True, "weight" : 5,  "children": [
            {"name": "nutrition_information" },
            {"name": "product_details", "critical": True, "weight" : 3, "children": [
                { "name": "short_description", "critical": True },
                { "name": "long_description", "critical": True },    
            ] },
            { "name": "specifications", "weight" : 2, "critical": True },
            { "name": "warnings" },
            { "name": "directions" },
            { "name": "highlights" },
            { "name": "ingredients" },
            { "name": "warranty" },
            { "name": "videos" },
        ]},
        {"name": "related_search"},
        {"name": "review_infomation", "critical": True, "children": [
            {"name": "average_rating", "weight" : 3, "critical": True },
            {"name": "total_review_count", "weight" : 3, "critical": True },
            { "name": "total_media_count" },
            { "name": "5_star_rating" },
            { "name": "5_star_percent" },
            { "name": "4_star_rating" },
            { "name": "4_star_percent" },
            { "name": "3_star_rating" },
            { "name": "3_star_percent" },
            { "name": "2_star_rating" },
            { "name": "2_star_percent" },
            { "name": "1_star_rating" },
            { "name": "1_star_percent" },
            { "name": "top_negative_review", "weight" : 3, "critical": True },
            { "name": "top_positive_review", "weight" : 3, "critical": True },
            { "name": "customer_reviews", "weight" : 3, "critical": True },
        ]},
        {"name": "related_pages", "weight" : 2, "critical": True}
    ]}
]

# Walmart Review Fields
WALMART_REVIEW_FIELDS = [
    {"name": "data", "critical": True, "children" : [
        {"name": "reviews", "critical": True, "weight" : 10, "children" : [
            {"name": "title", "weight" : 3, "critical": True },
            {"name": "text", "weight" : 3, "critical": True },
            {"name": "time", "critical": True},
            {"name": "rating" },
            {"name": "username", "critical": True },
            {"name": "badges" },
            {"name": "media" },
        ]},
        {"name": "top_negative_review", "weight" : 2, "critical": True, "children" : [
            {"name": "title", "weight" : 3, "critical": True },
            {"name": "text", "weight" : 3, "critical": True },
            {"name": "time", "critical": True},
            {"name": "rating" },
            {"name": "username", "critical": True },
            {"name": "badges" },
            {"name": "media" },
        ]},
        {"name": "top_positive_review", "weight" : 2, "critical": True, "children" : [
            {"name": "title", "weight" : 3, "critical": True },
            {"name": "text", "weight" : 3, "critical": True },
            {"name": "time", "critical": True},
            {"name": "rating" },
            {"name": "username", "critical": True },
            {"name": "badges" },
            {"name": "media" },
        ]},
        {"name": "product", "weight" : 2, "critical": True, "children" : [
            {"name": "name", "weight" : 2, "critical": True },
            {"name": "url", "weight" : 2, "critical": True },
            {"name": "type"},
            {"name": "seller" },
            {"name": "categories", "weight" : 2, "critical": True },
        ]},
        {"name": "frequent_mentions"},
        {"name": "pagination", "critical": True, "children" : [
            {"name": "page_count", "critical": True },
            {"name": "current_page", "critical": True },
            {"name": "current_span", "critical": True},
            {"name": "page_links", "weight" : 2, "critical": True },
        ]},
        {"name": "average_rating", "critical": True},
        {"name": "total_review_count", "critical": True},
        {"name": "review_withtext_count"},
        {"name": "aspect_review_count"},
        {"name": "total_media_count"},
        {"name": "aspect_review_count"},
    ]}
]

# Google Search Fields
GOOGLE_SEARCH_FIELDS = [
    { "name": "search_information", "critical": True, "weight" : 2, "children" : [
        { "name": "total_results", "critical": True },
        { "name": "time_taken_displayed", "critical": True },
        { "name": "query_displayed", "critical": True },
    ]},
    { "name": "ads", "children" : [
        { "name": "position"},
        { "name": "block_position"},
        { "name": "title", "weight" : 3},
        { "name": "link", "weight" : 3},
        { "name": "thumbnail"},
        { "name": "displayed_link"},
        { "name": "sitelinks", "children" : [
            { "name": "inline", "children" : [
                { "name": "link" },
                { "name": "title" },
            ] },
            { "name": "block", "children" : [
                { "name": "description" },
                { "name": "link" },
                { "name": "title" },
            ]},
        ]},
    ]},
    { "name": "knowledge_graph", "weight" : 2, "children" : [
        { "name": "title", "weight" : 3 },
        { "name": "image", "weight" : 2 },
        { "name": "description", "weight" : 3 },
        { "name": "source" },
        { "name": "related" },
        { "name": "related_link" },
        { "name": "social_media" },
        { "name": "see_more_about" },
    ]},
    { "name": "related_questions", "critical": True, "weight" : 2 },
    { "name": "organic_results", "critical": True, "weight" : 10, "children" : [
        { "name": "position" },
        { "name": "title", "critical": True, "weight" : 2 },
        { "name": "snippet", "critical": True, "weight" : 2 },
        { "name": "link", "critical": True, "weight" : 2 },
        { "name": "date" },
        { "name": "displayed_link" },
        { "name": "thumbnail" },
        { "name": "site_links" , "children" : [
            { "name": "inline", "children" : [
                { "name": "link" },
                { "name": "title" },
            ] },
            { "name": "block", "children" : [
                { "name": "description" },
                { "name": "link" },
                { "name": "title" },
            ]},
        ]},
    ]},
    { "name": "related_searches", "critical": True },
    { "name": "more", "critical": True },
]

def is_empty(value):
    if value is None:
        return True
    elif isinstance(value, int):
        return value == 0
    elif isinstance(value, float):
        return value == 0.0
    elif isinstance(value, str):
        return value == ""
    elif isinstance(value, list):
        return len(value) == 0
    elif isinstance(value, dict):
        return len(value.keys()) == 0
    else:
        return False
    
def calculate_field(field, record): 
    name = field.get("name", "unknown")
    children = field.get("children", [])
    overall_score = 0.0
    critical_score = 0.0
    if record.get(name, None) != None:
        data = record[name]
        if field.get("children", None) != None:
            if isinstance(data, list):
                critical_total_score = 0.0
                overall_total_score = 0.0
                for subdata in data:
                    overall, critical = calculate_fields(children, subdata)   
                    overall_total_score += overall
                    critical_total_score += critical
                overall_score = overall_total_score / len(data) if len(data) > 0 else 0
                critical_score = critical_total_score / len(data) if len(data) > 0 else 0
            else:
                overall_score, critical_score = calculate_fields(children, data)
        else:
            if not is_empty(data):
                overall_score = 1.0
                critical_score = 1.0
    return overall_score, critical_score

def calculate_fields(fields, record): 
    critical_total_max = 0.0
    critical_total_sum = 0.0
    overall_total_max = 0.0
    overall_total_sum = 0.0
    for childField in fields:
        overall, critical = calculate_field(childField, record)
        overall_total_max += childField.get("weight", 1)
        overall_total_sum += overall * childField.get("weight", 1)
        critical_total_max += childField.get("weight", 1) if childField.get("critical", False) else 0.0
        critical_total_sum += critical * childField.get("weight", 1) if childField.get("critical", False) else 0.0
        if childField.get("critical", False):
            print(f"*** {overall:.3f}: {critical:.3f} : {childField.get("name", "unknown")} * {childField.get("weight", 1)} : {overall_total_max:.3f} : {overall_total_sum:.3f}: {critical_total_max:.3f} : {critical_total_sum:.3f}")
        else:
            print(f"--- {overall:.3f}: {critical:.3f} : {childField.get("name", "unknown")} * {childField.get("weight", 1)} : {overall_total_max:.3f} : {overall_total_sum:.3f}: {critical_total_max:.3f} : {critical_total_sum:.3f}")
    overall_score = overall_total_sum / overall_total_max
    critical_score = critical_total_sum / critical_total_max if critical_total_max > 0 else 1.0
    print(f"### : {overall_score:.3f}: {critical_score:.3f}")
    return overall_score, critical_score

def calculate_score(result):
    if CURRENT_TEST == "walmart-search" :
        return calculate_fields(WALMART_SEARCH_FIELDS, result)
    elif CURRENT_TEST == "walmart-product" :
        return calculate_fields(WALMART_PRODUCT_FIELDS, result)
    elif CURRENT_TEST == "walmart-review" :
        return calculate_fields(WALMART_REVIEW_FIELDS, result)
    elif CURRENT_TEST == "google-search" :
        return calculate_fields(GOOGLE_SEARCH_FIELDS, result)
    else :
        return 0.0, 0.0
    
def save_result(result):
    with open("../data/result.json", "w", encoding="utf-8") as f:
        json.dump(result, fp=f, indent="\t")
    pass

def extract_sample_urls():
    response = requests.post("https://proxy.scrapeops.io/v1/sample-urls", json={
        "api_key": API_KEY_URL,
        "url": TEST_CONFIG[CURRENT_TEST]["url"],
        "domain": TEST_CONFIG[CURRENT_TEST]["domain"]
    })
    if response.status_code != 200:
        print(response.status_code)
        return None
    sample_urls = response.json().get("data", [])
    return sample_urls

def gather_html(url):
    response = requests.get('https://proxy.scrapeops.io/v1/',params={
        'api_key': API_KEY_PROXY,
        'url': url, 
    })
    if response.status_code != 200:
        print(f"ERROR : failed to gather html")
        return None
    return response.text

def parse_html(html):
    try:
        resp = requests.post(API_ENDPOINT_PARSE, json={
            "url": "url",
            "html": html
        })        
        if resp.status_code != 200:
            print(f"ERROR : failed to parse html")
            return None
        return resp.json()
    except Exception:
        print(f"ERROR : failed to parse html")
        return None

def write_log(log):
    with open("log.txt", '+a', encoding='utf-8') as f:
        print(f"{datetime.datetime.now()} : {log}", file=f)

# get sample urls

sample_urls = TEST_URLS
# sample_urls = extract_sample_urls()
if sample_urls is None:
    print("ERROR : failed to extract sample urls")
    exit()
# gather, parse, estimate html
count = 0
overall_scores = np.array([])
critical_scores = np.array([])

for url in sample_urls:
    html = gather_html(url["url"])
    if html is None:
        continue
    count += 1
    result = parse_html(html)
    save_result(result)
    if result is None:
        overall_scores = np.append(overall_scores, 0)
        critical_scores = np.append(critical_scores, 0)
        print(f"{count} : 0.00 : 0.00 : {url["url"]}")
    else:
        overall_score, critical_score = calculate_score(result)
        overall_scores = np.append(overall_scores, overall_score)
        critical_scores = np.append(critical_scores, critical_score)
        print(f"{count} : {overall_score:.3f} : {critical_score:.3f} : {url["url"]}")
    
# with open("./data/walmart-search.json", "r", encoding="utf-8") as f:
#     result = json.load(f)
# calculate_score(result)