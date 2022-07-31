import json
import pandas as pd

df = pd.read_csv("data/2022_sessions_data.csv")

locs = df["location"].unique().tolist()

session_data = {}
for loc in locs:
    session_data[loc] = {}
    for i, row in df[df["location"] == loc].iterrows():
        session_data[loc][row.session] = [row.start, row.end]

with open("data/sessions.json", "w") as f:
    f.write(json.dumps(session_data, indent=2))
