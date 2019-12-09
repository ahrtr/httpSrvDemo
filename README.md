# httpSrvDemo
httpSrvDemo is just a demo of web service application.

## Command line parameters

| Option | Default value | Description |
|--------|---------------|-------------|
| --port <port> | 8080 | listening port |
| --routepath <path> | / | route path to receive the http requests |
| --respmsg <resp message> | "Hello world!" | the response message to be sent back to clients |


## How to build & run httpSrvDemo
### Build
```
$ export TAG=abc # if not provided, then the HEAD commit id is used.
$ make
```

### Run
Use all the default values,
```
$docker run -d -p 8080:8080 httpsrvdemo:<tag> 
```

Then access the web service:
```
$ curl http://localhost:8080
```

You can also specify the command line parameters
```
$docker run -d -p 8080:8081 httpsrvdemo:<tag> --port 8081 --routepath '/hello' --respmsg 'hello benjamin!'
```

```
$ curl http://localhost:8080/hello
```
