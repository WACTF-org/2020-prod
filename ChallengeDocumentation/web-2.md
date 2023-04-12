| metadata | <> |
|--- | --- |
| Developer Name(s) | Laura |
| Best Contact Slack handle / Email address | [redacted] |
| Challenge Category | Web |
| Challenge Tier | 2 |
| Challenge Type | Container and Filedrop |

| Player facing | <> |
|--- | --- |
|Challenge Name | Hardcoded secrets |
|Challenge Description | You have managed to obtain part of the source of this nodejs app. It contains secrets! Use the secrets to to obtain the flag! | 
|Challenge Hint 1 | Secrets are often used to obtain authorisation. Look for the service running on port 3000.  |
|Challenge Hint 2 | 'Basic Authentication headers as base64 encoded. Perhaps the key is too!' |

| Admin Facing | <> |
|--- | --- |
|Challenge Flag| WACTF{"Happy as a Clam!"} |
|Challenge Vuln| Get the flag |
|Docker Usage Idle | 3% CPU / 10MB RAM |
|Docker usage Expected Peak | 8% CPU / 30MB RAM |
---

Challenge PoC
1. solution.txt
2. run the curl command to obtain the flag

the code supplied to WACTF competitors should be everything in /docker-files/ locally, apart from the /views directory containing the flag

command:

curl --request GET \
  --url http://127.0.0.1:3000/flag \
  --header 'accept: application/json' \
  --header 'authorization: Basic U3VwZXJTZWN1cmVQYXNzd29yZGZvcnVzZXJhZG1pbg==' \
  --header 'content-type: application/x-www-form-urlencoded'