| metadata | <> |
|--- | --- |
| Developer Name(s) | Chris Elliott |
| Best Contact Slack handle / Email address | ctcelliott@gmail.com |
| Challenge Category | Forensics |
| Challenge Tier | 4 |

| Player facing | <> |
|--- | --- |
|Challenge Name | Recovered Dropbox |
|Challenge Description | The security team has recovered this strange looking box from the corporate offices lobby. An image has been taken. What is going on? | 
|Challenge Hint 1  | Check the job scheduler |
|Challenge Hint 2 | Move by 13 places |

| Admin Facing | <> |
|--- | --- |
|Challenge Flag| WACTF{eps3.9_wactf4.bin} |
|Challenge Vuln| Open S3 Bucket, Weak WPA2 Passphrase, PCAP Analysis (Wireless) |
|Docker Usage Idle| NA |
|Docker Usage Expected Peak| NA |
---


# Challenge Writeup
1. The player must first mount or extract the filesystem. This can be done as follows (offset calculated by start of the root filesystem * 512)

![Mounting filesystem](writeup-screenshots/1.png)

2. Once mounted, basic enum should occur. A cron file for root can be located at /var/spool/cron/crontabs/root which contains the following: 30 * * * * /opt/fsociety/loader.sh

3. This script is strapping /opt/fsociety/fsociety.py in a tmux session and redirecting stdout to /var/log/fsociety.log.

4. The player must inspect both these files.

4. a) /var/log/fsociety.log contains the following relevant information:
    - [*] Starting aircrack-ng: ['aircrack-ng', '-w', 'wordlist', 'handshake-01.cap']
    - [*] Uploading: 1606128294.pcap
    
4. b) /opt/fsociety/fsociety.py contains the following relevant information:
    - bucket = "fsociety-pcaps"
    
5. The player should then attempt to download the pcap from the fsociety-pcaps S3 bucket. The structure for downloading an object form an s3 bucket is https://{{bucket-name}}.s3.amazonaws.com/{{object-name}}. In this case, the URL is "https://fsociety-pcaps.s3.amazonaws.com/1606128294.pcap"

6. The player must then inspect this PCAP. Encrypted 802.11 traffic is present.

7. The player must then crack the handshake stored in the PCAP to decrypt the traffic. This can be done using aircrack-ng with the wordlist from the provided image (/opt/fsociety/wordlist).

![Cracking handshake](writeup-screenshots/2.png)

8. The player must then decrypt the PCAP. This can be done through wireshark by going to Edit --> Preferences --> Protocols --> IEEE 802.11 --> Edit (next to Decryption keys) and adding FiveNine:ECorp_Office, key type wpa-pwd.

9. Once the decryption key is added, network traffic is revealed. 192.168.0.100 is regularly broadcasting a message over UDP port  5959. This can be seen in the following screenshot:

![Cracking handshake](writeup-screenshots/3.png)

10. This message is JNPGS{rcf3.9_jnpgs4.ova} which resembles the format of the flag. This has been ROT13 encoded which decodes to WACTF{eps3.9_wactf4.bin}
