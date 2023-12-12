# Publc Keys from JWKS

Get Public keys from the *modulus* and *exponent* values of the provided **JWKS**
endpoint of any provider

## Sample run

> Usage

```bash
λ go run main.go
usage: go run main <jwks url>
```


> Output for valid endpoint

```{shell}
λ go run main.go https://csmc.qa.auth.united.com/oauth2/v1/keys
2023/04/30 18:57:06 * For Kid: 0d720bb4-9512-4e56-b5bc-5f3030ff8363
2023/04/30 18:57:06 * Public Key:
-----BEGIN RSA PUBLIC KEY-----
MIIBITANBgkqhkiG9w0BAQEFAAOCAQ4AMIIBCQKCAQAAtOyHtF2xaYPm05MTh9Cl
X6KPBYpd/OjJMPfu3yWkcgiR1pBqT/ChP0fxlbJyvnT5hVj+jbxWMifeuMwPnxvN
uHHCqgLS1X3wtr5+u+NPs/G8+8P2vsDGSlnNNOif3d4scNEoS/3kKCsz2xeueeJc
fCETUg3xzTVhg07aWDD5PC6bICGB80wSD7lDtR2Yry/P1/ZuGceMMUnD/ntfODwM
T2ttW41Et1M3kwY+y5dVMQqgaNvOWxYnm0iolGTiaL4yqvW+NrEZUCJUYtajVi5M
ZRtV/ueSN4vfAqmmUvyzIgRIaVdgTS6pO+RTgwyinSOPj8K+2PnNALUfhYTbygik
AgMBAAE=
-----END RSA PUBLIC KEY-----
...
...
...
2023/04/30 18:57:06 * For Kid: 8b5eea37-ba11-44cc-824e-17115a59ed26
2023/04/30 18:57:06 * Public Key:
-----BEGIN RSA PUBLIC KEY-----
MIIBITANBgkqhkiG9w0BAQEFAAOCAQ4AMIIBCQKCAQAAyE2AhqRpGj1JB5Bqki/M
Tk5PKWKjiBXRR34WbKE7lkkkurPQ+nD7ZDAFtTW+NcHioHr+JPgAmrxkJ3fhP9Cb
CxAd+6FYfWTbVyNkiK8YhTPBtVygVShO+4Z89l8LkCcLgGHUwJ5eHjjDyVS8flNW
gkViC1ZDKZ/IA9MCLH+icWRBvAyPs0dvE7zE3AwbiKJMZGP1dPy5YfhDKSM2/APd
6psHWjs3scPk7Tzm9J156qgRIGTB3zoXITG1UmowKKThPVOq+yld4NP+IqfR0WsC
xE6bysEHCQKWjXFxad0k4RZ4xibm/8ni/od4mLvHG5ofe6VIH0bToVfbm164Xxd4
AgMBAAE=
-----END RSA PUBLIC KEY-----
```

## Other endpoints to try

We can extract `PublicKey` information from any available endpoints.

# from https://login.microsoftonline.com/common/v2.0/.well-known/openid-configuration
```{shell}
go run main.go https://login.microsoftonline.com/common/discovery/v2.0/keys

go run main.go https://www.googleapis.com/oauth2/v2/certs

```
