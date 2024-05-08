import requests
import numpy as np
import datetime

CURRENT_TEST = "walmart-search"
TEST_CONFIG = {
    "walmart-search": {
        "url": "*walmart.com/search*",
        "domain": "walmart.com",
        "endpoint": "walmart",
    }
}
API_KEY_URL = "81a3edae-ece8-47ed-a672-13e7961a6fcc"
API_KEY_PROXY = "d0f35644-3095-48f9-b109-9146894ae581"

WALMART_SEARCH_FIELDS = [
    {"name": "data", "critical": True, "weight" : 5, "children" : [
        {"name": "results", "critical": True, "weight" : 5, "children" : [
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
    ]},
    {"name": "status", "critical": True},
    {"name": "url", "critical": True},
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
        # print(f"# {childField.get("name", "unknown")} * {childField.get("weight", 1)}: {overall}: {critical} : {overall_total_max} : {overall_total_sum}: {critical_total_max} : {critical_total_sum}")
    overall_score = overall_total_sum / overall_total_max
    critical_score = critical_total_sum / critical_total_max if critical_total_max > 0 else 1.0
    # print(f"--- : {overall_score}: {critical_score}")
    return overall_score, critical_score

def calculate_score(result):
    if CURRENT_TEST == "walmart-search" :
        return calculate_fields(WALMART_SEARCH_FIELDS, result)
    else:
        return 0.0, 0.0

def save_result(name):
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
        resp = requests.post(f"http://localhost:8080/{TEST_CONFIG[CURRENT_TEST]["endpoint"]}", json={
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
sample_urls = extract_sample_urls()
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
    if result is None:
        overall_scores = np.append(overall_scores, 0)
        critical_scores = np.append(critical_scores, 0)
        print(f"{count} : 0.00 : 0.00 : {url["url"]}")
    else:
        overall_score, critical_score = calculate_score(result)
        overall_scores = np.append(overall_scores, overall_score)
        critical_scores = np.append(critical_scores, critical_score)
        print(f"{count} : {overall_score:.2f} : {critical_score:.2f} : {url["url"]}")
    break
# with open("./data/walmart-search-cake.json", "r", encoding="utf-8") as f:
#     result = json.load(f)
# calculate_score(result)