from bs4 import BeautifulSoup
import requests
import json
import pymongo

def scrapePost():
    pageURL = "https://www.warhammer-community.com/en-gb/articles/hjdxp86f/saturday-pre-orders-warhammer-40000-armageddon/"
    page = requests.get(pageURL)

    soup = BeautifulSoup(page.text, "html.parser")

    products = soup.find_all("span", attrs = {"class":"font-bold"})
    date = soup.find("time", attrs = {"class":"copy-bitter-xs whitespace-nowrap hidden @[175px]/video:block @[225px]/article:block xl:!block"})

    print(date.text)
    for product in products:
        print(product.text)
    

scrapePost()

    