import unittest
from urllib.request import urlopen

from bs4 import BeautifulSoup

from yasu_bs import *


class Test(unittest.TestCase):
    bs = None

    def setUpClass():
        url = "https://ocw.kyoto-u.ac.jp/course/68/"
        Test.page = Page(url)

    def test_subject(self):
        subject_title = Test.page.get_subject_title()
        self.assertEqual("細胞内情報発信学", subject_title)

    def test_department(self):
        department = Test.page.get_department()
        self.assertEqual("理学部", department)

    def test_language(self):
        language = Test.page.get_language()
        self.assertEqual("日本語", language)

    def test_academic_year(self):
        academic_year = Test.page.get_academic_year()
        self.assertEqual("2018", academic_year)

    def test_semester(self):
        semester = Test.page.get_semester()
        self.assertEqual("前期", semester)

    def test_targeted_audience(self):
        targeted_audience = Test.page.get_targeted_audience()
        self.assertEqual("３回生以上", targeted_audience)

    def test_subject_outline(self):
        subject_outline = Test.page.get_subject_outline()
        self.assertEqual(
            """生体および細胞にとって恒常性維持は生存に不可欠である。通常とは異なる状況が生じたときに細胞は，異常を細胞内に存在するセンサーによって感知し，核へ情報を伝達して遺伝子発現を調節することによって適応する仕組みを備えている。このような細胞内情報発信では，レセプターを介した細胞表面からの情報伝達の場合とは全く異なる分子機構が用いられている。本講義では，低酸素，コレステロール濃度低下，分泌系タンパク質の構造異常（小胞体ストレス）などの異常事態に対して細胞が，どのように変化を感知し，どのように核へ向けて情報を伝達し，核内でどのような遺伝子発現を促して恒常性を維持するのか解説する。

            第1回 低酸素応答の概要と発見の経緯を解説する。
            第2回 低酸素応答に関与するシス配列がゲノムのどこに存在するか明らかにする方法を解説する。
            第3回 低酸素応答に関与するシス配列を絞り込む方法を解説する。
            第4回 低酸素応答に関与するシス配列に結合するトランス因子を同定し、精製する方法を解説する。
            第5回 低酸素応答に関与するトランス因子をクローニングする方法を解説する。
            第6回 低酸素応答に関与するトランス因子の活性化機構を解説する。
            第7回 低酸素応答に関与するトランス因子の活性化に関与する酵素を同定する方法を解説する。
            第8回 低酸素という曖昧な情報を細胞がどのように感知しているか解説する。
            第9回 コレステロール枯渇応答の概要、発見の経緯、応答に関与するシス配列を絞り込む方法を解説する。
            第10回 コレステロール枯渇応答に関与するトランス因子を同定・精製し、クローニングする方法を解説する。
            第11回 コレステロール枯渇応答に関与するトランス因子の活性化方法を解説する。
            第12回 コレステロール枯渇応答に関与するトランス因子の活性化に関与する酵素を同定する方法を解説する。
            第13回 コレステロール枯渇という曖昧な情報を細胞がどのように感知しているか解説する。

            ※より詳しい内容の講義は「生物学セミナーB」の「小胞体ストレス応答の概要と発見の経緯を解説する。」に続きます。
            """,
            subject_outline,
        )


if __name__ == "__main__":
    unittest.main()
