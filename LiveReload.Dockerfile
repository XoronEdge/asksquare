FROM golang:alpine as builder

RUN apk update && apk upgrade  && \
    apk add  git \    
    make openssh-client

RUN apk --no-cache add curl

WORKDIR /app 

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD air