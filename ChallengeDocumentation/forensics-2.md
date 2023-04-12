| metadata | <> |
|--- | --- |
| Developer Name(s) | Jack |
| Best Contact Slack handle / Email address | [redacted] |
| Challenge Category | Forensics |
| Challenge Tier | 2 |
| Challenge Type | FileDrop |

| Player facing | <> |
|--- | --- |
|Challenge Name | Network Capture |
|Challenge Description | A machine on our network got infected with some kind of malware. The sample we captured showed that it probably uses DNS as an exfil  channel, and there is a bunch of requests that don't resolve. |
|Challenge Hint 1 | filtering for no such name might be a good idea |
|Challenge Hint 2 | (Don't forget, DNS is case insensitive, so the encoding has to be case insensitive too) |

| Admin Facing | <> |
|--- | --- |
|Challenge Flag| WACTF{4FC559DEDD41B16D1EC2C384931D8BBE} |
|Challenge Vuln| Flag in DNS queries |
---

Challenge PoC  
1. Run script: solution/solve.py
2. This will retrieve the flag
