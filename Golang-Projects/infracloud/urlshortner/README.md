# urlshortner

A service to Shorten the URL. 

## Build using Docker

```sudo docker build . -t urlshortner```

### Call to service inside docker-container

```
curl -d '{"url":"https://ok2.com"}' -H "Content-Type: application/json" -X POST http://172.17.0.2:8080/shorturl/

```

Run the example with the optional port address for the service: 

```bash
$ go run ./cmd/main.go -http.addr :8080
```

Make POST call to Shorten URL:

```bash
curl -d '{"url":"http://infracloud.com"}' -H "Content-Type: application/json" -X POST http://localhost:8080/shorturl/
{"GENUrl":"https://cflair.com/4299216071763539886"}
```
