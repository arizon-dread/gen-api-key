# gen-api-key

This CLI app produces a 256 bit hex key for api usage. The format of the key will be psk_applicationName_environment_<random string>. When outputting the result, the api key will be written along with the hash. 

## Usage

gen-api-key 
    -n <application name> - DO NOT use underscore in the application name since that is used as a separator in the api key parts.
    -e <environment> - short or long, depending on your flavor, example: prod | production
    [-j] - output JSON (optional)
