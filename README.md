# get-s3-metadata

An aws lambda function. Detect MIME type of the s3 Object.


## Installation
To setup the lambda, we have to create a lambda layer in aws.
**Layer** : A Lambda layer is an archive containing additional code, such as libraries, dependencies, or even custom runtimes. When you include a layer in a function, the contents are extracted to the /opt directory in the execution environment

## API Request Params

```json
    {
        "payload":base_64_encoded_format(
        {
            "source":{
                "path":"",
                "bucket":"",
                "region":""
            },
            "credentials":{
                "accesskey":"",
                "secretKey":""
            }
        }
        )
    }
```

## API Response Params
```json
    {
        "message" : "",
        "success" : true/false,
        "data" : {
            "MimeType": <MIME Type>
        }
    }
```
