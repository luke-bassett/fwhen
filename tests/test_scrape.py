from fwhen.scrape import F1DotComScraper


def test_find_session_times():
    s = F1DotComScraper("www.test.com")

    with open("tests/sample.html", "r") as f:
        s.html = f.read()
    s.find_session_times()
    assert s.sessions["fp1"] == ("2022-10-28T18:00:00Z", "2022-10-28T19:00:00Z")
    assert s.sessions["fp2"] == ("2022-10-28T21:00:00Z", "2022-10-28T22:00:00Z")
    assert s.sessions["fp3"] == ("2022-10-29T17:00:00Z", "2022-10-29T18:00:00Z")
