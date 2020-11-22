Challenge Data:

| metadata |  |
|--- | --- |
| Developer Name(s) | Luke |
| Best Contact Slack handle / Email address | Luke (WACTF Slack) / luke.healy@cybercx.com.au |
| Challenge Category | Exploit |
| Challenge Tier | 4 |
| Challenge Type | Container |

| Player facing |  |
|--- | --- |
|Challenge Name | snek |
|Challenge Description | Someone has rolled their own network access control tool. I'm sure it's fine. | 
|Challenge Hint 1 | I bet they stoled a crypto library off github... |
|Challenge Hint 2 | How's the input validation? |
|Docker Usage Idle | 5% CPU / 10MB RAM |
|Docker usage Expected Peak | 10% CPU / 50MB RAM |

| Admin Facing |  |
|--- | --- |
|Challenge Flag| WACTF{eval('Your life decisions if you use eval()')} |
|Challenge Vuln| Unsafe eval() leading to arb python command injection. |
| Docker Usage Idle | 0.1% CPU / 10MB RAM |
| Docker Usage Expected Peak | 1% CPU / 14MB RAM |
---

Challenge PoC
1. nc -lnvp 1300
2. python3 solution.py
3. Please note that the challenge is shellable.
