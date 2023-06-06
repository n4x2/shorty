# Shorty
Simple URL Shortener

## Getting Started

At this time server running at http://127.0.0.1:8081. It provides the following endpoints:

- `POST /url`: create new short URL
- `GET /url`: return all shortened URLs
- `GET /:shortURl`: redirect URL

```shell
# start the server
make start
```

### Create a Shortened URL
Create new short URL

Example:
```shell
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"longUrl":"https://google.com"}' \
  http://localhost:8081/url
```
```json
{
  "shortUrl": "http://localhost:8081/e1a8159a"
}
```

### Redirect URL

Open directly shortened URL from browser or test using curl

```shell
curl -X GET http://localhost:8081/e1a8159a
```
