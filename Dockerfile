FROM golang:alpine AS builder

ARG GITHUB_USER="nothing"
ARG GITHUB_PASS="nothing"

RUN apk update && apk upgrade && \
    apk add --no-cache git make gcc build-base upx

#create home folder

RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

# Create a netrc file using the credentials specified using --build-arg
RUN printf "machine github.com\n\
    login ${GITHUB_USER}\n\
    password ${GITHUB_PASS}\n\
    \n\
    machine api.github.com\n\
    login ${GITHUB_USER}\n\
    password ${GITHUB_PASS}\n"\
    >> /root/.netrc
RUN chmod 600 /root/.netrc

RUN go env -w GOPRIVATE=github.com/nri4nudge/api-common
ENV GO111MODULE=on
RUN mkdir /app
WORKDIR /app

# copy go.mod and go.sum first then download deps to improve cache
COPY ./go.mod ./go.sum ./
RUN go get github.com/pwaller/goupx
RUN go get golang.org/x/lint/golint
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN goupx main

FROM alpine:3.12.1
RUN mkdir /app && \
    apk add --no-cache tzdata ca-certificates && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
WORKDIR /app
COPY --from=builder /app/main /app
# COPY --from=builder /app/migrations/* /app/migrations/
EXPOSE 80
ENTRYPOINT ["/app/main"]