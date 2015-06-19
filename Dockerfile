FROM gliderlabs/alpine

RUN apk --update add go bash git
RUN git clone https://github.com/sstephenson/bats.git; ./bats/install.sh /usr/local
ENV GOPATH=$HOME/gopath
ENV PATH=$HOME/gopath/bin:$PATH
RUN go get golang.org/x/tools/cmd/vet
RUN go get golang.org/x/tools/cmd/cover
