from flask import Flask, render_template
import json

app = Flask(__name__)


@app.route("/")
def home():
    with open("data/sessions.json", "r") as f:
        sessions = json.loads(f.read())
    return render_template("home.html", sessions=sessions)
