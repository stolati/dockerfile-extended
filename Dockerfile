FROM golang:1.9

# Here a regular dockerfile because it's how we bootstrap


WORKDIR /srv
ADD . .

ENV GOPATH=/srv/src/

RUN go get github.com/Masterminds/sprig
RUN go get github.com/Masterminds/sprig


