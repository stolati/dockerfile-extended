# Here a regular dockerfile because it's how we bootstrap
FROM golang:1.9

WORKDIR /srv
ADD . .

ENV TERM=xterm
CMD bash # Do you stuff

