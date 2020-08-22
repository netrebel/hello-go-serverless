# hello-go-serverless

Source: [Create a Serverless Application in Golang With Aws](https://schadokar.dev/posts/create-a-serverless-application-in-golang-with-aws)

# Build

```
env GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go 
env GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go
```

# Deploy

* `serverless deploy` 
* `serverless remove`