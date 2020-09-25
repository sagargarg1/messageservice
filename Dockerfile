# Sttart from base image 1.12.13:
FROM golang:1.12.13

ENV GO111MODULE=on

# Setup out $GOPATH
ENV GOPATH=/root/go/
RUN mkdir -p $GOPATH

ENV APP_PATH=$GOPATH/src/
RUN mkdir -p $APP_PATH

# Copy the entire source code from the current directory to $WORKPATH
ENV WORKPATH=$APP_PATH
COPY * $WORKPATH
RUN ls $WORKPATH
WORKDIR $WORKPATH

RUN go build -o messageservice .
# Expose port 8081 to the world:

EXPOSE 8081
CMD ["./messageservice"]
