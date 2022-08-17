# The above code is scraping the Kyoto University OCW website to get the list of courses in each
# series.import json
import json
import re

import requests
from bs4 import BeautifulSoup, SoupStrainer

if __name__ == "__main__":

    # get jp version

    series_dic = {}
    links = []
    series_titles = []

    url = "https://ocw.kyoto-u.ac.jp/series"
    # get text inside <dt> tag
    soup = BeautifulSoup(requests.get(url).text, "html.parser")
    for dt in soup.find_all("dt"):
        # get text inside <dt> tag
        series_titles.append(dt.text)

    # get all links from the page
    for link in BeautifulSoup(requests.get(url).text, "html.parser", parse_only=SoupStrainer("a")):
        if link.has_attr("href"):
            links.append(link["href"])

    links = [link for link in links if "series=" in link]

    # get strings after "series="
    series_strings = [link.split("series=")[1] for link in links]
    series_strings_to_series_title_dic = {
        series_string: series_title for series_string, series_title in zip(series_strings, series_titles)
    }
    # print(series_strings_to_series_title_dic)

    series_dic = {key: [] for key in series_strings}

    full_links = ["https://ocw.kyoto-u.ac.jp/" + link for link in links]

    for i, series_string in enumerate(series_strings):
        top_link = f"https://ocw.kyoto-u.ac.jp/?s=&faculty=&category=&series={series_string}&subject=&year="
        soup = BeautifulSoup(requests.get(top_link).text, "html.parser")
        search_result = soup.find("h2", {"class": "c-title__content--sub"})
        # get number inside parentheses with regex
        num_videos = int(re.findall(r"\d+", search_result.text)[0])
        num_pages = num_videos // 16 + 1 if num_videos % 16 != 0 else num_videos // 16
        # print(num_pages)

        subject_links = []
        for n_page in range(1, num_pages + 1):
            link = (
                f"https://ocw.kyoto-u.ac.jp/page/{n_page}?s=&faculty=&category=&series={series_string}&subject=&year="
            )
            # print(link)
            # get all links from the page
            for link in BeautifulSoup(requests.get(link).text, "html.parser", parse_only=SoupStrainer("a")):
                if link.has_attr("href"):
                    # if link include "/course/{number}"
                    if "/course/" in link["href"]:
                        # get only numbers with regex
                        number = re.findall(r"\d+", link["href"])
                        if len(number) == 1:
                            series_dic[series_strings[i]].append(number[0])

    # print(series_dic)
    # get total number of each category
    # for key, value in series_dic.items():
    #    print(key, len(value))
    # export series_string_to_series_title_dic to json file
    with open("data/series_jp_string_to_series_title_dic.json", "w") as f:
        json.dump(series_strings_to_series_title_dic, f)

    # export series_dic to json file
    with open("data/series_jp_dic.json", "w") as f:
        json.dump(series_dic, f)

    # Get English version

    series_en_dic = {}
    links = []
    series_en_titles = []

    url_en = "https://ocw.kyoto-u.ac.jp/en/series"
    # get text inside <dt> tag
    soup = BeautifulSoup(requests.get(url_en).text, "html.parser")
    for dt in soup.find_all("dt"):
        # get text inside <dt> tag
        series_en_titles.append(dt.text)
    print(series_en_titles)

    # get all links from the page
    for link in BeautifulSoup(requests.get(url_en).text, "html.parser", parse_only=SoupStrainer("a")):
        if link.has_attr("href"):
            links.append(link["href"])

    links = [link for link in links if "series=" in link]
    # print(links)
    # get strings after "series="
    series_en_strings = [link.split("series=")[1] for link in links]
    series_en_strings_to_series_title_dic = {
        series_string: series_title for series_string, series_title in zip(series_en_strings, series_en_titles)
    }
    print(series_en_strings_to_series_title_dic)

    series_en_dic = {key: [] for key in series_en_strings}

    full_links = ["https://ocw.kyoto-u.ac.jp/" + link for link in links]

    for i, series_string in enumerate(series_en_strings):
        top_link = f"https://ocw.kyoto-u.ac.jp/en/?s=&faculty=&category=&series={series_string}&subject=&year="
        soup = BeautifulSoup(requests.get(top_link).text, "html.parser")
        search_result = soup.find("h2", {"class": "c-title__content--sub"})
        # get number inside parentheses with regex
        num_videos = int(re.findall(r"\d+", search_result.text)[0])
        num_pages = num_videos // 16 + 1 if num_videos % 16 != 0 else num_videos // 16
        # print(num_pages)

        subject_links = []
        for n_page in range(1, num_pages + 1):
            link = f"https://ocw.kyoto-u.ac.jp/en/page/{n_page}?s=&faculty=&category=&series={series_string}&subject=&year="
            # print(link)
            # get all links from the page
            for link in BeautifulSoup(requests.get(link).text, "html.parser", parse_only=SoupStrainer("a")):
                if link.has_attr("href"):
                    # if link include "/course/{number}"
                    if "/course/" in link["href"]:
                        # get only numbers with regex
                        number = re.findall(r"\d+", link["href"])
                        if len(number) == 1:
                            series_en_dic[series_en_strings[i]].append(number[0])

    # print(series_dic)
    # get total number of each category
    # for key, value in series_dic.items():
    #    print(key, len(value))
    # export series_string_to_series_title_dic to json file
    with open("data/series_en_string_to_series_title_dic.json", "w") as f:
        json.dump(series_en_strings_to_series_title_dic, f)

    # export series_dic to json file
    with open("data/series_en_dic.json", "w") as f:
        json.dump(series_en_dic, f)
