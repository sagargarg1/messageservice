# Sttart from base image 1.13:
FROM golang:1.13

# Configure the repo url so we can configure our work directory:
ENV REPO_URL=github.com/saggarg/messageservice

# Setup out $GOPATH
ENV GOPATH=/app

ENV APP_PATH=$GOPATH/src/$REPO_URL


# Copy the entire source code from the current directory to $WORKPATH
ENV WORKPATH=$APP_PATH/src
COPY src $WORKPATH
WORKDIR $WORKPATH

RUN go build -o messageservice .

# Expose port 8081 to the world:
EXPOSE 8081

CMD ["./messageservice"]
