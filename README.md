
# Shortener

Shorten long urls and keep hash:url key-value maps in MongoDB.

Service database agnostic and can use any DB.

### How to

Start a mongo docker container

`$ docker run -it -p 27017:27017 mongo`

Run service

`$ go run . -address "localhost"`

Use either curl for testing:

```
curl -X POST localhost/api/v1/short-url -d '{"Url": "https://test.com/very-long-url?testurl=123&x=y"}'

curl -X GET localhost/api/v1/short-url/c647adf52c439e35daf186bc2a516966
```

Or SwaggerUI:

http://localhost/swagger/index.html#/
