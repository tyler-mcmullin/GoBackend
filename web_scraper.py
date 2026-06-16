from bs4 import BeautifulSoup
from dotenv import load_dotenv
from playwright.sync_api import sync_playwright
from pymongo import MongoClient
from datetime import datetime
import requests
import json
import os

def scrapePost():
    pageURL = getURL()

    if not pageURL:
        print("Error: Failed to get URL")
        return

    page = requests.get(pageURL)

    soup = BeautifulSoup(page.text, "html.parser")

    products = soup.find_all("span", attrs = {"class":"font-bold"})

    fmtProducts = []
    for product in products:
        fmtProducts.append(product.text)

    if not fmtProducts:
        print("Error: Failed to generate products")
        return

    date = soup.find("time", attrs = {"class":"copy-bitter-xs whitespace-nowrap hidden @[175px]/video:block @[225px]/article:block xl:!block"}).text.strip()

    time = datetime.now()
    formatted = time.strftime("%Y-%m-%d %H:%M:%S")

    MONGO_URI = os.getenv("MONGO_URI")
    client = MongoClient(MONGO_URI)
    db = client["posts"]
    collection = db["posts"]

    data = {
        "products" : fmtProducts,
        "date" : date,
        "source" : pageURL,
        "time_added" : time
    }

    collection.insert_one(data)

def getURL():
    baseURL = "https://www.warhammer-community.com"
    whComURL = "https://www.warhammer-community.com/en-gb/topics/pre-orders/"

    with sync_playwright() as p:
        browser = p.chromium.launch(headless=True)
        page = browser.new_page()
        page.goto(whComURL)

        page.wait_for_selector("article", timeout=10000)
        
        html = page.content()
        browser.close()

    soup = BeautifulSoup(html, "html.parser")

    links = soup.find_all("a", class_=["btn-cover", "mb-15", "md:mb-20", "block", "link-underline", "link-underline--light", "text-grimDarkBlack"])
    
    for link in links:
        if link["href"].startswith("/en-gb/articles/"):
            return baseURL + link["href"]


load_dotenv()
scrapePost() 