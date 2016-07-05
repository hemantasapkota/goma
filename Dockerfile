# Based on https://github.com/docker-library/golang/blob/0ce80411b9f41e9c3a21fc0a1bffba6ae761825a/1.6/Dockerfile
FROM buildpack-deps:jessie-scm

# Setup GO
# gcc for cgo
RUN apt-get update && apt-get install -y --no-install-recommends \
                g++ \
                gcc \
                libc6-dev \
                make \
                git \
        && rm -rf /var/lib/apt/lists/*

ENV GOLANG_VERSION 1.6.2
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 e40c36ae71756198478624ed1bb4ce17597b3c19d243f3f0899bb5740d56212a

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
        && echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
        && tar -C /usr/local -xzf golang.tar.gz \
        && rm golang.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# End GO setup

# Setup gomobile
RUN go get golang.org/x/mobile/cmd/gomobile
RUN gomobile init

# Copy our source
COPY . $GOPATH/src/GetStrong/
WORKDIR $GOPATH/src/GetStrong/goandroid

#Setup Android
## TODO
#End Setup Android

# RUN go get
# CMD gomobile bind -target android -v
