import json
import re

import requests
from bs4 import BeautifulSoup, SoupStrainer

if __name__ == "__main__":

    series_dic = {}
    links = []

    url = "https://ocw.kyoto-u.ac.jp/series"
    # get all links from the page
    for link in BeautifulSoup(requests.get(url).text, "html.parser", parse_only=SoupStrainer("a")):
        if link.has_attr("href"):
            links.append(link["href"])

    links = [link for link in links if "series=" in link]

    # get strings after "series="
    series_strings = [link.split("series=")[1] for link in links]
    series_dic = {key: [] for key in series_strings}

    full_links = ["https://ocw.kyoto-u.ac.jp/" + link for link in links]

    for i, link in enumerate(full_links):
        subject_links = []
        # get all links from the page
        for link in BeautifulSoup(requests.get(link).text, "html.parser", parse_only=SoupStrainer("a")):
            if link.has_attr("href"):
                # if link include "/course/{number}"
                if "/course/" in link["href"]:
                    # get only numbers with regex
                    number = re.findall(r"\d+", link["href"])
                    if len(number) == 1:
                        series_dic[series_strings[i]].append(number[0])

    # export series_dic to json file
    with open("data/series_dic.json", "w") as f:
        json.dump(series_dic, f)
