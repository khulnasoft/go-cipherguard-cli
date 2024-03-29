doc/cipherguard_delete_resource.md## cipherguard delete resource

Deletes a Cipherguard Resource

### Synopsis

Deletes a Cipherguard Resource

```
cipherguard delete resource [flags]
```

### Options

```
  -h, --help   help for resource
```

### Options inherited from parent commands

```
      --config string               Config File
      --debug                       Enable Debug Logging
      --id string                   ID of the Entity to Delete
      --mfaDelay duration           Delay between MFA Attempts, only used in noninteractive modes (default 10s)
      --mfaMode string              How to Handle MFA, the following Modes exist: none, interactive-totp and noninteractive-totp (default "interactive-totp")
      --mfaRetrys uint              How often to retry TOTP Auth, only used in nointeractive modes (default 3)
      --serverAddress string        Cipherguard Server Address (https://cipherguard.example.com)
      --timeout duration            Timeout for the Context (default 1m0s)
      --totpOffset duration         TOTP Generation offset only used in noninteractive-totp mode
      --totpToken string            Token to generate TOTP's, only used in nointeractive-totp mode
      --userPassword string         Cipherguard User Password
      --userPrivateKey string       Cipherguard User Private Key
      --userPrivateKeyFile string   Cipherguard User Private Key File, if set then the userPrivateKey will be Overwritten with the File Content
```

### SEE ALSO

* [cipherguard delete](cipherguard_delete)	 - Deletes a Cipherguard Entity

