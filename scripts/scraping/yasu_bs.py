import json
import os
import re
from typing import List, Tuple

import fire
import httpx
import requests
from bs4 import BeautifulSoup

""""
type Subject struct {
	id				 int64 
	category		 string 
	title			 string 
	facultyIds		 []int64
	videoIds		 []int64
	location		 string
	remark 			 string
	pdfLinks		 []string
	relatedSubjectIds []int64
	department       string
	language         string
	academicYear     string
	semester         string
	firstHeldOn      string
	numCredit       int8
	courceFormat     string
	targetedAudience string
	dayOfWeek        string
	courcePeriod     string
	outline		  string
	objective	  string
	lessonPlan	  string
	gradingMethod	  string
	courceRequirement string
	reference	  string
	subpageIds	  []int64
}

type Chapter struct {
	start_at time.Time
	topic	string
	thumbnail string
}

type subpage struct {
	id				 int64
	content			 string
}

type Video struct {
	id				 int64
	title			 string
	link			 string
	chapters		 []*Chapter
	facultyIds		 []int64
	lecturedOn time.Time
	videoLength time.Duration
	Language		 string
}

type Faculty struct {
	id				 int64
	name			 string
	department		 string
	rank			 string
}
"""
class Page:
    url = None
    soup = None

    def __init__(self, url: str):
        self.url = url
        self.soup = BeautifulSoup(requests.get(url).text, "html.parser")

    # for subject struct
    def get_subject_title(self) -> str:
        # select class with "c-title__content"
        return self.soup.find('h2', class_='c-title__content').text
        
    # find a tag with text "開講部局名" and select the next tag
    def get_department(self) -> str:
        return self.soup.find('dt', string='開講部局名').find_next('dd').text

    def get_language(self) -> str:
        return self.soup.find('dt', string='使用言語').find_next('dd').text
    
    def get_academic_year(self) -> int:
        year_regex = re.compile(r'^\d{4}')
        try:
            year = self.soup.find('th', string='開講年度・開講期').find_next('td').text
        except AttributeError:
            try:
                year = self.soup.find('th', string='年度').find_next('td').text
            except AttributeError:
                year = self.soup.find('th', string='年度・期').find_next('td').text
        
        matched_year = year_regex.match(year)
        
        return int(matched_year.group())

    def get_semester(self) -> str:
        try:
            year = self.soup.find('th', string='開講年度・開講期').find_next('td').text
        except AttributeError:
            year = self.soup.find('th', string='開講期').find_next('td').text

        # regex for 前期 or 後期
        semester_regex = re.compile(r'(前期|後期)')
        matched_semester = semester_regex.search(year)
        return matched_semester.group()

    def get_targeted_audience(self) -> str:
        return self.soup.find('th', string='対象学生').find_next('td').text

    def get_subject_outline(self) -> str:
        
        text = self.soup.find('th', string='授業の概要・目的').find_next('div').text
        # remove all special characters and spaces from text
        return re.sub(r'[^\w]', '', text)

subject_dict = {
    "id": 0,
    "category": "",
    "title": "",
    "facultyIds": [],
    "videoIds": [],
    "location": "",
    "remark": "",
    "pdfLinks": [],
    "relatedSubjectIds": [],
    "department": "",
    "language": "",
    "academicYear": "",
    "semester": "",
    "firstHeldOn": "",
    "numCredit": 0,
    "courceFormat": "",
    "targetedAudience": "",
    "dayOfWeek": "",
    "courcePeriod": "",
    "outline": "",
    "objective": "",
    "lessonPlan": "",
    "gradingMethod": "",
    "courceRequirement": "",
    "reference": "",
    "subpageIds": [],
}

faculty_dict = {
    "id": 0,
    "name": "",
    "department": "",
    "rank": "",
}


chapter_dict = {
    "start_at": "",
    "topic": "",
    "thumbnail": "",
}

subpage_dict = {
    "id": 0,
    "content": "",
}

video_dict = {
    "id": 0,
    "title": "",
    "link": "",
    "chapters": [],
    "facultyIds": [],
    "lecturedOn": "",
    "videoLength": "",
    "Language": "",
}


def fetch_body(url: str) -> str:
    res = httpx.get(url)
    soup = BeautifulSoup(res.text, "html.parser")
    data = soup.select_one("#__NEXT_DATA__").get_text()
    print(data)
    return json.loads(data)["props"]["pageProps"]["tasks"]["body"]


def download_images_and_replace_links(text: str, question_number) -> str:
    soup = BeautifulSoup(text, "html.parser")
    dir_name = f"questions/{str(question_number)}"
    # crate image folder
    os.makedirs(f"{os.path.dirname(__file__)}/../../{dir_name}/images", exist_ok=True)
    regex = r"(http|ftp|https):\/\/([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:\/~+#-]*[\w@?^=%&\/~+#-])"
    all_links = re.findall(regex, text)
    print(all_links)
    # filter out links that does not include png file
    all_png_links = [link for link in all_links if link[1].endswith(".png")]
    # download images from links
    for link in all_png_links:
        # download png file with requests
        # TODO: check if this line can download image with requests
        r = requests.get(link[0] + "://" + link[1] + link[2])
        with open(f"{os.path.dirname(__file__)}/../../{dir_name}/images/{link[1]}", "wb") as f:
            f.write(r.content)
        # replace link with local path
        text = text.replace(link[0] + "://" + link[1] + link[2], f"/{dir_name}/images/{link[1]}")
    return text

    # get all patterns

    return soup.prettify()


if __name__ == "__main__":
    # fire.Fire(get_question)

    url = "https://ocw.kyoto-u.ac.jp/course/68/"
    res = httpx.get(url)
    soup = BeautifulSoup(res.text, "html.parser")
    title = soup.title.text
    print(title)

    # select element with css selector '.c-title__content'
    subject_title = soup.select(".c-title__content")[0].text  # 細胞内情報発信学
