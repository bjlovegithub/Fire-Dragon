# Use the official go docker image built on debian.
FROM golang:1.11

# set up the environment vars
ENV MYSQL_USER=root \
  MYSQL_PASS=billjeff \
  MYSQL_HOST=127.0.0.1:3306 \
  MYSQL_DB=test \
  JWT_SECRECT_KEY=This_is_my_test_key

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

