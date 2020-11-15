# 2020-prod
Oh god why

# Git aw yeah such devops
Since we are (mis)using git, you may run into problems when commiting etc. Realistically we should have a branch + merge workflow, but that's way too nerdy. You are most likely to run into git merge issues/conflicts on two files, this readme.md (because you're updating the status), and the main docker-compose.yaml. Sometimes it will work fine, sometimes it really won't. [This link might help you in your struggles](https://docs.github.com/en/free-pro-team@latest/github/collaborating-with-issues-and-pull-requests/resolving-a-merge-conflict-using-the-command-line), otherwise just shout about it in the QA slack chan, and someone will either laugh at your pain (C_Sto), or might help you out.

# Progress

- (empty) = Unsubmitted
- X = The repo has something in it
- R = Passed QA ready for K8 integration test
- $$$ = Production ready

|Challenge | Complete | Individual Responsible (challenge description) | Link to Repo | Link to QA Issue |
|---|---|---|---|---|
|Web - 0 |  | Deem |
|Web - 1 | X | Lincoln, Laura | [web-1](https://github.com/WACTF-org/legendoflynkle-challenges/tree/master/web-1) (possible dupe [web-1](https://github.com/WACTF-org/ld-preload-challenges/tree/master/web-1-dependencycheck-filedrop) possible another dupe [web-1](https://github.com/WACTF-org/sysophost-challenges/tree/master/web-1)) |
|Web - 2 | X | Dr Clinton Carpington (CEH), Laura | [web-2](https://github.com/WACTF-org/ld-preload-challenges/tree/master/web-2-hardcodedsecrets) (!possible dupe) | |
|Web - 3 |  | Sajeeb |
|Web - 4 |  | Lincoln |
|Web - 5 | X | Dono | [web-5](https://github.com/WACTF-org/dzflack-challenges/tree/master/web-5) | |
|Exploit - 0 | X | Luke | [exploit-0](https://github.com/WACTF-org/lukehealy-challenges/tree/master/exp-0) | [exploit-0](https://github.com/WACTF-org/lukehealy-challenges/issues/1) |
|Exploit - 1 |  | Chris |
|Exploit - 2 |  | Jorel |
|Exploit - 3 | X | Cam | [exploit-3](https://github.com/WACTF-org/c-sto-challenges/tree/master/crypto-2) | [exploit-3](https://github.com/WACTF-org/c-sto-challenges/issues/1) |
|Exploit - 4 | X | Luke | [exploit-4](https://github.com/WACTF-org/lukehealy-challenges/tree/master/exp-4) | [exploit-4](https://github.com/WACTF-org/lukehealy-challenges/issues/2)
|Exploit - 5 | X | Luke | [exploit-5](https://github.com/WACTF-org/lukehealy-challenges/tree/master/exp-5) | |
|Crypto - 0 |  | Cam |
|Crypto - 1 |  | Cam I guess |
|Crypto - 2 | X | Cam | [crypto-2](https://github.com/WACTF-org/c-sto-challenges/tree/master/crypto-2) | [crypto-2](https://github.com/WACTF-org/c-sto-challenges/issues/2) |
|Crypto - 3 |  | Dono |
|Crypto - 4 | X | Dylan | [crypto-4](https://github.com/WACTF-org/dpindur-challenges/tree/master/crypto-4) | [crypto-4](https://github.com/WACTF-org/dpindur-challenges/issues/2)
|Crypto - 5 | X | Dylan | [crypto-5](https://github.com/WACTF-org/dpindur-challenges/tree/master/crypto-5) | [crypto-5](https://github.com/WACTF-org/dpindur-challenges/issues/3)
|Forensics - 0 | X | Russ | [forensics-0](https://github.com/WACTF-org/rustla-challenges/blob/master/forensics-0.md) | [forensics-0](https://github.com/WACTF-org/rustla-challenges/issues/1)
|Forensics - 1 | X |  Dr Pepper Pig (CISSP) | [forensics-1](https://github.com/WACTF-org/kronicd-challenges/tree/master/df-1) |
|Forensics - 2 | X |  Jack N | [forensics-2](https://github.com/WACTF-org/jib1337-challenges/tree/master/forensics-2) | [forensics-2](https://github.com/WACTF-org/jib1337-challenges/issues/2)
|Forensics - 3 | X | CJ, Chris | [forensics-3](https://github.com/WACTF-org/xyantix-challenges/blob/master/forensics-3.md) (! possible dupe) | |
|Forensics - 4 | X | CJ (jack???) | [forensics-4](https://github.com/WACTF-org/jib1337-challenges/tree/master/forensics-4) | [forensics-4](https://github.com/WACTF-org/jib1337-challenges/issues/1) |
|Forensics - 5 |  | Jack N |
|Misc - 0 | ??? | (?) Sudosammy (dns-check) | |
|Misc - 1 | X | sysophost | [misc-1](https://github.com/WACTF-org/sysophost-challenges/tree/master/misc-1) | [misc-1](https://github.com/WACTF-org/sysophost-challenges/issues/1) |
|Misc - 2 | X | Dr Perpper Pig (CISSP) | [misc-2](https://github.com/WACTF-org/kronicd-challenges/tree/master/misc-2) |
