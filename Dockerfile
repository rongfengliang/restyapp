# build stage
FROM golang:1.9-alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ENV D=/go/src/github.com/rongfengliang/restyapp
ADD . $D
RUN cd $D && go build -o gwserver && cp gwserver /tmp/

FROM alpine:latest
WORKDIR /app
EXPOSE 8089
COPY --from=build-env /tmp/gwserver /app/gwserver
CMD ["./gwserver"]