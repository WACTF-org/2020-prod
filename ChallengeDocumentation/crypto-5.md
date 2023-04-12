## Info

| metadata                                  | <>                                        |
| ----------------------------------------- | ----------------------------------------- |
| Developer Name(s)                         | Dylan                                     |
| Best Contact Slack handle / Email address | [redacted] |
| Challenge Category                        | Crypto                                    |
| Challenge Tier                            | 5                                         |
| Challenge Type                            | Container                                 |

| Player facing         | <>                                                    |
| --------------------- | ----------------------------------------------------- |
| Challenge Name        | The Oracle of Delphi                                  |
| Challenge Description | Gain access to the admin panel and retrieve the flag. http://crypto-5  | 
| Challenge Hint 1      | Try fiddling with the SecureSession cookie            |
| Challenge Hint 2      | Padding Oracles can be used to encrypt too            |

| Admin Facing   | <>                                      |
| -------------- | --------------------------------------- |
| Challenge Flag | WACTF{29e1aeea77ac559ef1c0d75e36ed2818} |
| Challenge Vuln | Padding Oracle Auth Bypass              |
|Docker Usage Idle | 5% CPU / 30MB RAM |
|Docker usage Expected Peak | 25% CPU / 60 RAM |

---

## Challenge PoC
Setup `exploit.py` as follows.
```
kali@kali:~$ python3 -m venv env
kali@kali:~$ . ./env/bin/activate
kali@kali:~$ pip install requests
Collecting requests
  Using cached requests-2.24.0-py2.py3-none-any.whl (61 kB)
Collecting urllib3!=1.25.0,!=1.25.1,<1.26,>=1.21.1
  Using cached urllib3-1.25.11-py2.py3-none-any.whl (127 kB)
Collecting certifi>=2017.4.17
  Using cached certifi-2020.6.20-py2.py3-none-any.whl (156 kB)
Collecting chardet<4,>=3.0.2
  Using cached chardet-3.0.4-py2.py3-none-any.whl (133 kB)
Collecting idna<3,>=2.5
  Using cached idna-2.10-py2.py3-none-any.whl (58 kB)
Installing collected packages: urllib3, certifi, chardet, idna, requests
Successfully installed certifi-2020.6.20 chardet-3.0.4 idna-2.10 requests-2.24.0 urllib3-1.25.11
```

The `exploit.py` script can be used to obtain and decrypt a cookie as follows.
```
kali@kali:~$ python exploit.py http://192.168.83.1:5000 decrypt
Decrypted block 1
Decrypted block 2
Decrypted block 3
Decrypted session
[b'{"Username":"gue', b'st","IsAdmin":fa', b'lse}\x0c\x0c\x0c\x0c\x0c\x0c\x0c\x0c\x0c\x0c\x0c\x0c']
```

The `exploit.py` script can be used to obtain the flag on the admin page by exploiting the padding oracle to re-encrypt the sesion cookie with `IsAdmin` set to true.
```
kali@kali:~$ python exploit.py http://192.168.83.1:5000/ getflag
Encrypted block 2
Encrypted block 1
Encrypted block 0
WACTF{29e1aeea77ac559ef1c0d75e36ed2818}
```
