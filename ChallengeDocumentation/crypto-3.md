| metadata                                  | <>                           |
| ----------------------------------------- | ---------------------------- |
| Developer Name(s)                         | Dono                         |
| Best Contact Slack handle / Email address | zack.flack@hivint.com / dono |
| Challenge Category                        | Crypto                       |
| Challenge Tier                            | 3                            |
| Challenge Type                            | Filedrop                     |

| Player facing         | <>                                                                                                                                          |
| --------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- |
| Challenge Name        | Cache em All                                                                                                                                |
| Challenge Description | We retrieved this Enterprise Application which seems to cache credentials. Can you retrieve the credentials of the application's last user? |
| Challenge Hint 1      | dnSpy for the win                                                                                                                           |
| Challenge Hint 2      | hrrm if only that decrypt method was implemented                                                                                            |

| Admin Facing   | <>                                                    |
| -------------- | ----------------------------------------------------- |
| Challenge Flag | WACTF{This_is_a_good_way_to_cache_AD_creds_right?}    |
| Challenge Vuln | Weak obfuscation of cached credentials via RSA crypto |

---

Challenge PoC

1. Download dnspy 32bit (https://github.com/dnSpy/dnSpy/releases)
2. Open crypto-3.exe and crypto-3-library.dll in dnspy
3. Replace the `decrypt` function in `crypto-3.crypto_3.Form1` with the function below
4. At the end of the `button1_Click` function in `crypto-3.crypto_3.Form1` add the decrypt call, as shown below
5. SAVE THE BINARY TO DISK!!!!!!!!!!!!!!!!!!!
6. Run crypto-3.exe, supply random inputs to username and password, observe flag in message box

Decrypt function

````C#

private void decrypt(X509Certificate2 certificate, string currentDirectory)
{
    var privateKey = certificate.PrivateKey as RSACryptoServiceProvider;

    string path = currentDirectory + "\\.cachedCredentials";
    using (StreamReader sr = new StreamReader(path))
    {
        string line;
        while ((line = sr.ReadLine()) != null)
        {
            var encryptedBytes = System.Convert.FromBase64String(line);
            var data = privateKey.Decrypt(encryptedBytes, false);
            MessageBox.Show(System.Text.Encoding.UTF8.GetString(data), "decrypted", MessageBoxButtons.OK, MessageBoxIcon.Information);
        }
    }
}

```

Decrypt Call

```C#
this.decrypt(x509Certificate2Collection[0], currentDirectory);
````
