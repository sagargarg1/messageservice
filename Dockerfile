# Sttart from base image 1.12.13:
FROM golang:1.12.13

# Configure the repo url so we can configure our work directory:
ENV REPO_URL=github.com/saggarg1/messageservice

# Setup out $GOPATH
ENV GOPATH=/root/go

ENV APP_PATH=$GOPATH/src/$REPO_URL


# Copy the entire source code from the current directory to $WORKPATH
ENV WORKPATH=$APP_PATH/pkg
COPY pkg $WORKPATH
WORKDIR $WORKPATH

RUN go build -o messageservice .

# Expose port 8081 to the world:
EXPOSE 8081

CMD ["./messageservice"]
