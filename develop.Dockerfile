FROM public.ecr.aws/docker/library/golang:1.24.5-alpine

WORKDIR /go/src/github.com/yuita-yoshihiko/daredemo-design-backend
COPY go.mod .
COPY go.sum .

RUN apk add --no-cache git alpine-sdk
RUN set -x \
    && go mod download

COPY . .
