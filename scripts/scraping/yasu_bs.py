import json
import os
import re
from typing import List, Tuple

import fire
import httpx
import requests
from bs4 import BeautifulSoup

import ulid

# define python dictionaries for each go struct

dic_video = {
    "id": "",
    "title": "",
    "link": "",
    "chapters": [],
    "facultyIds": [],
    "lecturedOn": "",
    "videoLength": "",
    "langugage": "",
}

dic_pdf = {
    "id": "",
    "title": "",
    "description": "",
    "url": "",
}

# convert the following struct to python
dic_syllabus = {
    "id": "",
    "language": "",
    "subjectNumbering": "",
    "academicYear": "",
    "semester": "",
    "numCredit": "",
    "courseFormat": "",
    "assignmenedGrade": "",
    "targettedAudience": "",
    "dayOfWeek": "",
    "coursePeriod": "",
    "outline": "",
    "objective": "",
    "lessonPlan": "",
    "gradingMethod": "",
    "courseRequirement": "",
    "outClassLearning": "",
    "outClassLearnig": "",
    "reference": "",
    "reference": "",
    "remark": "",
    "subpageIds": [],
}

dic_faculty = {
    "id": "",
    "name": "",
    "department": "",
    "rank": ""
}

dic_subpage = {
    "id": "",
    "content": ""
}

def get_subject_attributes(subject_url: str) -> dict:


    dic_subject = {
        "id": "",
        "category": "",
        "title": "",
        "videoIds": [],
        "location": "",
        "pdfIds": [],
        "relatedSubjectIds": [],
        "department": "",
        "firstHeldOn": "",
        "facultyIds": [],
        "language": "",
        "freeDescription": "",
        "syllabusId": "",
        "series": "",
    }


class Subject():
    def __init__(self, url: str):
        self.url = url
        self.soup = BeautifulSoup(requests.get(url).text, "html.parser")

    # for subject struct
    def get_subject_title(self) -> str:
        # select class with "c-title__content"
        return self.soup.find('h2', class_='c-title__content').text

class Video():
    def __init__(self, url: str):
        self.url = url
    


class Syllabus():
    url = None
    soup = None

    def __init__(self, url: str):
        self.url = url
        self.soup = BeautifulSoup(requests.get(url).text, "html.parser")

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


if __name__ == "__main__":
    # fire.Fire(get_question)

    url = "https://ocw.kyoto-u.ac.jp/course/68/"
    res = httpx.get(url)
    soup = BeautifulSoup(res.text, "html.parser")
    title = soup.title.text
    print(title)

    # select element with css selector '.c-title__content'
    subject_title = soup.select(".c-title__content")[0].text  # 細胞内情報発信学
