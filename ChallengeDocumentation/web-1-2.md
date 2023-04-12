| metadata | <> |
|--- | --- |
| Developer Name(s) | Lynkle |
| Best Contact Slack handle / Email address | [redacted] |
| Challenge Category | Web |
| Challenge Tier | 1 |
| Challenge Type | Container |

| Player facing | <> |
|--- | --- |
|Challenge Name | Bird-House |
|Challenge Description | The admin has taken a photo of some rare birds. Find the photo with the flag hiden in the site. http://web-1-2  | 
|Challenge Hint 1 | Pay attention to the URLs |
|Challenge Hint 2 | Look for the IDs and work backwards |

| Admin Facing | <> |
|--- | --- |
|Challenge Flag| WACTF{birds_everywhere} |
|Challenge Vuln| Insecure Direct Object Reference |
|Docker Usage Idle| 0% CPU / 15~MB RAM |
|Docker Usage Expected Peak| 15% CPU / 50~MB RAM |
---

# docker-compose.yml

```
version: '3'
services:
  web-1-2:
    container_name: web-1-2
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
Create an account or login to the web application.

Navigate to any individual photo by clicking on it in the gallery.

Roll the ID in the url http://web-1-2/images/view/{{ID}} back until you reach ID 1.

The flag is in the photo itself.
