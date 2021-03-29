FROM golang:1.15

WORKDIR $GOPATH/src/app
ENV GO111MODULE=on
ENV PORT=8080
ENV BASE_URL=https://www.alphavantage.co
COPY . .
# RUN go install
RUN go build -o finance_tracker
RUN chmod +x finance_tracker