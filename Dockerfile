FROM golang as build

ENV GOPROXY=https://goproxy.io

ADD . /ee-gallery-backend

WORKDIR /ee-gallery-backend

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_server

FROM alpine:3.7

WORKDIR /www

COPY --from=build /ee-gallery-backend/api_server /usr/bin/api_server
ADD ./configs /www/configs

RUN chmod +x /usr/bin/api_server

ENTRYPOINT ["api_server"]