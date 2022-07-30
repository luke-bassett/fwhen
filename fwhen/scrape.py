"""Impliments the scraper class for collection schedule information"""

from datetime import datetime
import re
import requests
from typing import Dict, Tuple

SESSIONS = ["Practice 1", "Practice 2", "Practice 3", "Sprint", "Qualifying", "Race"]


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

    @staticmethod
    def get_startdate_pattern(session_name: str) -> str:
        return r'(?sm)"name": "' + session_name + r'.*?startDate": "(\S*)"'

    @staticmethod
    def get_enddate_pattern(session_name: str) -> str:
        return r'(?sm)"name": "' + session_name + r'.*?endDate": "(\S*)"'

    def get_session_times(self, session_name: str) -> Tuple[str, str]:
        start_date = re.findall(self.get_startdate_pattern(session_name), self.html)[0]
        end_date = re.findall(self.get_enddate_pattern(session_name), self.html)[0]
        return start_date, end_date

    def find_session_times(self) -> Dict[str, datetime]:
        self.session_times = {}
        for session in SESSIONS:
            try:
                self.session_times[session] = self.get_session_times(session)
            except IndexError:
                self.session_times[session] = None
