# hello-go-serverless

Source: [Create a Serverless Application in Golang With Aws](https://schadokar.dev/posts/create-a-serverless-application-in-golang-with-aws)

# Build

```
env GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go 
env GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go
```

> You will get the smallest binaries if you compile with -ldflags '-w -s'.
>
> The -w turns off DWARF debugging information: you will not be able to use gdb on the binary to look at specific functions or set breakpoints or get stack traces, because > all the metadata gdb needs will not be included. You will also not be able to use other tools that depend on the information, like pprof profiling.
> 
> The -s turns off generation of the Go symbol table: you will not be able to use go tool nm to list the symbols in the binary. strip -s is like passing -s to -ldflags but it > doesn't strip quite as much. go tool nm might still work after strip -s. I am not completely sure.
> 
> None of these — not -ldflags -w, not -ldflags -s, not strip -s — should affect the execution of the actual program. They only affect whether you can debug or analyze the > program with other tools.

https://stackoverflow.com/a/22276273/2500817

# Deploy

* `serverless deploy` 
* `serverless remove`