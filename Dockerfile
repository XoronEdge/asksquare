FROM golang:alpine

WORKDIR /build
# . is our WORKDIR

COPY go.mod .
COPY go.sum .

# download modules
RUN go mod download
# copy all file from project to build 
COPY . .


RUN go build  cmd/main.go

WORKDIR /dist


COPY .env .
RUN ls

RUN cp /build/main  .

EXPOSE 4000

CMD ["/dist/main"]