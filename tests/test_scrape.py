from fwhen.scrape import F1DotComScraper


def test_find_session_times():
    s = F1DotComScraper("www.test.com")

    with open("tests/sample.html", "r") as f:
        s.html = f.read()
    s.find_session_times()
    assert s.session_times["Practice 1"] == (
        "2022-10-28T18:00:00Z",
        "2022-10-28T19:00:00Z",
    )
    assert s.session_times["Practice 2"] == (
        "2022-10-28T21:00:00Z",
        "2022-10-28T22:00:00Z",
    )
    assert s.session_times["Practice 3"] == (
        "2022-10-29T17:00:00Z",
        "2022-10-29T18:00:00Z",
    )
    assert s.session_times["Sprint"] == None
    assert s.session_times["Qualifying"] == (
        "2022-10-29T20:00:00Z",
        "2022-10-29T21:00:00Z",
    )
    assert s.session_times["Race"] == ("2022-10-30T20:00:00Z", "2022-10-30T22:00:00Z")
