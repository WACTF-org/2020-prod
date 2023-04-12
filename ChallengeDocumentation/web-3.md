| metadata | <> |
|--- | --- |
| Developer Name(s) | Sajeeb |
| Best Contact Slack handle / Email address | [redacted] |
| Challenge Category | Web |
| Challenge Tier | 3 |
| Challenge Type | Container |

| Player facing | <> |
|--- | --- |
|Challenge Name | Old is gold |
|Challenge Description | Back in the old days, we used to create websites with very little functionality. We didn't have all the fancy frameworks available now. I don't think anyone would want to hack our manifesto, but I wonder if they could pull it off! http://web-3 | 
|Challenge Hint 1 | robots.txt often provides some additional information for challenges |
|Challenge Hint 2 | Learn about php filters and other protocols. An understand of how they can be used as a part of an attack may prove helpful! |

| Admin Facing | <> |
|--- | --- |
|Challenge Flag| WACTF{Inclusion_is_key!} |
|Challenge Vuln| LFI in loading different language versions of the same website |
|Docker Usage Idle| 0% CPU / 6MB RAM |
|Docker Usage Expected Peak| 20% CPU / 44MB RAM |
---
# docker-compose.yml
```
# The docker-compose file you deliver needs to:
# - Define which ports are exposed, and what they map to in your challenge
# - Name and tag the container for use on the docker registry
# - Define the folder that your dockerfile (and challenge source code) exists in
# The version MUST be 3.
version: '3'
# This is general structure of the docker-compose file that will be used in prod
services:
  web-3:
  # Your container name is <category>-<tier>
    container_name: web-3
    # This should be the directory your dockerfile exists in.
    # Note, other challenges will live in ./ too, so ensure your directory name is adequately unique (ie. exp-3)
    build: ./web-3/
    # Your challenges need to be tagged for pushing to the docker registry. The syntax is <domain>:<port>/wactf0x04/<container name>
    image: registry.capture.tf:5000/wactf0x04/web-3
    # Ports! the syntax is <external>:<containerlocal>
    # You can use whatever local container port you wish (as long as it's above 1024 so a non-root user can bind to it)
    # If you need to bind to a privileged port (such as running a webserver) checkout the Examples-Dockerfiles/ directory
    ports:
      - 3333:80
# Want to include something else in docker-compose that's not in this template?
# Look up its compatibility here: https://kompose.io/conversion/
```

# Challenge PoC
Using your browser, visit the URL
Click on one of the language options
Observe that the `lang` param has `.php` in it
Test out LFI payloads here
Either guess from a dir bruteforce, or visit `robots.txt`, to find that the path `config/config.php` exists
Run the final payload to get that file:
http://localhost:3333/index.php?lang=php://filter/convert.base64-encode/resource=config/config.php
Base64 decode the output and find the DB password and username, alongside the flag.