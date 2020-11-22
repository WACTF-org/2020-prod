| metadata | <> |
|--- | --- |
| Developer Name(s) | Russ |
| Best Contact Slack handle / Email address | russell.frame@hivint.com / rustla (WACTF Slack) |
| Challenge Category | Forensics |
| Challenge Tier | 0 |
| Challenge Type | FileDrop |

| Player facing | <> |
|--- | --- |
|Challenge Name | Is serverless, is secure |
|Challenge Description | To impress his agile DevOps friends, Batman has moved his BatCave computer to a "serverless". Oh no, someone has managed to write to his intel folders and retrieve all of his intel. Can you work out which user is the naughty agent? Luckily he's got logs enabled and has been identifying his BatCave computer on all requests, so you can rule out BatCaveSuperComputer as the flag. | 
|Challenge Hint 1 | How do browsers identify themselves to web servers? |
|Challenge Hint 2 | Maybe there's some bash tools to help cut through the noise |

| Admin Facing | <> |
|--- | --- |
|Challenge Flag| WACTF{ALL_THOSE_BATGADGETS_AND_NO_AUTHENTICATION} |
|Challenge Vuln| User-Agent leaks the flag, vuln is API with no authentication |
---

# Challenge Solution
The challenge can be solved by searching User-Agent strings (or WACTF if the player figures out the flag likely contains the flag prefix).

Run the solution PoCs, correct User-Agent (not Batman's) is in the WACTF flag format, nice and obvious for players:
- Linux: forensics-0-solution.sh
- Windows: forensics-0-solution.ps1

The player can also open the JSON file in an IDE or text editor, search for User-Agent and review all results.