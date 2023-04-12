| metadata | <> |
|--- | --- |
| Developer Name(s) | Dean |
| Best Contact Slack handle / Email address | [redacted] |
| Challenge Category | Misc |
| Challenge Tier | 1 |
| Challenge Type | Filedrop |

| Player facing | <> |
|--- | --- |
|Challenge Name | Meta |
|Challenge Description | Look at this image we have here. Can you find the flag? |
|Challenge Hint 1| Challenge name is where it's at|
|Challenge Hint 2| *Who* created the image might also tell you *how*| 

| Admin Facing | <> |
|--- | --- |
|Challenge Flag| wactf{always_be_checking_the_metadata} |
|Challenge Vuln| Zlib compressed string stored in Copyright field within an image. The Author field acts as a hint to the compression in use |

---

Challenge PoC

1. `exiftool -s -s -s -copyright -author ./src/image.png`
2. `python3 solve.py`