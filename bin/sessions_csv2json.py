"""Converts CSV to json, text for frontend taken from json"""
import json
import pandas as pd

df = pd.read_csv("data/2022_sessions_data.csv")
df['startdate'] = pd.to_datetime(df['start'])
df = df.sort_values('startdate', ascending=True)

df['location'] = df['location'].str.lower()
locs = df["location"].unique().tolist()

df['session'] = df['session'].str.replace('Practice ', 'fp')
df['session'] = df['session'].str.replace('Qualifying', 'quali ')
df['session'] = df['session'].str.lower()


session_data = {}
for loc in locs:
    session_data[loc] = {}
    for i, row in df[df["location"] == loc].iterrows():
        session_data[loc][row.session] = [row.start, row.end]


with open("data/sessions.json", "w") as f:
    f.write(json.dumps(session_data, indent=2))
