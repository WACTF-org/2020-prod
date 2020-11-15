| metadata | <> |
|--- | --- |
| Developer Name(s) | sudosammy |
| Best Contact Slack handle / Email address | sudosammy |
| Challenge Category | misc |
| Challenge Tier | 0 |
| Challenge Type | Container |

| Player facing | <> |
|--- | --- |
|Challenge Name | dns-check |
|Challenge Description | This container will be used to help players check their connection to OpenVPN and DNS | 
|Challenge Hint 1 | |
|Challenge Hint 2 | |

| Admin Facing | <> |
|--- | --- |
|Challenge Flag| N/A |
|Challenge Vuln| Hopefully none... |
|Docker Usage Idle| 0% CPU / 2MB RAM |
|Docker Usage Expected Peak| 10% CPU / 9MB RAM |
---

# docker-compose.yml

```
version: '3'
services:
  misc-0:
    container_name: misc-0-dns-check
    build: ./misc-0/
    image: registry.capture.tf:5000/wactf0x04/misc-0-dns-check
    ports:
      - 80:8000
    deploy:
      resources:
        limits:
          cpus: '0.10'
          memory: 15M
        reservations:
          cpus: '0.05'
          memory: 10M
    cap_drop:
      - NET_RAW
```

# Challenge PoC.py

Connect to the OpenVPN, visit http://dns-check in your browser. Win!
