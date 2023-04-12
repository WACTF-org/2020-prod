| metadata | <> |
|--- | --- |
| Developer Name(s) | C_Sto |
| Best Contact Slack handle / Email address | [redacted] |
| Challenge Category | Crypto |
| Challenge Tier | 0 |
| Challenge Type | Filedrop |

| Player facing | <> |
|--- | --- |
|Challenge Name | It's always the same |
|Challenge Description | Crypto is quite tricky, and this is hard for a lvl 0 flag, but such is life. The flag was XORed with the value 0xab, then converted to hex. Can you get the flag out of the ciphertext? `fceae8ffedd0fce3f2f4e2f8f4e2fff4eae7fceaf2f8f4f3e4f9d6` | 
|Challenge Hint 1 | Many online tools do XOR |
|Challenge Hint 2 |  even more online tools do hex decoding |

| Admin Facing | <> |
|--- | --- |
|Challenge Flag| WACTF{WHY_IS_IT_ALWAYS_XOR} |
|Challenge Vuln| xor sucks |
---

# docker-compose.yml

```
No
```

# Challenge PoC.py
```
https://gchq.github.io/CyberChef/#recipe=From_Hex('Auto')XOR(%7B'option':'Hex','string':'ab'%7D,'Standard',false)&input=ZmNlYWU4ZmZlZGQwZmNlM2YyZjRlMmY4ZjRlMmZmZjRlYWU3ZmNlYWYyZjhmNGYzZTRmOWQ2
```