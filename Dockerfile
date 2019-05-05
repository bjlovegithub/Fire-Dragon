# Use the official go docker image built on debian.
FROM golang:1.11

ARG mysql_user
ARG mysql_pass
ARG mysql_host_port
ARG mysql_db
ARG jwt_sec_key

# set up the environment vars
ENV MYSQL_USER=$mysql_user \
  MYSQL_PASS=$mysql_pass \
  MYSQL_HOST=$mysql_host_port \
  MYSQL_DB=$mysql_db \
  JWT_SECRECT_KEY=$jwt_sec_key

# Grab the source code and add it to the workspace.
ADD . /go/src/Fire-Dragon

# Install revel and the revel CLI.
RUN go get github.com/revel/revel && \
  go get github.com/revel/cmd/revel && \
  revel build /go/src/Fire-Dragon /go/src/app dev

# Open up the port where the app is running.
EXPOSE 443

# start the app
ENTRYPOINT /go/src/app/run.sh

