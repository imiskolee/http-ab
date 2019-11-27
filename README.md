# http-ab
`http-ab` is a test tool for quickly compare two API endpoint response schema(included status code,response header,response body).

# Use Case
* Migrate old API to new implementation.

# Install

```sh
go get -u https://github.com/imiskolee/http-ab
```

# Configuration

```yaml
common: &common
  uri: /a/b?c=d&e=f#g=1
  method: GET
  content_type: application/json
  headers:
    FOO: BAR
  payload: "anything here..."
    
mainly: 
  <<: *common
  url: https://for.bar

secondly:
  <<: *common
  url: https://foo.bar
  
response_parser: json
```

# Try

```bash
http-ab --config ./example.yaml
```

# Thanks

* [JSON Diff library](https://github.com/josephburnett/jd)