import json
import re

import requests
from bs4 import BeautifulSoup, SoupStrainer

if __name__ == "__main__":

    categories = []
    category_code_to_category_name_dic = {}

    academic_fields_dic = {}
    links = []
    series_titles = []
    academi_field_names = []

    url = "https://ocw.kyoto-u.ac.jp/course/"

    # select tag with name == "category"
    soup = BeautifulSoup(requests.get(url).text, "html.parser")
    category_tag = soup.find_all("select", {"name": "category"})[0]

    for child in category_tag.findChildren("option"):

        # get value of option tag
        value = child.get("value")

        if len(value) > 0:
            category_name = child.text
            categories.append(value)
            print(value, category_name)
            category_code_to_category_name_dic[value] = category_name

full_links = [
    f"https://ocw.kyoto-u.ac.jp/?s=&faculty=&category={category}&series=&subject=&year=" for category in categories
]

categories_dic = {category: [] for category in categories}
num_pages = [22, 30, 6, 6, 6]  # 手動

for i, category in enumerate(categories):

    num_page = num_pages[i]
    for n_page in range(1, num_page + 1):
        link = f"https://ocw.kyoto-u.ac.jp/page/{n_page}?s=&faculty=&category={category}&series=&subject=&year="

        for link in BeautifulSoup(requests.get(link).text, "html.parser", parse_only=SoupStrainer("a")):

            if link.has_attr("href"):
                # if link include "/course/{number}"
                if "/course/" in link["href"]:
                    # get only numbers with regex
                    number = re.findall(r"\d+", link["href"])
                    if len(number) == 1:
                        categories_dic[category].append(number[0])

# print(categories_dic)
# count all courses in each category
# for category, courses in categories_dic.items():
#    print(f"{category}: {len(courses)}")

# export category_code_to_category_name_dic and categoeis_dic to json file
with open("data/category_code_to_category_name_dic.json", "w") as f:
    json.dump(category_code_to_category_name_dic, f)

with open("data/categories_dic.json", "w") as f:
    json.dump(categories_dic, f)
