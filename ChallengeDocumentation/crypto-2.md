| metadata | <> |
|--- | --- |
| Developer Name(s) | C_Sto |
| Best Contact Slack handle / Email address | C_Sto |
| Challenge Category | Crypto |
| Challenge Tier | 2 |
| Challenge Type | Filedrop |

| Player facing | <> |
|--- | --- |
|Challenge Name | Psychotic Secrets |
|Challenge Description | Hey C_Sto, we pulled a bunch of secrets from this server and the encryption key, but no IV was stored with them (maybe it's the same for all secrets?)... This feels like a crypto thing, can you help? We know it was encrypted with blowfish CBC. All the ciphertexts seem like they have part of the flag in the first block, with the following block showing you the order to put them in... which sure is handy for a CTF! | 
|Challenge Hint 1 | If the IV isn't with the ciphertext, it's probably the same value for all of them |
|Challenge Hint 2 |  Look at the diagram for CBC, does the same IV for multiple ciphertexts look like a familiar crypto problem? |

| Admin Facing | <> |
|--- | --- |
|Challenge Flag| WACTF{crypto_is_actuallykinda_hard_to_get_right} |
|Challenge Vuln| Re-Used IV = many time pad |
---

# docker-compose.yml

```
No
```

# Challenge PoC.py
```
see cmd/solve.go (run with go run solve.go after pasting in the values etc)
```