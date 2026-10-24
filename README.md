# gen-api-key

This CLI app produces a 256 bit hex key for api usage. The format of the key will be psk_applicationName_environment_*random string*. When outputting the result, the APIKey will be written along with the hash.

## Usage

> Note: **DO NOT use underscore in the flag values!**

```bash
gen-api-key  
    -n *application name* - the client application
    -e *environment* - short, long or none, depending on your flavor
    [-j] - output as JSON
    -p prefix - defaults to "psk", supply "" to omit.
    -h - display help text.

```

All flags are optional. If none are supplied, only the default prefix `psk_` (pre shared key) will be added.  
Supply `-p ""` to omit the default prefix (`psk_`) or supply a value for your own, custom prefix.  
The idea behind the default initial part is to just indicate generic information that it's a pre shared key.  
The intended usage is:  

| Part of key      | Intended purpose                                                 |
|------------------|------------------------------------------------------------------|
| Prefix           | The serving application that produces the key and grants access. |
| Application name | The client application that supplies the key on the requests.    |
| Environment      | The environment that the application and APIKey works in         |

### Examples

#### Standard output

```bash
gen-api-key -p my-api -n my-client -e prod
APIKey: my-api_my-client_prod_5d3c4df902ae71ad3ca30fbba0eeafbcd4c84c54d990361f9869f1f50c9062cf
Hash: 1af6a0f50954805cc4fa58a12c49746087391b78c6c0201a72ac424c21b9f3ad
```

#### JSON output (piped to `jq` for formatting)

```bash
gen-api-key -p my-api -n my-client -e prod -j | jq
{
  "api-key": "my-api_my-client_prod_6822dc1b6683b15f3558f705413ef521052a3fdd797b09b9431a330a7ed0e410",
  "hash": "948364bd1f55954772686942e2e6ada08e74cb1b1838e48632c16b372b9e15fc",
  "prefix": "my-api",
  "app-name": "my-client",
  "env": "prod"
}
```
