| metadata | <> |
|--- | --- |
| Developer Name(s) | CJ |
| Best Contact Slack handle / Email address | cj@hivint.com/xyantix (WACTF Slack) |
| Challenge Category | Forensics |
| Challenge Tier | 3 |
| Challenge Type | FileDrop |

| Player facing | <> |
|--- | --- |
|Challenge Name | Flag Collection |
|Challenge Description | My computer died and I lost my flag collection! I think they're all there but one of them looks weird...| 
|Challenge Hint 1 | Such Volatility. |
|Challenge Hint 2 | Magic from Head to toe |

| Admin Facing | <> |
|--- | --- |
|Challenge Flag| WACTF{always_rushin_around} |
|Challenge Vuln| Get the flag |
---

Challenge PoC
1. Run Volatility over the image with the "imageinfo" command to determine the appropriate profile to use. 
```volatility imageinfo -f image.vmem```
Suggested Profile(s) : Win7SP1x64, Win7SP0x64, Win2008R2SP0x64, Win2008R2SP1x64_23418, Win2008R2SP1x64, Win7SP1x64_23418


2. Run Volatility over the image file and grep for the word "Flag"
```volatility filescan -f image.vmem --profile Win7SP1x64 | grep Flag```

3. Observe the files listed as well as their offsets.

4. Retrieve each file looking for the flag.
```
italy.jpg - volatility dumpfiles -f image.vmem --profile Win7SP1x64 -Q 0x000000001e53ea20 -n --dump-dir outputdir 
iran.jpg -  volatility dumpfiles -f image.vmem --profile Win7SP1x64 -Q 0x000000001e5ff7f0 -n --dump-dir outputdir 
russia.jpg - volatility dumpfiles -f image.vmem --profile Win7SP1x64 -Q 0x000000001e14b070 -n --dump-dir outputdir 
usa.jpg - volatility dumpfiles -f image.vmem --profile Win7SP1x64 -Q 0x000000001e814f20 -n --dump-dir outputdir 
sweden.jpg - volatility dumpfiles -f image.vmem --profile Win7SP1x64 -Q 0x000000001e5beb10 -n --dump-dir outputdir 
yemen.jpg - volatility dumpfiles -f image.vmem --profile Win7SP1x64 -Q 0x000000001e996920 -n --dump-dir outputdir 
japan.jpg - volatility dumpfiles -f image.vmem --profile Win7SP1x64 -Q 0x000000001f268440 -n --dump-dir outputdir 
australia.jpg - volatility dumpfiles -f image.vmem --profile Win7SP1x64 -Q 0x000000001fe9f4f0 -n --dump-dir outputdir
```

5. Rename the files as required (they will likely have a .dat extensions following their actual extensions)

6. Open the files and observe that all are ok except for russia.jpg. 

7. Open russia.jpg in our favourite hex editor and observe the missing Magic Header. 
```FF D8```
8. Add the required header information (can be retrieved via google or from any of the other files retrieved).

9. Open the russia.jpg file and retrieve the flag.
