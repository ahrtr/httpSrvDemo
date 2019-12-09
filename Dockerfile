FROM golang:1.13.4-alpine AS build
LABEL maintainer="Benjamin Wang <benjamin_wang@aliyun.com>"

WORKDIR /go/src/httpSrvDemo/
COPY . /go/src/httpSrvDemo/
RUN CGO_ENABLED=0 go build -o /bin/httpSrvDemo

FROM scratch
LABEL maintainer="Benjamin Wang <benjamin_wang@aliyun.com>"

COPY --from=build /bin/httpSrvDemo /bin/httpSrvDemo 
ENTRYPOINT ["/bin/httpSrvDemo"]
