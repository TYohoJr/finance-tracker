FROM golang:1.15

WORKDIR $GOPATH/src/app
ENV GO111MODULE=on
ENV PORT=8080
ENV BASE_URL=https://www.alphavantage.co
ENV API_TOKEN=X86NAKJDEADUW1PG
COPY . .
RUN go install
RUN go build -o stocktracker_binary
RUN chmod +x stocktracker_binary