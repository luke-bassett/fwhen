from fwhen.scrape import F1DotComScraper

with open('race_pages.txt', 'r') as f:
    race_pages = f.read().splitlines()

for race_page in race_pages:
    print(race_page)
    scraper = F1DotComScraper(race_page)
    scraper.request_html()
    scraper.find_session_times()
    with open('2022_sessions.csv', 'a') as f:
        for session in scraper.session_times.keys():
            if scraper.session_times[session]:
                f.write(f"{race_page},{session},{scraper.session_times[session][0]},{scraper.session_times[session][1]}\n")



