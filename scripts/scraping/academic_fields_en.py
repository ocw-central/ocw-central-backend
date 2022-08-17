import json
import re

import requests
from bs4 import BeautifulSoup, SoupStrainer

if __name__ == "__main__":

    academic_fields_dic = {}
    links = []
    fields = []
    academi_field_names = []
    fields_to_filed_names_dic = {}

    url = "https://ocw.kyoto-u.ac.jp/en/course/"

    # select tag with name == "category"
    soup = BeautifulSoup(requests.get(url).text, "html.parser")
    field_tag = soup.find_all("select", {"name": "subject"})[0]

    for child in field_tag.findChildren("option"):

        # get value of option tag
        value = child.get("value")

        if len(value) > 0:
            field_name = child.text
            fields.append(value)
            print(value, field_name)
            fields_to_filed_names_dic[value] = field_name

academic_fields_dic = {field: [] for field in fields}

for i, field in enumerate(fields):

    top_link = f"https://ocw.kyoto-u.ac.jp/en/?s=&faculty=&category=&series=&subject={field}&year="
    #    print(top_link)
    soup = BeautifulSoup(requests.get(top_link).text, "html.parser")
    search_result = soup.find("h2", {"class": "c-title__content--sub"})
    # get number inside parentheses with regex
    num_videos = int(re.findall(r"\d+", search_result.text)[0])
    # print(num_videos)
    num_page = num_videos // 16 + 1 if num_videos % 16 != 0 else num_videos // 16
    # print(num_page)

    for n_page in range(1, num_page + 1):
        link = f"https://ocw.kyoto-u.ac.jp/en/page/{n_page}?s=&faculty=&category=&series=&subject={field}&year="

        for link in BeautifulSoup(requests.get(link).text, "html.parser", parse_only=SoupStrainer("a")):

            if link.has_attr("href"):
                # if link include "/course/{number}"
                if "/course/" in link["href"]:
                    # get only numbers with regex
                    number = re.findall(r"\d+", link["href"])
                    if len(number) == 1:
                        academic_fields_dic[field].append(number[0])

print(academic_fields_dic)
# get total number of videos in each category
for field, courses in academic_fields_dic.items():
    print(f"{field}: {len(courses)}")

print(fields_to_filed_names_dic)

# print(categories_dic)
# count all courses in each category
# for category, courses in categories_dic.items():
#    print(f"{category}: {len(courses)}")

# export category_code_to_category_name_dic and categoeis_dic to json file
with open("data/academic_fields_en_code_to_names.json", "w") as f:
    json.dump(fields_to_filed_names_dic, f)

with open("data/academic_fields_en_dic.json", "w") as f:
    json.dump(academic_fields_dic, f)
