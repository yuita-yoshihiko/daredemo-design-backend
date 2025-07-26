FROM public.ecr.aws/docker/library/golang:1.24.5-alpine AS builder

WORKDIR /go/src/github.com/yuita-yoshihiko/daredemo-design-backend
COPY go.mod .
COPY go.sum .

RUN apk add --no-cache git alpine-sdk
RUN set -x \
    && go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /server ./cmd/main.go


FROM public.ecr.aws/docker/library/alpine:3.22

EXPOSE 80
RUN apk update \
    && apk upgrade \
    && apk add ca-certificates \
    && rm -rf /var/cache/apk/* \
    && apk add tzdata \
    && mkdir -m 666 -p /tmp

COPY --from=builder /server ./app/server

RUN apk add libcap \
    && setcap cap_net_bind_service=+ep ./app/server

USER nobody
CMD ["./app/server"]
