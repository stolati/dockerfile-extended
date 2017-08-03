# Here a regular dockerfile because it's how we bootstrap
FROM golang:1.9

WORKDIR /srv
ADD . .

RUN ls
RUN bash build.bash

