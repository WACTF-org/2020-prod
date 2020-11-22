| metadata | <> |
|--- | --- |
| Developer Name(s) | Clinton Carpene, Peter Hannay |
| Best Contact Slack handle / Email address | clinton.carpene@nccgroup.com/swarley (WACTF Slack) peter.hannay@nccgroup.com/kronicd (WACTF Slack) |
| Challenge Category | Web |
| Challenge Tier | 2 |
| Challenge Type | Container |

| Player facing | <> |
|--- | --- |
|Challenge Name | Every vote counts! |
|Challenge Description | It's important to do your democratic duty and vote. But wouldn't it be a shame to see your preferred candidate lose in what is obviously an unfair system? You only get one vote; make it count. Your login credentials for voting are: - Username: `lastvoter` - Password: `yourvotecounts123` - Preferred Candidate: `LA Ice Cola`. http://web-2-2 | 
|Challenge Hint 1 | You only get one vote, but how many times can you use it? |
|Challenge Hint 2 | When should you actually vote in order to rig the election. (Google TOCTOU) |

| Admin Facing | <> |
|--- | --- |
|Challenge Flag| WACTF{democracy_wins_everytime} |
|Challenge Vuln| Election fraud through vote reuse / TOCTOU |
|Docker Usage Idle | 1% CPU / 12MB RAM |
|Docker usage Expected Peak | 35% CPU / 30MB RAM |

---

Challenge PoC
1. Run `python3 solution/solution.py`
2. This will retrieve the flag
