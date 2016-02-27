FROM golang
MAINTAINER Stoney Kang, sikang99@gmail.com
RUN mkdir /app
ADD app.go /app/main.go
WORKDIR /app
RUN go build -o main .
EXPOSE 4444
CMD ["/app/main"]
