| metadata                                  | <>                                             |
| ----------------------------------------- | ---------------------------------------------- |
| Developer Name(s)                         | Laura                                          |
| Best Contact Slack handle / Email address | [redacted] |
| Challenge Category                        | Web                                            |
| Challenge Tier                            | 1                                              |
| Challenge Type                            | FileDrop                                       |

| Player facing         | <>                                                                                                                                                                                                                                                                             |
| --------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Challenge Name        | Find the vulnerable dependency                                                                                                                                                                                                                                                 |
| Challenge Description | This bad JS app includes vulnerable dependencies. Use OWASP dependency check to identify the out-of-date JS packages. Note: the flag does not have the WACTF prefix in the challenge, but must be submitted with the WACTF prefix on the scoreboard (eg WACTF{flag_goes_here}) |
| Challenge Hint 1      | Use the --scan argument to search for the vulnerable packages within the bad source.                                                                                                                                                                                           |
| Challenge Hint 2      | 'dependency-check.sh --scan Secure-js-app/ -f CSV .'                                                                                                                                                                                                                           |

| Admin Facing   | <>                                |
| -------------- | --------------------------------- |
| Challenge Flag | WACTF{park-rotation-banana!!!!!!} |
| Challenge Vuln | Get the flag                      |

---

Challenge PoC

1. Run solution/solution.sh
2. This will retrieve the flag
