Challenge Data:

|metadata | <> |
|--- | --- |
|Company | NCC Group |
|Developer Name(s) | Peter Hannay & Clinton Carpene |
|Contact | [redacted] |
|Challenge Category | Forensics |
| Challenge Tier | 1 |
| Challenge Type | FileDrop |


| Player facing | <> |
|--- | --- |
|Challenge Name | Parity |
|Challenge Description | We've recovered the hard disks from a machine. They were running ZFS, but we're having trouble mounting it because computers are hard. Mount the volume, get the flag. | 
|Challenge Hint 1 | The raw disks may need something done to them to import them into your VM software of choice. Importing and mounting are different things, might need some descriptor files. | 
|Challenge Hint 2 | Do some googling, you might need to force some things to make it work. zpool it up  |

| Admin Facing | <> |
|--- | --- |
|Challenge Flag| WACTF{parity_get_it_lol} |
|Challenge Vuln| Mount the array, get the flag |
Challenge PoC
---

 * Add disks to your VM of choice, depending on the type of VM you use this might be easy or hard idk.
   * For VMware, create a descriptor file for each disk as follows (https://communities.vmware.com/t5/ESXi-Discussions/getting-a-raw-hard-disk-image-into-a-VM/m-p/2708613):
        ~~~
        # Disk DescriptorFile
        version=1
        CID=11307a7e
        parentCID=ffffffff
        createType="vmfs"

        # Extent description
        RW 188416 VMFS "disk1.vmdk"

        # The Disk Data Base 
        #DDB

        ddb.adapterType = "lsilogic"
        ddb.encoding = "windows-1252"
        ddb.geometry.cylinders = "11"
        ddb.geometry.heads = "255"
        ddb.geometry.sectors = "63"
        ddb.longContentID = "421a21a1a1d4be24b6c4aced11307a7e"
        ddb.toolsInstallType = "4"
        ddb.toolsVersion = "10304"
        ddb.virtualHWVersion = "16"
        ~~~
    * Then add each disk via it's descriptor to a linux VM.
 * run zpool import and observe the name of the pool
   * players will likely need to install this first.
 * zpool import -f tank
 * cd /tank
 * cat password.txt
 * 7z e flag.7z
 * cat flag.txt
