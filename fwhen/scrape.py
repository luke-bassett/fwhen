"""Impliments the scraper class for collection schedule information"""

from datetime import datetime
import re
import requests
from typing import Dict


class Scraper:
    """Contains basic html collection and parsing methods, 1 object per url."""

    def __init__(self, url: str):
        self.url = url

    def request_html(self):
        self.response = requests.get(self.url)
        self.html = self.response.text


class F1DotComScraper(Scraper):
    def __init__(self, url):
        super().__init__(url)

    def find_session_times(self) -> Dict[str, datetime]:
        self.sessions = {}
        for i in range(1, 4):
            start_pattern = f'(?sm)Practice {i}.*?startDate": "(\S*)"'
            end_pattern = f'(?sm)Practice {i}.*?endDate": "(\S*)"'
            start_date = re.findall(start_pattern, self.html)[0]
            end_date = re.findall(end_pattern, self.html)[0]
            self.sessions[f"fp{i}"] = (start_date, end_date)
