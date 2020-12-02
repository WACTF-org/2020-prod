| metadata | <> |
|--- | --- |
| Developer Name(s) | C_Sto |
| Best Contact Slack handle / Email address | C_Sto |
| Challenge Category | Crypto |
| Challenge Tier | 1 |
| Challenge Type | Filedrop |

| Player facing | <> |
|--- | --- |
|Challenge Name | It's always the same... again |
|Challenge Description | The flag was XORed with an unknown value, then converted to hex. Can you get the flag out of the ciphertext `16000215073a0e0f041e031815041e08121e0f0e151e001e0d001306041e1204001302091e12110002043c` | 
|Challenge Hint 1 | Many online tools do XOR, but you might want to script this |
|Challenge Hint 2 | 255 values isn't a lot to look through, just brute force it |

| Admin Facing | <> |
|--- | --- |
|Challenge Flag| WACTF{ONE_BYTE_IS_NOT_A_LARGE_SEARCH_SPACE} |
|Challenge Vuln| xor sucks |
---

# docker-compose.yml

```
No
```

# Challenge PoC.py
```
https://gchq.github.io/CyberChef/#recipe=From_Hex('Auto')XOR(%7B'option':'Hex','string':'41'%7D,'Standard',false/disabled)XOR_Brute_Force(1,100,0,'Standard',false,true,false,'')&input=MTYwMDAyMTUwNzNhMGUwZjA0MWUwMzE4MTUwNDFlMDgxMjFlMGYwZTE1MWUwMDFlMGQwMDEzMDYwNDFlMTIwNDAwMTMwMjA5MWUxMjExMDAwMjA0M2M
```