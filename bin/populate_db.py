import sqlite3
import pandas as pd

conn = sqlite3.connect('data/fwhen.db')

cur = conn.cursor()

cur.execute("""create table if not exists events(
    id integer primary key,
    location_name text,
    year int
    );"""
)
cur.execute("""create table if not exists sessions(
    id integer primary key,
    eventid integer not null,
    session_type text not null,
    start_time text not null,
    end_time text not null,
    foreign key (eventid) references events(id)
    );"""
)
conn.commit()

sessions_df = pd.read_csv('data/2022_sessions.csv')
location_names = sessions_df['url'].str.extract(r'.*/(.*)\.html').unique().tolist()
print(location_names)
for location_name in location_names:
    cur.execute(f"insert into events values({location_name}, 2022);")

