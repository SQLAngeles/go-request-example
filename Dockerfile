# start from base golang image which installs golang and sets GOPATH
FROM golang

# copy the local project code into a directory in the container’s GOPATH
COPY ./ /go/src/github.com/[your-github-account]/go-request-example/
# set this to the working directory so all subsequent commands will run from this directory
WORKDIR /go/src/github.com/[your-github-account]/go-request-example/

#  install all dependencies
RUN go get ./
# build the binary
RUN go build

# download fresh package
RUN go get github.com/pilu/fresh
# use pilu/fresh to hot reload the code when it changes
CMD fresh

# This is the port the app was programed to serve on
EXPOSE 8000