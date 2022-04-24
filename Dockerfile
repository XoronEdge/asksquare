FROM golang:alpine as builder

WORKDIR /build
# . is our WORKDIR

COPY go.mod .
COPY go.sum .

# download modules
RUN go mod download
# copy all file from project to build 
COPY . .


RUN go build  cmd/main.go


FROM alpine:latest  
WORKDIR /dist
# COPY .env .
COPY --from=builder /build/main .

EXPOSE 3000

CMD ["/dist/main"]