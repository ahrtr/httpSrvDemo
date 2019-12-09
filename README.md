# httpSrvDemo
httpSrvDemo is just a demo http server application.

## Command line parameters

| Option | Default value | Description |
|--------|---------------|-------------|
| --port <port> | 8080 | listening port |
| --routepath <path> | / | route path to receive the http requests |
| --respmsg <resp message> | "Hello world!" | the response message to be sent back to clients |


## How to build & run httpSrvDemo

```
export TAG=abc # if not provided, then the HEAD commit id is used.
make
docker run -d -p 8080:8080 httpsrvdemo:<tag> 
# You can also specify the command line parameters
#docker run -d -p 8080:8081 httpsrvdemo:<tag> --port 8081 --routepath '/hello' --respmsg 'hello benjamin!'
```