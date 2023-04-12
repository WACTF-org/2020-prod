| metadata | <> |
|--- | --- |
| Developer Name(s) | Dono |
| Best Contact Slack handle / Email address | [redacted] |
| Challenge Category | Web |
| Challenge Tier | 5 |
| Challenge Type | Container |

| Player facing | <> |
|--- | --- |
|Challenge Name | ggrf |
|Challenge Description | Check out this crackin HTML5 web site. Our GRC team told us that they run a highly sensitive flag service on this server... http://web-5 |
|Challenge Hint 1 | Check all the params |
|Challenge Hint 2 | No file for you. Check the top ones. |

| Admin Facing | <> |
|--- | --- |
|Challenge Flag| WACTF{59d7dd1cfd283ead53a8c9e28719903a} |
|Challenge Vuln| XSS in PDF generation which leads to SSRF, which can be used to enumerate local ports and retrieve the flag |
|Docker Usage Idle | 3% CPU / 10MB RAM |
|Docker usage Expected Peak | 8% CPU / 30MB RAM |
---

Challenge PoC
1. Upload `b.js` to shellbox.
2. Serve `b.js` on port 8000, vi `python3 -m http.server`
3. Run `solve.py`, open the retrieved PDF, and view the flag.