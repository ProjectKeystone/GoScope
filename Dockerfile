FROM golang:alpine  
WORKDIR /scanner

COPY . /scanner

RUN go build -o GoScope 

ENTRYPOINT ["./GoScope"]
