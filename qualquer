
FROM golang:1.16-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY /conect ./conect
COPY /controler ./controler
COPY /model ./model
COPY /view ./view 
COPY index.go ./

RUN go build -o /crud
EXPOSE 9090
CMD ["/crud"]

